module redl

go 1.17

require (
	gek_downloader v0.0.0
	gek_github v0.0.0
)

replace (
	gek_downloader => ../gek/gek_downloader
	gek_exec => ../gek/gek_exec
	gek_file => ../gek/gek_file
	gek_github => ../gek/gek_github
	gek_json => ../gek/gek_json
	gek_math => ../gek/gek_math
)
