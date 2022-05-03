# harbourlivingapi

![Build Workflow](https://github.com/BigListRyRy/harbourlivingapi/actions/workflows/ci.yaml/badge.svg)


### Setup

- Install the Go version specified in go.mod
- Add the details required in api/app.env

### Start the Service Locally with make command

```make postgres```

```make createdb```

```make migrateup```

```make start```


### Run Tests
```make test```

### Technology
- GraphQL
- Postgres Database 
- Clean Architecture 
- Swagger 


### Database Entity Diagram
[ER Diagram](https://dbdiagram.io/d/612d650b825b5b0146eb97b0)

### Build and Push Docker Image
``docker build``

### Create new migration 
 migrate create -ext sql -dir db/migration -seq rating

###
```     latitude: Float
        longitude: Float
        miles: Float  ```