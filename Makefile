.PHONY: dep lint test bench godoc

dep:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run

test:
	go test -mod=vendor -gcflags=all=-l $(shell go list ./... | grep -v mock | grep -v docs) -covermode=count -coverprofile .coverage.cov
	go tool cover -func=.coverage.cov

bench:
	go test -run=nonthingplease -benchmem -bench=. $(shell go list ./... | grep -v /vendor/)

godoc:
	echo "http://127.0.0.1:6060"
	godoc -http=127.0.0.1:6060 -goroot="."
