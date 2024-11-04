build: clean build/gotex
	@echo "build created in: $(shell pwd)/build/gotex"

build/gotex:
	@go build -o build/gotex .

clean:
	@rm -rf build
	@rm -f *.tex *.pdf