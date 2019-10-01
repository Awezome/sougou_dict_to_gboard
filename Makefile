build:
	go build -o release/sougou_to_gboard_mac src/*
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/sougou_to_gboard_win.exe src/*
