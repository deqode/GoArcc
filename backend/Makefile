get-linter:
	go get github.com/mgechev/revive

run-linter:
	revive -formatter friendly ./...

get-fmt:
	go get fmt

run-fmt:
	go fmt ./...

run-all-linter: get-linter get-fmt run-linter run-fmt
