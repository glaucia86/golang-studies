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

## Significado de `http.Handler` em Go

`http.Handler` é uma interface que define o método `ServeHTTP` que recebe um `http.ResponseWriter` e um `*http.Request` como argumentos.

Exemplo:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Mais detalhes, acesse o link: **[http.Handler](https://pkg.go.dev/net/http#Handler)**

## Significado de `http.ListenAndServe` em Go

`http.ListenAndServe` é uma função que recebe dois argumentos: um endereço de rede e um `http.Handler` e retorna um erro.

Exemplo:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
  // Hello World, the web server

  helloHandler := func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
  }

  http.HandleFunc("/hello", helloHandler)
  log.Fatal(http.ListenAndServe(":4000", nil))
}
```

## Lidando com environment variables em Go

Para lidar com environment variables em Go, basta importar o pacote `os` e utilizar a função `os.Getenv()`.

Exemplo:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
  // Hello World, the web server

  helloHandler := func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
  }

  http.HandleFunc("/hello", helloHandler)
  log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
```

## Exemplo de Closures for dependency injection

> observação: seguir o exemplo feito **[AQUI](https://gist.github.com/alexedwards/5cd712192b4831058b21)** e **[AQUI](https://www.alexedwards.net/blog/organising-database-access)**

## How to execute MySQL at WSL2?

Para executar o MySQL no WSL2, basta executar o seguinte comando:

```bash
sudo service mysql start
```

E depois, entrar com os dados de usuário e senha.

```bash
sudo mysql -u root -p
```

Em caso de usar, sem ser o root, basta executar o seguinte comando:

```bash
mysql -D snippetbox -u web -p
```

Estaremos usando o driver do MySQL para Go: **[mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql)**

## Shorthand single-record queries

Para fazer uma query de um único registro, podemos usar o método `QueryRow()`.

Exemplo:

* snippets.go (método `Get()`)

```go
func (m *SnippetModel) Get(id int) (Snippet, error) {
    var s Snippet
    
    err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return Snippet{}, ErrNoRecord
        } else {
             return Snippet{}, err
        }
    }

    return s, nil
}
```