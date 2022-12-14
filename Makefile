test:mocks
	go test ./...

run:
	go run main.go

cover:mocks
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

mocks:
	mockery --dir databases --all --output ./databases/mocks
