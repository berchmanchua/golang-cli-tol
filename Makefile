.PHONY: build clean run

build:
	go build -o bin/tol cmd/tol/main.go

clean:
	rm -rf bin/

run:
	go run cmd/tol/main.go
