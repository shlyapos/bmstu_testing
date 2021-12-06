go-test: build-mocks test-all

run:
	go run .

test-all:
	go test .\test\...

build-mocks:
	mockery --dir .\app\repo\ --all