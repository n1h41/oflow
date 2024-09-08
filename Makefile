build:
	@go build -o build/oflow-server cmd/main.go

test:
	@go test -v ./...

go:
	@air --build.cmd "go build -o build/oflow-server cmd/main.go" --build.bin "build/oflow-server"

debug:
	@dlv debug --headless --api-version=2 --listen=127.0.0.1:38697 ./cmd/main.go
