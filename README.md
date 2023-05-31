
# API GO

Proyecto de simulación de billetera electronica con uso de la api Background Check de Truora 

## Tech Stack

**Server:** Go

**Database:** Postgres

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`

`DB_PORT`

`DB_NAME`

`DB_USER`

`DB_PASSWORD`

`Truora_API_Key`

## API Reference

#### Get all wallets

```http
  GET /api/v1/wallets
```

#### Get One wallet

```http
  GET /api/v1/wallets/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of wallet to fetch |

#### Post wallet

```http
  POST /api/v1/wallets
```
`Body Request`

```json
{
    "national_id": "12345678",
    "country": "PE",
    "type": "person",
    "user_authorized": true
}

```

#### Update wallet
```http
  PUT /api/v1/wallets/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of wallet to fetch |

`Example Body Request`

```json
{
    "haveCard":true
}
```

#### Delete item
```http
  DELETE api/v1/wallets/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of wallet to fetch |

#### Status wallet
```http
  GET /api/v1/wallets/status?personId={personId}
```
| Query Params | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `personId`      | `string` | **Required**. Personal identification associated with a wallet to fetch |

#### Post Transaction

```http
  POST /api/v1/transactions
```
`Body Request`

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `source`      | `int` | **Required**. Id de la wallet que realizará el envio |
| `destiny`      | `int` | **Required**. Id de la wallet que recibirá el envio |
| `type`      | `int` | **Required**. Tipo de envio: 1 para deposito, 2 para retiro |

```json
{
    "amount": 30,
    "source": 9,
    "destiny": 10,
    "type": 1
}

```

## Run Locally

Start the server

```bash
  go run main.go
```
## Authors

- [Raisa Orellana](https://github.com/Raisa320)


