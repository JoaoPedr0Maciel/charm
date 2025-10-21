# 🚀 Como Fazer uma Release

## Primeira Release (v0.1.0)

### 1. Certifique-se de que tudo está commitado

```bash
git status
git add .
git commit -m "feat: configuração inicial com instalação automática"
```

### 2. Envie para o GitHub

```bash
git push origin main
```

### 3. Crie uma tag

```bash
git tag -a v0.1.0 -m "Release v0.1.0 - Primeira versão"
git push origin v0.1.0
```

### 4. Aguarde o GitHub Actions

Após fazer o push da tag:
- Vá até: https://github.com/JoaoPedr0Maciel/charm/actions
- O workflow "Release" será executado automaticamente
- Aguarde cerca de 2-5 minutos

### 5. Verifique a Release

- Acesse: https://github.com/JoaoPedr0Maciel/charm/releases
- Você verá os binários para:
  - Linux (amd64, arm64)
  - Windows (amd64)
  - macOS (amd64, arm64)

## 🎉 Pronto!

Agora qualquer pessoa pode instalar com:

### Linux/Mac
```bash
curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/charm/main/install.sh | bash
```

### Windows
Baixar o `.zip` da página de releases

### Go Install
```bash
go install github.com/JoaoPedr0Maciel/charm@latest
```

## Próximas Releases

Para releases futuras, apenas repita os passos 1-5 incrementando a versão:

```bash
# Bugfixes
git tag -a v0.1.1 -m "Release v0.1.1 - Correção de bugs"

# Novas features
git tag -a v0.2.0 -m "Release v0.2.0 - Nova feature X"

# Breaking changes
git tag -a v1.0.0 -m "Release v1.0.0 - Versão estável"

git push origin <tag-name>
```

## Testar Localmente Antes de Publicar

Para testar como ficará a release sem publicar:

```bash
make snapshot
```

Isso criará binários em `dist/` para você testar.

## Troubleshooting

### Erro: "GoReleaser not found"

Se o GitHub Actions falhar, certifique-se de que:
1. O repositório é público ou você tem GitHub Pro/Team
2. As permissões do workflow estão corretas (já configuradas no `.github/workflows/release.yml`)

### Erro ao baixar com install.sh

Se o script de instalação falhar:
1. Verifique se a tag foi criada corretamente
2. Aguarde alguns minutos para o GitHub Actions terminar
3. Verifique se os assets foram publicados na release

## Dicas

- Use [Conventional Commits](https://www.conventionalcommits.org/) para mensagens de commit
- Crie um CHANGELOG.md para documentar mudanças
- Teste sempre com `make snapshot` antes de fazer release oficial

