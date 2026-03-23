#include "pdf.hpp"

#include "cache.hpp"

#include <fcntl.h>
#include <sys/wait.h>
#include <unistd.h>

std::vector<std::string> BuildPDFCommand(const std::string &width, const std::string &height, const fs::path &file)
{
	ClearCacheIfTooBig(GetCacheDirectory(), 500ull * 1024 * 1024);
	auto cache = MakeCachePath(file);
	std::string path = cache.string().substr(0, cache.string().size() - 4); // drop the .jpg
	std::vector<std::string> args = {"pdftoppm",    "-f", "1",           "-l",    "1",  "-scale-to-x", "1920",
	                                 "-scale-to-y", "-1", "-singlefile", "-jpeg", file, path};

	if (!IsCacheValid(file, cache)) {
		GenerateThumbnail(file, cache, args);
	}

	return {"chafa", "-s", width + "x" + height, "-f", "sixels", "--bg", "black", "--polite", "on", cache};
}
