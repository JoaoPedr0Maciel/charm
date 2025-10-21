#!/bin/bash

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Informações do projeto
REPO="JoaoPedr0Maciel/charm"
BINARY_NAME="charm"

echo -e "${GREEN}🚀 Instalando Charm...${NC}"

# Detectar OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux*)
        OS="linux"
        ;;
    darwin*)
        OS="darwin"
        ;;
    *)
        echo -e "${RED}❌ Sistema operacional não suportado: $OS${NC}"
        exit 1
        ;;
esac

# Mapear arquitetura
case $ARCH in
    x86_64)
        ARCH="x86_64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}❌ Arquitetura não suportada: $ARCH${NC}"
        exit 1
        ;;
esac

echo -e "${YELLOW}📦 Detectado: $OS $ARCH${NC}"

# Pegar a última versão
echo -e "${YELLOW}🔍 Buscando última versão...${NC}"
LATEST_VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}❌ Não foi possível encontrar a última versão${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Última versão: $LATEST_VERSION${NC}"

# Construir nome do arquivo
if [ "$OS" = "darwin" ]; then
    OS_TITLE="Darwin"
elif [ "$OS" = "linux" ]; then
    OS_TITLE="Linux"
fi

FILE_NAME="${BINARY_NAME}_${LATEST_VERSION#v}_${OS_TITLE}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$FILE_NAME"

echo -e "${YELLOW}⬇️  Baixando $FILE_NAME...${NC}"

# Criar diretório temporário
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Baixar arquivo
if ! curl -fsSL "$DOWNLOAD_URL" -o "$FILE_NAME"; then
    echo -e "${RED}❌ Erro ao baixar: $DOWNLOAD_URL${NC}"
    exit 1
fi

# Extrair
echo -e "${YELLOW}📂 Extraindo...${NC}"
tar -xzf "$FILE_NAME"

# Determinar diretório de instalação
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
elif [ -w "$HOME/.local/bin" ]; then
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
else
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
fi

# Instalar
echo -e "${YELLOW}📥 Instalando em $INSTALL_DIR...${NC}"

if [ -w "$INSTALL_DIR" ]; then
    mv "$BINARY_NAME" "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    echo -e "${YELLOW}🔐 Permissão necessária para instalar em $INSTALL_DIR${NC}"
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

# Limpar
cd -
rm -rf "$TMP_DIR"

echo -e "${GREEN}✨ Charm instalado com sucesso!${NC}"
echo ""
echo -e "${YELLOW}📝 Uso:${NC}"
echo -e "  ${GREEN}charm get https://api.github.com/users/github${NC}"
echo ""

# Verificar se está no PATH
if ! command -v "$BINARY_NAME" &> /dev/null; then
    echo -e "${YELLOW}⚠️  O diretório $INSTALL_DIR não está no seu PATH${NC}"
    echo -e "${YELLOW}   Adicione ao seu ~/.bashrc ou ~/.zshrc:${NC}"
    echo -e "   ${GREEN}export PATH=\"\$PATH:$INSTALL_DIR\"${NC}"
    echo ""
fi

echo -e "${GREEN}🎉 Pronto para usar!${NC}"

