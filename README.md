# Release Downloader

- Download Release from GitHub and Sourceforge

## Usage

```
# download MPV windows from sourceforge
redl -sf "https://sourceforge.net/projects/mpv-player-windows/rss?path=/64bit" -p "x86_64" -p ".7z"

# download OBS windows from github
redl -gh "obsproject/obs-studio" -p "OBS-Studio" -p ".zip" -ep "pdbs"
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
