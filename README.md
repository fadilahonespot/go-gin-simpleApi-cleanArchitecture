# Simple API With Gin and Gorm Framework

CRUD simple RESTFUL API with routing in using gin framework and ORM database with Gorm auto migrate. Database using mysql. Check config/config.go for configurate database. just create a database with the name simple_api in MySQL and then run this API program 

## Installation

- configure database in folder config/config.go
- open the program via the terminal.
- run the program by typing the command

```bash
go run main.go
```

## Routing in API
Testing in localhost

- GET Method 
```
localhost:7861/person
```
- GET Method by id
```
localhost:7861/person/id
```
```
example
localhost:7861/person/1
```
- POST METHOD
```
localhost:7861/person
```
Input body
```
{
  "first_name": "Agung",
  "last_name": "Sanjaya"
}
```
- PUT Method 
```
localhost:7861/person
```
Input body #input ID will be used to update the data
```
{
  "id": 4,
  "first_name": "Novan",
  "last_name": "Gunawan"
}
```

- DELETE METHOD by id
```
localhost:7861/person/id
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
