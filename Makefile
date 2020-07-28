build:
	rm -rf release/*
	CGO_ENABLED=0 go build -o release/sougou_to_gboard_mac main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/sougou_to_gboard_win_64.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o release/sougou_to_gboard_win_32.exe main.go
