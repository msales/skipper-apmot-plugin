build:
	@GO111MODULE=on go build -buildmode=plugin -o apmot.so plugin.go
