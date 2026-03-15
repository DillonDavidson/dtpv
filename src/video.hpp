#pragma once

#include <filesystem>
#include <vector>

namespace fs = std::filesystem;

std::vector<std::string> BuildVideoCommand(const std::string &width, const std::string &height, const fs::path &file);
void ClearCacheIfTooBig(const fs::path &cache_dir, std::uintmax_t max_size);
std::string GetCacheDirectory();
fs::path MakeCachePath(const fs::path &file);
bool IsCacheValid(const fs::path &src, const fs::path &cache);
void GenerateThumbnail(const fs::path &src, const fs::path &cache);
