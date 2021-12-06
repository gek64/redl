# REDL (Release Downloader)

- Support github now, support more in the future
- Written in golang 

## Usage
```
Usage:
  redl [Options]

Options:
  -r  <repo>    : set repo
  -p  <part>    : set the search part of the file name to be downloaded
  -o  <output>  : set output file
  -h            : show help
  -v            : show version

Example:
  1) redl -r "gek64/redl" -p "windows-amd64"
  2) redl -r "gek64/redl" -p "windows" "amd64"
  3) redl -r "gek64/redl" -p "windows-amd64" -o "./release-downloader-windows-amd64.exe"
  4) redl -h
  5) redl -v
```

## Build
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/redl.git

cd redl

go build -v -trimpath -ldflags "-s -w"
```

## QA

### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This report occurred after `Windows 10 21h2`. This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. For more information, see https://go.dev/doc/faq#virus

### Q: Why should I clone `https://github.com/gek64/gek.git` before building
A: I don’t want the project to depend on a certain cloud service provider, and this is also a good way to avoid network problems.


## License

**GNU Lesser General Public License v2.1**

See `LICENSE` for details
