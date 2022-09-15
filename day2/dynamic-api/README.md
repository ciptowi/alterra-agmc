# REST API example application

This is a example of an application that provides REST
API to existing Books model.

## Get all users
### Request
`GET /users`

    curl -i -H 'Accept: application/json' -X GET http://localhost:5000/users

### Response

    HTTP/1.1 200 Ok
    Status: 200 Ok
    Connection: close
    Content-Type: application/json

    {"message": "Success", "data": []}


## Create new user
### Request
`POST /users`

    curl -i -H 'Accept: application/json' -X POST -d 'title=Foo&writer=Bob&isbn=EDS123' http://localhost:5000/users

### Response
    HTTP/1.1 201 Created
    Status: 201 Created
    Connection: close
    Content-Type: application/json

    {"message": "Create Success", "data": {"id": 1, "title": "Foo", "isbn": "EDS123", "writer": "Bob"}}

## Get users by id
### Request
`GET /users/:id`

    curl -i -H 'Accept: application/json' -X GET http://localhost:5000/users/1

### Response

    HTTP/1.1 200 Ok
    Status: 200 Ok
    Connection: close
    Content-Type: application/json

    {"message": "Success", "data": {"id": 1, "title": "Foo", "isbn": "EDS123", "writer": "Bob"}}

## Update user by id
### Request
`PUT /users/:id`

    curl -i -H 'Accept: application/json' -X PUT -d 'title=Changed&writer=Bob&isbn=EDS123' http://localhost:5000/users/1

### Response
    HTTP/1.1 201 Created
    Status: 201 Created
    Connection: close
    Content-Type: application/json

    {"message": "Update Success", "data": {"id": 1, "title": "Changed", "isbn": "EDS123", "writer": "Bob"}}

## Delete users by id
### Request
`DELETE /users/:id`

    curl -i -H 'Accept: application/json' -X DELETE http://localhost:5000/users/1

### Response

    HTTP/1.1 204 No Content
    Status: 204 No Content
    Connection: close
    Content-Type: application/json

