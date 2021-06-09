clean:
	rm -f bin/*

generate:
	go generate ./...

piccolo: clean generate
	go build \
		-o bin/piccolo \
		cmd/piccolo/*

test: generate
	go test ./... -cover