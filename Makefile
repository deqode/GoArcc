get-linter:
	go get github.com/mgechev/revive

get-richgo:
	go get -u github.com/kyoh86/richgo

get-dupl:
	go get -u github.com/mibk/dupl

get-errcheck:
	go get -u github.com/kisielk/errcheck

get-gocritic:
	go get -u github.com/go-critic/go-critic/cmd/gocritic

run-gocritic:
	gocritic check ./...

run-errcheck: get-errcheck
	errcheck ./...

run-dupl: get-dupl
	dupl */**.go



run-linter: get-linter
	revive -formatter friendly ./...

get-fmt:
	go get fmt

run-fmt:
	go fmt ./...

tidy:
	go mod tidy
	go mod vendor

# Run all the linters
run-all-linter: get-linter get-fmt run-linter run-fmt tidy

#Get the test coverage
test: get-richgo
	go test -cover ./...
	go test -coverprofile=./cover/coverage.out ./...
	cat cover/coverage.out
	go tool cover -html=./cover/coverage.out

build:
	go build -o bin/goarcc -buildmode pie ./cmd/api
