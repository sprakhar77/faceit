## Faceit User service 

### Hexagonal Architecture (Ports & Adapter)

[Wiki: Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

###### Architecture benefits:
- loosely coupled
- testable
- flexible (e.g: can change database/queues easily from Postgres to MongoDB or Kafka to RabbitMQ etc)

More information about this [here](https://dzone.com/articles/hexagonal-architecture-what-is-it-and-how-does-it)


###### How to run ?

Note: If you have Postgres already running you should turn it off 

`docker-compose up --build`

or

`docker-compose up`

## API Documentation

### Creates a new user

```http
POST /users/
```

Body
```json
{
  "first_name": "Prakhar",
  "last_name": "Ssdfrivastava",
  "nickname": "soasasmename",
  "password": "aaa",
  "email": "aaa@z.com",
  "country": "ABC",
}
```

### Update a user

```http
PUT /users/:user_id
```

Body
```json
{
  "first_name": "Prakhar",
  "last_name": "Ssdfrivastava",
  "nickname": "soasasmename",
  "password": "aaa",
  "email": "aaa@z.com",
  "country": "ABC",
}
```

### Get a user

```http
GET /users/:user_id
```

Body
```json
{
  "id": 2,
  "first_name": "Prakhar",
  "last_name": "Ssdfrivastava",
  "nickname": "soasasmename",
  "password": "aaa",
  "email": "aaa@z.com",
  "country": "ABC",
  "created_at": "2022-01-14T11:59:05.492625Z",
  "updated_at": "2022-01-14T12:00:20.317411Z"
}
```

### Get a list of users

```http
GET /users?country=ABC&limit=10&offset=3
```

Body
```json
[
    {
      "id": 2,
      "first_name": "Prakhar",
      "last_name": "Ssdfrivastava",
      "nickname": "soasasmename",
      "password": "aaa",
      "email": "aaa@z.com",
      "country": "ABC",
      "created_at": "2022-01-14T11:59:05.492625Z",
      "updated_at": "2022-01-14T12:00:20.317411Z"
    }
]
```
### Delete a user

```http
DELETE /users/:user_id
```
