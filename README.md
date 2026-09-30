# XM Golang Exercise - v22.0.0

### How to run
Run from the root directory of the project
```sh
docker-compose -f deploy/dev/docker-compose.yml up --build
```
Tested on OrbStack

### Examples
For getting jwt token
```sh
GET http://localhost:8000/auth/token
```
The resulting value is used below as <token>

#### Work with companies
**GET**

If company 01a0ee99-09b9-7679-90ed-574c05701342 exists
```sh
GET http://localhost:8000/v1/company/01a0ee99-09b9-7679-90ed-574c05701342
```

**CREATE**
```sh
POST http://localhost:8000/v1/company/
Authorization: Bearer <token>
Content-Type: application/json

{
"name": "CDF",
"description": "ERF",
"amountOfEmployees": 10,
"registered": true,
"type": "Cooperative"
}
```

**PATCH**
```sh
PATCH http://localhost:8000/v1/company/
Authorization: Bearer <token>
Content-Type: application/json

{
"name": "CDFRE",
"amountOfEmployees": 10,
"registered": false
}
```

**DELETE**
```sh
DELETE http://localhost:8000/v1/company/01a0ee99-09b9-7679-90ed-574c05701342
Authorization: Bearer <token>
```
