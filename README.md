# api-social

## Dev

A API abaixo esta configurada com Dockerfile, para que ela funcione é necessario ter instalado o docker e GO abaixo temos o passa a passo para roda-la no 
Linux sem o uso de docker.
```bash
cd /src/router
go run *.go
```

## Rodar Aplicação
A API abaixo esta configurada com Dockerfile, para que ela funcione é necessario ter instalado o docker e GO abaixo temos o passa a passo para roda-la em qualquer sistema operacional usando docker.
```bash
docker buil -t social
docker run -p 5000:5000 social
```

## User
```bash
GET - /user/ 
POST - /user/
PUT - /user/{id}
DELETE - /user/{id}
```
## User Json
```JSON
{
  "name": "Nome",
  "email": "email@example.com",
  "password": "senha",
  "city": "Cidade",
  "state": "Estado",
  "coutry": "País",
  "cep": "CEP",
  "number": "Número",
  "account_billers": false
}

```
## Category
```bash
GET - /category/ 
POST - /category/
PUT - /category/{id}
```
## Category Json
```JSON
{
"Name": "name category",
}
```

## Comments

```bash
POST - /comment

```
## Comments Json


```json
{
  "idPost": "12345",
  "userComments": "NomeDoUsuario",
  "commentsText": "Texto do comentário"
}
```



## Posts

```bash
GET - /post/index
POST - /post/

```
## Posts Json


```json
{
  "title": "Título do Post",
  "titleSlug": "slug-do-titulo",
  "text": "Texto do Post",
  "name": "Nome do Usuário",
  "linkYoutube": "URL do YouTube",
  "category": "Categoria do Post",
  "image": {
    "imageOne": "URL da Imagem 1",
    "imageTwo": "URL da Imagem 2",
    "imageThree": "URL da Imagem 3",
    "imageFour": "URL da Imagem 4",
    "imageFive": "URL da Imagem 5"
  },
  "bannerAltText": "Texto alternativo do banner",
  "commentsQuantity": 10,
  "approved": true
}

```
## Login
Ao Realizar o Login vamos ter o retorno de um webtoken que será utilizado para fazer a autenticacao. 
```bash

POST -  /login/

```

## Login JSON
```JSON
{
  "email": "@example.com",
  "password": "password"
}

```

## JSON A Ser Retornado Pela Requisição Da Rota De Posts

```JSON
{
  "category": "Categoria do Post",
  "post": {
    "id": "5f343f2d95b9a86d157548c0",
    "title": "Título do Post",
    "titleSlug": "titulo-do-post",
    "text": "Texto do Post",
    "nameUser": "Nome do Usuário",
    "category": "Categoria do Post",
    "image": {
      "imageOne": "url-imagem-1",
      "imageTwo": "url-imagem-2",
      "imageThree": "url-imagem-3",
      "imageFour": "url-imagem-4",
      "imageFive": "url-imagem-5"
    },
    "bannerAltText": "Texto alternativo do banner",
    "linkYoutube": "url-youtube",
    "approved": true,
    "comments": [
      {
        "_id": "5f343f2d95b9a86d157548c1",
        "idPost": "12345",
        "userComments": "NomeDoUsuario",
        "commentsText": "Texto do comentário",
        "createdAt": "2021-09-01T10:30:00Z"
      },
      {
        "_id": "5f343f2d95b9a86d157548c2",
        "idPost": "12345",
        "userComments": "NomeDoUsuario2",
        "commentsText": "Outro comentário",
        "createdAt": "2021-09-02T12:45:00Z"
      }
    ],
    "createdAt": "2021-09-01T09:15:00Z",
    "updatedAt": "2021-09-02T14:20:00Z"
  }
}
```

## Json Post Cards Home


```JSON
[
	{
		"id": "649b1b7029df18c79939e5b2",
		"title": "texto",
		"titleSlug": "",
		"text": "Get the most out of your new inbox by quickly and easily marking all of your previously",
		"nameUser": "teste",
		"category": "teste",
		"image": {
			"imageOne": "http://localhost:8080/integration/?event=01GZICVD1V9RED",
			"imageTwo": "",
			"imageThree": "",
			"imageFour": "",
			"imageFive": ""
		},
		"bannerAltText": "",
		"createdAt": "2023-06-27T17:25:04.135Z"
	},
	{
		"id": "649b33ce33c467e7e1be5eb8",
		"title": "teste mateus",
		"titleSlug": "",
		"text": "Get the most out of your new inbox by quickly and easily marking all of your previously",
		"nameUser": "local",
		"category": "teste",
		"image": {
			"imageOne": "http://localhost:8080/integration/?event=01GZICVD1V9RED",
			"imageTwo": "",
			"imageThree": "",
			"imageFour": "",
			"imageFive": ""
		},
		"bannerAltText": "",
		"createdAt": "2023-06-27T19:09:02.824Z"
	}
]
```



