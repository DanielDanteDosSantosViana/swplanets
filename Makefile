test:
		go test $$(go list ./... | grep -v /vendor/)

fmt:
		go fmt $$(go list ./... | grep -v /vendor/)

run:
		go run cmd/swplanetsd/main.go

.PHONY: test fmt
