#include "markdown.hpp"

std::vector<std::string> BuildMarkdownCommand(const std::string &width, const std::filesystem::path &file)
{
	return {"glow", "-w", width, file};
}
