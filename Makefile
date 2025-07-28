install:
	go build -o aurgo
	sudo mv aurgo /usr/local/bin/

test: 
	go test -v ./...