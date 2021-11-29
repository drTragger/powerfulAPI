run: build
	./api -path configs/api.toml
build:
	go build -v ./cmd/api/