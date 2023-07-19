# Golang Studies :pencil:

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

Golang Expert from [LINUXtips](https://www.youtube.com/watch?v=YwmMFu0yEvw)

## Install

download tar file

from https://go.dev/doc/install
```
rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz
```

Ponto de entrada `main.go`.

rodar
```
go run main.go
```

criar pacote?
```
go mod init github.com/mateusoliveira43/golang-studies
go mod tidy
```

PascalCase -> publico (não consigo acessar de outro pacote se não for assim)

camelCase
