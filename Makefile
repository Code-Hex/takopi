.PHONY: all wasm wasm_exec
all: wasm wasm_exec
wasm:
	GOOS=js GOARCH=wasm go build -trimpath -o takopi.wasm ./cmd/wasm

wasm_exec:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .