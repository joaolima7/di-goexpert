# Guia Rápido para Usar Google Wire em Go
O Google Wire é uma ferramenta poderosa para injeção de dependências em aplicações Go. Ele ajuda a organizar e gerenciar as dependências de forma eficiente, facilitando a manutenção e escalabilidade do código. A seguir, apresentamos os principais pontos para começar a usar o Google Wire.
## Instalação
Para instalar o Google Wire, execute o seguinte comando no terminal:

```bash
go install github.com/google/wire/cmd/wire@latest
```
## Uso Básico
1. **Criação do Arquivo Wire**: Crie um arquivo chamado `wire.go` no seu projeto. Este arquivo conterá as definições de injeção de dependências.
2. **Registro de Sets**: Utilize `wire.NewSet` para agrupar as dependências que serão injetadas. Por exemplo, no seu caso, você criou um set para o repositório de produtos:

```go
var setRepositoryDependence = wire.NewSet(product.NewProductRepository, wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))
```
3. **Anotações no Arquivo**: Utilize as anotações `//go:generate` e `//go:build` para indicar como o Wire deve gerar o código. Por exemplo:

```go
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject
```
4. **Definição de Injetores**: Crie funções que retornam as instâncias das dependências. No seu exemplo, você definiu um injetor para o caso de uso do produto:

```go
func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
    productRepository := product.NewProductRepository(db)
    productUseCase := product.NewProductUseCase(productRepository)
    return productUseCase
}
```
5. **Geração de Código**: Após definir suas dependências e injetores, execute o comando para gerar o código necessário:

```bash
wire
```
6. **Execução do Código**: Finalmente, você pode executar seu programa com o comando:

```bash
go run main.go wire_gen.go
```