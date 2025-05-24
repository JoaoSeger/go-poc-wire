# Documentação do Projeto Wire

## O que é este projeto?

Este é um projeto de exemplo que demonstra o uso do Google Wire para injeção de dependência em Go. O projeto implementa um simples serviço de usuário com uma arquitetura de camadas (domain, application, infrastructure) seguindo princípios de Clean Architecture.

## Requisitos

- Go 1.21 ou superior
- Wire (`go install github.com/google/wire/cmd/wire@latest`)

## Como executar o projeto

1. Clone o repositório:

2. Execute o gerador do Wire (caso precise fazer alterações):
```
cd cmd
go generate
```

3. Execute a aplicação:
```
go run .
```

Observação: É importante executar `go run .` e não `go run main.go`, pois o Wire utiliza tags de build para gerenciar qual implementação do `InitializeUserService` será usada durante a compilação.

## Por que usar o Wire?

O Wire é uma ferramenta de injeção de dependência para Go desenvolvida pelo Google. Algumas vantagens de usar o Wire:

1. **Injeção de dependência em tempo de compilação**: Diferente de frameworks que resolvem dependências em tempo de execução, o Wire gera código Go que é compilado junto com sua aplicação.

2. **Sem reflexão**: O Wire não usa reflexão, o que torna o código mais rápido e seguro.

3. **Facilidade de manutenção**: Com o Wire, as dependências são explícitas, facilitando a compreensão do fluxo de código.

4. **Testabilidade**: Facilita a substituição de implementações reais por mocks durante testes.

## Estrutura do projeto

```
wire/
├── cmd/
│   ├── main.go         # Ponto de entrada da aplicação
│   ├── wire.go         # Definição dos provedores para o Wire
│   └── wire_gen.go     # Código gerado pelo Wire (não editar manualmente)
└── internal/
    └── user/
        ├── domain/
        │   └── entity.go       # Entidades e interfaces de domínio
        ├── application/
        │   └── service.go      # Serviços da aplicação
        └── infrastructure/
            └── repository.go   # Implementações de repositórios
```

## Como funciona?

1. O arquivo `wire.go` define os "providers" - funções que criam instâncias dos componentes necessários.
2. Quando executamos o comando `wire`, o código em `wire_gen.go` é gerado automaticamente.
3. O código gerado implementa a função `InitializeUserService()` com a lógica correta para instanciar todos os componentes e suas dependências.
4. A aplicação principal usa `InitializeUserService()` para obter uma instância configurada do serviço de usuário.

Esta abordagem permite que você mantenha um código limpo, organizado e fácil de testar, ao mesmo tempo que evita a complexidade de gerenciar dependências manualmente. 