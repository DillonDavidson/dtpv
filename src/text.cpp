#include "text.hpp"

std::vector<std::string> BuildTextCommand(const std::string &width, const std::filesystem::path &file)
{
	return {"bat",
	        "--color=always",
	        "--style=plain",
	        "--paging=never",
	        "--wrap=character",
	        "--terminal-width",
	        width,
	        "--",
	        file};
}
