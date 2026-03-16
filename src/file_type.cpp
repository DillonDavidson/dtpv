#include "file_type.hpp"

#include <unordered_map>

FileType DetermineFileType(const std::filesystem::path &file)
{
	static const std::unordered_map<std::string, FileType> extensionMap = {
	    {".pdf", FileType::PDF},   {".md", FileType::Markdown}, {".mkv", FileType::Video},
	    {".mp4", FileType::Video}, {".jpg", FileType::Image},   {".jpeg", FileType::Image},
	    {".png", FileType::Image}, {".webp", FileType::Image},
	};

	std::string lowerExt = file.extension();
	std::transform(lowerExt.begin(), lowerExt.end(), lowerExt.begin(),
	               [](unsigned char c) { return std::tolower(c); });

	auto it = extensionMap.find(lowerExt);
	return it != extensionMap.end() ? it->second : FileType::Error;
}
