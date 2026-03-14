#pragma once

#include <string>

enum class ExtensionType {
	Directory,
	Image,
	Markdown,
	PDF,
	Text,
	Video,
};

ExtensionType DetermineExtensionType(const std::string &file);
