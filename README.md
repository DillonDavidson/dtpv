# dtpv

Dillon's Terminal PreViewer - a terminal-based file previewer for text, images, videos, and markdown, designed to integrate with [lf](https://github.com/gokcehan/lf).

## Dependencies

- [atool](https://www.nongnu.org/atool/) - preview archive files
- [chafa](https://github.com/hpjansson/chafa/) - renders images/video thumbnails in the terminal
- [bat](https://github.com/sharkdp/bat) - syntax-highlighted text previews
- [ffmpegthumbnailer](https://github.com/dirkvdb/ffmpegthumbnailer) - video thumbnail generation
- [glow](https://github.com/charmbracelet/glow) - markdown previews
- [poppler](https://poppler.freedesktop.org/) - pdf previews with pdftoppm

## Building

```bash
git clone https://github.com/DillonDavidson/dtpv
cd dtpv
go install
```

## Usage

Add the following to your `~/.config/lf/lfrc`:
```bash
set previewer dtpv
```

## License

dtpv is licensed under the GNU GPL Version 3 license. See [LICENSE](https://github.com/DillonDavidson/dtpv/blob/master/LICENSE) for more information.
