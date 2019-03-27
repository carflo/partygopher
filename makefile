build:
	cd src && go install
databind:
	go-bindata -prefix ./frames/gopherframes_txt/ -o ./src/bindata.go ./frames/gopherframes_txt/