# API de Filmes
Exemplo de uma API de filmes escrita em *Golang* usando os servicos [API Gateway](https://aws.amazon.com/pt/api-gateway/), [Lambda](https://aws.amazon.com/pt/lambda/) e [DynamoDB](https://aws.amazon.com/pt/dynamodb/).

## DynamoDB
Criar uma tabela chamada `movies` com os campos `id`, `name` e `image`:

![Detalhes da tabela movies](screenshots/dynamodb1.png)

![Itens da tabela movies](screenshots/dynamodb2.png)

## Lambda
Criar uma funcao Lambda chamada `movies`:

Para compilar e zipar os arquivos do projeto usar os comandos:

```
env GOOS=linux go build -o main
zip main.zip main
```

## API Gateway

Criar uma API com o recurso `movies` e os métodos `GET` e `POST`:

![API Gateway](screenshots/apigateway1.png)

Integracao com o Lambda:

![Integracao do API Gateway com o Lambda](screenshots/apigateway2.png)

## Testando
Recuperando lista de filmes:

![Recuperando lista de filmes](screenshots/postman1.png)

Inserindo um novo filme:

![Inserindo um novo filme](screenshots/postman2.png)

## Créditos
[https://medium.com/iq360/go-serverless-construindo-uma-api-usando-golang-e-aws-lambda-2a7d6a3019b9](https://medium.com/iq360/go-serverless-construindo-uma-api-usando-golang-e-aws-lambda-2a7d6a3019b9)