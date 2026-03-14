#include "file_extension.hpp"

#include <string>
#include <unistd.h>
#include <vector>

int main(int argc, char *argv[])
{
	if (argc < 2) {
		return 0;
	}

	std::string file = argv[1];

	ExtensionType ext = DetermineExtensionType(file);

	std::vector<std::string> args;
	switch (ext) {
	case ExtensionType::Directory:
		break;
	case ExtensionType::Image:
		break;
	case ExtensionType::Markdown:
		args = {"glow", file};
	case ExtensionType::PDF:
		break;
	case ExtensionType::Text:
		args = {"bat", "--color=always", file};
		break;
	case ExtensionType::Video:
		break;
	}

	std::vector<char *> cargs;
	cargs.reserve(args.size() + 1);
	for (auto &s : args) {
		cargs.push_back(s.data());
	}

	cargs.push_back(nullptr);

	execvp(cargs[0], cargs.data());

	perror("execvp failed");
	return 1;
}
