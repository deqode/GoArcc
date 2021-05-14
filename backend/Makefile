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
	richgo test -cover ./...
	mkdir -p cover
	ls cover -l
	go test -coverprofile=cover/coverage.out ./...
	ls cover -l
	go tool cover -html=coverage.out -o cover/coverage.html

build-api:
	go build -o bin/backend_api cmd/api/*.go
