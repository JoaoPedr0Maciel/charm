.PHONY: build install clean test run help

# Variáveis
BINARY_NAME=charm
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR=./bin
MAIN_PATH=./main.go

# Cores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

## help: Mostra esta mensagem de ajuda
help:
	@echo "$(GREEN)Comandos disponíveis:$(NC)"
	@echo "  $(YELLOW)make build$(NC)       - Compila o projeto"
	@echo "  $(YELLOW)make install$(NC)     - Instala o binário localmente"
	@echo "  $(YELLOW)make clean$(NC)       - Remove arquivos compilados"
	@echo "  $(YELLOW)make test$(NC)        - Roda os testes"
	@echo "  $(YELLOW)make run$(NC)         - Compila e executa"
	@echo "  $(YELLOW)make release$(NC)     - Cria release com GoReleaser (requer tag)"
	@echo "  $(YELLOW)make snapshot$(NC)    - Cria build snapshot sem publicar"

## build: Compila o projeto
build:
	@echo "$(GREEN)🔨 Compilando $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)✅ Binário criado em $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

## install: Instala o binário localmente
install:
	@echo "$(GREEN)📦 Instalando $(BINARY_NAME)...$(NC)"
	@go install -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN_PATH)
	@echo "$(GREEN)✅ $(BINARY_NAME) instalado com sucesso!$(NC)"

## clean: Remove arquivos compilados
clean:
	@echo "$(YELLOW)🧹 Limpando arquivos...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -rf dist/
	@go clean
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

## test: Roda os testes
test:
	@echo "$(GREEN)🧪 Rodando testes...$(NC)"
	@go test -v ./...

## run: Compila e executa
run: build
	@echo "$(GREEN)🚀 Executando $(BINARY_NAME)...$(NC)"
	@$(BUILD_DIR)/$(BINARY_NAME)

## release: Cria release com GoReleaser
release:
	@if [ -z "$(shell git describe --tags --exact-match 2>/dev/null)" ]; then \
		echo "$(YELLOW)⚠️  Nenhuma tag encontrada. Crie uma tag primeiro:$(NC)"; \
		echo "  git tag -a v0.1.0 -m 'Release v0.1.0'"; \
		echo "  git push origin v0.1.0"; \
		exit 1; \
	fi
	@echo "$(GREEN)📦 Criando release...$(NC)"
	@goreleaser release --clean

## snapshot: Cria build snapshot sem publicar
snapshot:
	@echo "$(GREEN)📸 Criando snapshot build...$(NC)"
	@goreleaser build --snapshot --clean
	@echo "$(GREEN)✅ Binários criados em dist/$(NC)"

## dev: Compila e roda em modo desenvolvimento
dev:
	@echo "$(GREEN)🔧 Modo desenvolvimento...$(NC)"
	@go run $(MAIN_PATH)

.DEFAULT_GOAL := help

