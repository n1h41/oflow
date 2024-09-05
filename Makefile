build:
	@go build -o build/oflow-server cmd/main.go

test:
	@go test -v ./...

go:
	@air --build.cmd "go build -o build/oflow-server cmd/main.go" --build.bin "build/oflow-server"
