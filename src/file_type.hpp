#pragma once

#include <filesystem>

enum class FileType {
	Error,
	Directory,
	Image,
	Markdown,
	PDF,
	Text,
	Video,
};

FileType DetermineFileType(const std::filesystem::path &file);
