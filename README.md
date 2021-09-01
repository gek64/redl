# REDL (Release Downloader)

- Support github now, support more in the future
- Written in golang 

## Usage
```
Version:
  1.00

Usage:
  redl [Options]

Options:
  -r  <repo>    : set repo
  -p  <part>    : set the search part of the file name to be downloaded
  -o  <output>  : set output file
  -h            : show help
  -v            : show version

Example:
  1) redl -r gek64/redl -p windows-amd64
  2) redl -r gek64/redl -p windows-amd64 -o ./release-downloader-windows-amd64.exe
  3) redl -h
  4) redl -v
```

## Build
### Example
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/redl.git

cd redl

go build -v -trimpath -ldflags "-s -w"
```
