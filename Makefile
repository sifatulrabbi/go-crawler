run:
	go run ./cmd/crawler/main.go
buildx86:
	mkdir build; go build -o build/crawer ./cmd/crawler/main.go
.PHONY: run, buildx86
