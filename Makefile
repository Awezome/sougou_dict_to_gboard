build:
	fyne-cross darwin --env GOPROXY=https://goproxy.cn -arch=amd64
	fyne-cross windows --env GOPROXY=https://goproxy.cn -arch=*