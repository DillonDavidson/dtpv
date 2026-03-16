#pragma once

#include <filesystem>
#include <string>
#include <vector>

namespace fs = std::filesystem;

std::vector<std::string> BuildImageCommand(const std::string &width, const std::string &height, const fs::path &file);
