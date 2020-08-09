build:
	rm -rf release/*
	go build -o release/sougou_to_gboard_mac main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui" -o release/sougou_to_gboard_win_64.exe main.go
