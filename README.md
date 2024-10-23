# Release Downloader

- Download release file with or without specified names from GitHub, Gitlab and Sourceforge
- Automatically select the latest version or manually select by tag name
- Get the release file download link and use it with other download tools, such as curl, wget, aria2, etc.
- Have a built-in file downloader to download file if no other file downloader installed

## Usage

### Github

```sh
# download latest release
redl -gh "obsproject/obs-studio" -p "OBS-Studio" -p ".zip" -ep "PDBs.zip"

# download release by tag
redl -gh "obsproject/obs-studio" -t "30.2.2" -p "OBS-Studio" -p ".zip" -ep "PDBs.zip"
```

### Gitlab

```sh
# download latest release
redl -gl "36189" -p ".apk" -ep ".asc"

# download release by tag
redl -gl "36189" -t "1.18.0" -p ".apk" -ep ".asc"
```

### SourceForge

```sh
# download release by rss url
redl -sf "https://sourceforge.net/projects/mpv-player-windows/rss?path=/64bit" -p "x86_64" -p ".7z"
```

### Use with other download tools

```sh
# output download link to stdout without download file
redl -gl "36189" -p ".asc" -nd

# download use other download tool(aria2 curl wget etc)
aria2c "$(redl -gl '36189' -p '.asc' -nd)"
curl -LOJ "$(redl -gl '36189' -p '.asc' -nd)"
wget --content-disposition "$(redl -gl '36189' -p '.asc' -nd)"

# use with other download tools in Windows bat file
for /f "tokens=*" %%i in ('redl -gl 36189 -p .asc -nd') do (aria2c -x 4 %%i)
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/redl https://github.com/gek64/redl/releases/latest/download/redl-linux-amd64
chmod +x /usr/local/bin/redl

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/redl https://github.com/gek64/redl/releases/latest/download/redl-freebsd-amd64
chmod +x /usr/local/bin/redl
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/redl.git
cd redl
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
