# Variables
wasmPath = "./static/miho.wasm"
servPath = "./initServer"

# Commands
clean:
	@if [ -d $(wasmPath) ]; then \
		rm -rf $(wasmPath); \
	fi
	@if [ -d $(servPath) ]; then \
		rm -rf $(servPath); \
	fi

build: clean
	clear
	@GOOS=js GOARCH=wasm go build -o $(wasmPath) ./wasm/wasm_main.go
	@go build -o $(servPath) ./main.go
	@$(servPath)

dev:
	clear
	@GOOS=js GOARCH=wasm go build -o $(wasmPath) ./wasm/wasm_main.go
	@go build -o ./tmp/main ./main.go