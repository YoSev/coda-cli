test:
	go test -race ./...

build+darwin+amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-darwin-amd64 main.go
build+darwin+arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-darwin-arm64 main.go
build+linux+amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-linux-amd64 main.go
build+linux+arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-linux-arm64 main.go
build+windows+amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-windows-amd64.exe main.go
build+windows+arm64:
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w -extldflags "-static"" -o _bin/coda-windows-arm64.exe main.go

build+all: clean build+darwin+amd64 build+darwin+arm64 build+linux+amd64 build+linux+arm64 build+windows+amd64 build+windows+arm64

clean:
	rm -rf _bin
	mkdir _bin