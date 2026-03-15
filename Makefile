# Variabel
BINARY_DIR=bin
GO_FILES=$(shell find cmd -name "main.go")

.PHONY: all help clean test

help: ## Menampilkan daftar perintah yang tersedia
	@echo "Perintah yang tersedia:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## --- RUN COMMANDS ---

run-array-add: ## Menjalankan latihan Array Addition
	go run cmd/arrays/addition/main.go

run-array-rot: ## Menjalankan latihan Array Rotation
	go run cmd/arrays/rotation/main.go

run-num-fib: ## Menjalankan latihan Fibonacci
	go run cmd/numbers/fibonacci/main.go

run-api-fetch: ## Menjalankan latihan JSON Fetch API
	go run cmd/api/json_fetch/main.go

## --- UTILS ---

test: ## Menjalankan semua unit test di folder internal
	go test -v ./internal/...

tidy: ## Merapikan dependencies (go mod tidy)
	go mod tidy

fmt: ## Melakukan formatting kode Go
	go fmt ./...

clean: ## Menghapus file binary yang terkompilasi
	rm -rf $(BINARY_DIR)
	@echo "Folder bin telah dibersihkan."

build-all: ## Mengkompilasi semua main.go menjadi binary (opsional)
	@mkdir -p $(BINARY_DIR)
	@for file in $(GO_FILES); do \
		target_name=$$(dirname $$file | sed 's|cmd/||' | tr '/' '-'); \
		go build -o $(BINARY_DIR)/$$target_name $$file; \
		echo "Built: $(BINARY_DIR)/$$target_name"; \
	done