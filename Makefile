build:
	go build -o kuchihira-bot

build-debug:
	go build -o kuchihira-bot_debug

debug: build-debug
	./kuchihira-bot_debug post --debug

daemondebug: build
	./kuchihira-bot_debug daemon --debug
