build:
	rm -rf release/* && mkdir release
	qtdeploy build desktop
	qtdeploy -docker build windows_64_static
