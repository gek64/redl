# Release Downloader

- Support GitHub(rope mode) and Sourceforge(rss mode)

## Usage

```
Usage:                                                                 
  redl {-r rope | -rss rss_link} [Options] -p [part1, part2, par3, ...]

Args:
  -r   <repo>    : set repo
  -p   <part>    : set the search part of the file name to be downloaded
  -rss <rss>     : set rss link

Options:
  -o   <output>  : set output file

Other:
  -h             : show help
  -v             : show version

Example:
  1) redl -r "gek64/redl" -p "windows-amd64"
  2) redl -r "gek64/redl" -p "windows" "amd64"
  3) redl -r "gek64/redl" -o "./release-downloader-windows-amd64.exe" -p "windows-amd64" ".exe"
  4) redl -rss "https://sourceforge.net/projects/mpv-player-windows/rss?path=/64bit" -p "x86_64" ".7z"
  5) redl -h
  6) redl -v
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/redl.git
cd redl
go build -v -trimpath -ldflags "-s -w"
```

## QA

### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`

A: This application does not contain any malware, backdoors, and advertisements, all released files are build by GitHub
actions. For more information, see https://go.dev/doc/faq#virus

## License

- **GPL-3.0 License**
- See `LICENSE` for details
