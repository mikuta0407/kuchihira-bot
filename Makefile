build:
	go build .

debug: build
	./kuchihira-bot post --debug
