generate:
	go generate && go build -o main

setup:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
