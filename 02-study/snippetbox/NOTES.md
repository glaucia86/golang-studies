# Notas 

Aqui segue algumas notas de estudo sobre o projeto `snippetbox`.

## Go Modules

Para criar um modulo em Go, vá até a pasta do projeto. No nosso caso, `snippetbox`. E dentro dele, execute o comando:

```bash
go mod init github.com/yourusername/snippetbox
```

Isso irá criar um arquivo `go.mod` na pasta do projeto. Esse arquivo é responsável por gerenciar as dependências do projeto.

## Lista de `Constants`

Podemos consultar a lista de `constants` disponíveis no pacote `net/http`, através do link: **[Constants](https://pkg.go.dev/net/http#pkg-constants)**

## Como estruturar um projeto em Go

Basta seguir a documentação oficial do Go, que explica como podemos estruturar um projeto.

**[Server project](https://go.dev/doc/modules/layout#server-project)**

### Projeto refatorado

O projeto foi refatorado. Assim sendo, para executar o projeto, basta executar o comando:

```bash
go run ./cmd/web/
```

Depois, basta acessar o endereço: **[http://localhost:4000](http://localhost:4000)** e aparecerá a página inicial do projeto.

**[Code Developed - commit](https://github.com/glaucia86/golang-studies/commit/8fb8ba757a301559576fda3056697f09db028913)**

