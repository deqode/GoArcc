get-linter:
	go get github.com/mgechev/revive

run-linter:
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
test:
	go test -cover ./...
