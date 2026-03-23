#include "image.hpp"

std::vector<std::string> BuildImageCommand(const std::string &width, const std::string &height, const fs::path &file)
{
	return {"chafa", "-s", width + "x" + height, "-f", "sixels", "--bg", "black", "--polite", "on", file};
}
