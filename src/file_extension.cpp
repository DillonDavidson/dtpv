#include "file_extension.hpp"

ExtensionType DetermineExtensionType(const std::string &file)
{
	if (file.ends_with(".pdf")) {
		return ExtensionType::PDF;
	}

	if (file.ends_with(".md")) {
		return ExtensionType::Markdown;
	}

	return ExtensionType::Text;
}
