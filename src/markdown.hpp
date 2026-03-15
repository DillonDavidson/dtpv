#pragma once

#include <filesystem>
#include <string>
#include <vector>

std::vector<std::string> BuildMarkdownCommand(const std::string &width, const std::filesystem::path &file);
