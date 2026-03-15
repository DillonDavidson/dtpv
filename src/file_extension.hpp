#pragma once

#include <filesystem>

enum class ExtensionType {
	Directory = 0,
	Image = 1,
	Markdown = 2,
	PDF = 3,
	Text = 4,
	Video = 5,
};

ExtensionType DetermineExtensionType(const std::filesystem::path &file);
