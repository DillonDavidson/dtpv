#include "file_extension.hpp"

ExtensionType DetermineExtensionType(const std::filesystem::path &file)
{
	if (file.extension() == ".pdf") {
		return ExtensionType::PDF;
	}

	if (file.extension() == ".md") {
		return ExtensionType::Markdown;
	}

	if (file.extension() == ".mkv") {
		return ExtensionType::Video;
	}

	return ExtensionType::Text;
}
