install:
	go build -o aurgo ./cmd/aurgo
	sudo mv aurgo /usr/local/bin/