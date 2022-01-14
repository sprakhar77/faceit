## Faceit User Service 

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


## What was not done/ could be extended
- I did not write tests (unit/integration) as it required much more time than I had at hand. Also most of validation tests are already automatically handled by the database (have a look at init.sql). I added some basic validation for email and encrypting the password before storing it in DB.
- I did not implement a concrete Kafka writer as this was out of scope. It could be easily plugged in to the UserPublisher port
- I wrote the database filter to filter on all properties but only exposed Country (via the API package) to external clients as exposing others properties was out of scope. However they can be easily added by exteniding the request ypes in api package, the database layer is already prepared to handle that.
- I did not use UUID to store ids instead I used auto-incremented int64 values which are very effective in postgres.

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
