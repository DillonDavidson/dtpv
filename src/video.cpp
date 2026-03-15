#include "video.hpp"

#include <fcntl.h>
#include <string>
#include <sys/wait.h>
#include <unistd.h>

std::vector<std::string> BuildVideoCommand(const std::string &width, const std::string &height, const fs::path &file)
{
	auto cache = MakeCachePath(file);

	if (!IsCacheValid(file, cache)) {
		GenerateThumbnail(file, cache);
		ClearCacheIfTooBig(GetCacheDirectory(), 500ull * 1024 * 1024);

		// This may not be necessary
		if (!fs::exists(cache)) {
			return {}; // thumbnail generation failed
		}
	}

	return {"chafa", "-s", width + "x" + height, "-f", "sixels", cache};
}

void ClearCacheIfTooBig(const fs::path &cache_dir, std::uintmax_t max_size)
{
	std::uintmax_t total = 0;

	for (const auto &entry : fs::directory_iterator(cache_dir)) {
		if (fs::is_regular_file(entry)) {
			total += fs::file_size(entry);
		}
	}

	if (total <= max_size) {
		return;
	}

	for (const auto &entry : fs::directory_iterator(cache_dir)) {
		fs::remove_all(entry);
	}
}

std::string GetCacheDirectory()
{
	const char *xdg = std::getenv("XDG_CACHE_HOME");
	fs::path cacheDirectory = xdg ? xdg : fs::path(std::getenv("HOME")) / ".cache";
	cacheDirectory /= "dtpv";

	if (!fs::exists(cacheDirectory)) {
		fs::create_directories(cacheDirectory);
	}

	return cacheDirectory.string();
}

fs::path MakeCachePath(const fs::path &file)
{
	fs::directory_entry entry(file);

	if (!entry.exists() || !entry.is_regular_file()) {
		return {};
	}

	fs::path canonicalPath;

	try {
		canonicalPath = fs::weakly_canonical(file);
	} catch (...) {
		return {};
	}

	std::stringstream key;
	key << canonicalPath.string() << entry.file_size() << entry.last_write_time().time_since_epoch().count();

	std::size_t hash = std::hash<std::string>{}(key.str());

	return fs::path(GetCacheDirectory()) / (std::to_string(hash) + ".jpg");
}

bool IsCacheValid(const fs::path &src, const fs::path &cache)
{
	fs::directory_entry srcEntry(src);
	fs::directory_entry cacheEntry(cache);

	if (!srcEntry.exists() || !cacheEntry.exists()) {
		return false;
	}

	return cacheEntry.last_write_time() >= srcEntry.last_write_time();
}

void GenerateThumbnail(const fs::path &src, const fs::path &cache)
{
	std::string lock = cache.string() + ".lock";

	if (access(lock.c_str(), F_OK) == 0) {
		return;
	}

	int fd = open(lock.c_str(), O_CREAT | O_EXCL, 0644);
	if (fd < 0) {
		return;
	}

	close(fd);

	pid_t pid = fork();

	if (pid == 0) {
		std::string srcStr = src.string();
		std::string cacheStr = cache.string();

		execlp("ffmpegthumbnailer", "ffmpegthumbnailer", "-i", srcStr.c_str(), "-o", cacheStr.c_str(), "-s",
		       "0", "-t", "50%", nullptr);

		_exit(1);
	}

	if (pid > 0) {
		waitpid(pid, nullptr, 0);
	}

	unlink(lock.c_str());
}
