# ✨ Charm

Uma ferramenta CLI moderna e elegante para fazer requisições HTTP com saída formatada e colorida.

## 🚀 Instalação Rápida

### Linux / macOS 🐧 🍎

```bash
curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/charm/main/install.sh | bash
```

### Windows 🪟

1. Baixe o arquivo `.zip` da [última release](https://github.com/JoaoPedr0Maciel/charm/releases/latest)
2. Extraia o arquivo
3. Adicione o diretório ao PATH ou mova o `charm.exe` para um diretório que já esteja no PATH

### Via Go Install

Se você já tem Go instalado:

```bash
go install github.com/JoaoPedr0Maciel/charm@latest
```

## 📖 Uso

### GET Request

```bash
# Simples
charm get https://api.github.com/users/github

# Com Bearer token
charm get https://api.example.com/data --bearer seu-token-jwt

# Com Basic auth
charm get https://api.example.com/data --basic usuario:senha
```

### POST Request

```bash
# Com dados JSON
charm post https://api.example.com/users --data '{"name":"João","email":"joao@example.com"}'

# Com autenticação
charm post https://api.example.com/users \
  --bearer seu-token \
  --data '{"name":"João"}'
```

### PUT Request

```bash
charm put https://api.example.com/users/1 \
  --data '{"name":"João Pedro","email":"jp@example.com"}' \
  --bearer seu-token
```

### PATCH Request

```bash
charm patch https://api.example.com/users/1 \
  --data '{"active":true}' \
  --bearer seu-token
```

### DELETE Request

```bash
charm delete https://api.example.com/users/1 --bearer seu-token
```

### Ver Versão

```bash
charm version
```

## 🎨 Features

- ✨ Output colorido e formatado
- 📊 Exibição de status code com emojis
- ⏱️ Medição de tempo de resposta
- 📋 Visualização clara de headers
- 🎯 Suporte para autenticação (Bearer e Basic)
- 🌈 JSON formatado e colorido
- 🚀 Suporte completo para GET, POST, PUT, PATCH, DELETE
- 📦 Envio de dados JSON no body

## 🛠️ Desenvolvimento

### Requisitos

- Go 1.25 ou superior
- Make (opcional, mas recomendado)

### Comandos disponíveis

```bash
make build      # Compila o projeto
make install    # Instala localmente
make test       # Roda os testes
make run        # Compila e executa
make clean      # Limpa arquivos compilados
make snapshot   # Cria build local para todos os OS
```

### Build manual

```bash
git clone https://github.com/JoaoPedr0Maciel/charm.git
cd charm
go build -o charm
```

## 📦 Criar uma Release

1. Crie uma tag:
```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

2. O GitHub Actions automaticamente:
   - Compila para Linux, Windows e macOS
   - Cria os binários
   - Publica na página de releases

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se livre para abrir issues ou pull requests.

## 📝 Licença

Este projeto está sob a licença MIT.

## 🌟 Status do Projeto

Em desenvolvimento ativo. Mais features em breve!

---

Feito com ❤️ por [João Pedro Maciel](https://github.com/JoaoPedr0Maciel)

