build:
	go build -o release/sougou_to_gboard src/*
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/sougou_to_gboard.exe src/*
