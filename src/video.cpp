#include "video.hpp"

#include "cache.hpp"

#include <fcntl.h>
#include <string>
#include <sys/wait.h>
#include <unistd.h>
#include <vector>

std::vector<std::string> BuildVideoCommand(const std::string &width, const std::string &height, const fs::path &file)
{
	ClearCacheIfTooBig(GetCacheDirectory(), 500ull * 1024 * 1024);

	auto cache = MakeCachePath(file);
	std::vector<std::string> args = {"ffmpegthumbnailer", "-i", file, "-o", cache, "-s", "0", "-t", "50%"};

	if (!IsCacheValid(file, cache)) {
		GenerateThumbnail(file, cache, args);
	}

	return {"chafa", "-s", width + "x" + height, "-f", "sixels", "--bg", "black", "--polite", "on", cache};
}
