
build:
	go build
	
run: build
	./demyst-todo status $(FLAGS)

test:
	go test -coverprofile=coverage.out ./...

test_html: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
