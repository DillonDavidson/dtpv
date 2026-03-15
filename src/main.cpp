#include "file_extension.hpp"
#include "markdown.hpp"
#include "text.hpp"
#include "video.hpp"

#include <string>
#include <unistd.h>
#include <vector>

int main(int argc, char *argv[])
{
	if (argc < 2) {
		return 0;
	}

	const fs::path file = argv[1];
	const std::string width = argc > 2 ? argv[2] : "80";
	const std::string height = argc > 3 ? argv[3] : "40";

	ExtensionType ext = DetermineExtensionType(file);

	std::vector<std::string> args;

	switch (ext) {
	case ExtensionType::Directory:
		break;
	case ExtensionType::Image:
		break;
	case ExtensionType::Markdown:
		args = BuildMarkdownCommand(width, file);
		break;
	case ExtensionType::PDF:
		break;
	case ExtensionType::Text:
		args = BuildTextCommand(width, file);
		break;
	case ExtensionType::Video:
		args = BuildVideoCommand(width, height, file);
		break;
	}

	std::vector<char *> cargs;
	cargs.reserve(args.size() + 1);
	for (auto &s : args) {
		cargs.push_back(s.data());
	}

	cargs.push_back(nullptr);

	// This fixes the bleed but kind of bugs lf, so I'm ignoring it for now
	// I leave this as a future problem for nerds
	// std::cout << "\033[H\033[2J" << std::flush;

	execvp(cargs[0], cargs.data());

	perror("execvp failed");
	return 1;
}
