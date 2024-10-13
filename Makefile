# 変数定義
GO = go
GOOS = js
GOARCH = wasm
GO_DIR = art/go
WASM_DIR = public/wasm

# デフォルトのアート名
ART_NAME ?= ruby_image

# ターゲット
.PHONY: all clean build

help: 
	@echo 'make go-build ART_NAME=hogehoge'

go-build: $(WASM_DIR)/wasm_exec.js
	cd $(GO_DIR) && GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o ../../$(WASM_DIR)/go_$(ART_NAME).wasm $(ART_NAME)/main.go

clean:
	rm -f $(WASM_DIR)/*.wasm

# wasm_exec.jsのコピー
$(WASM_DIR)/wasm_exec.js:
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js $(WASM_DIR)/

run:
	npm run dev