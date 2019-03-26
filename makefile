build:
	go build -o gopher .
databind:
	go-bindata -prefix ./frames/gopherframes_txt/ ./frames/gopherframes_txt/

