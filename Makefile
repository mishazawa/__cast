include .env
export

build:
	go build -o bin/$(BASE) .

run:
	go run .

clean:
	go clean -testcache

test: clean
	go test ./...

commit:
	gofmt -w .
	git add .
	git commit --allow-empty
