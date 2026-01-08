.PHONY: all

test:
	@go test -race -covermode atomic -coverprofile=coverprofile ./...

detect_race:
	@go test -v -race
