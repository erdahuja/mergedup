# Project: mergedup
# 📁 Collection: users 


## End-point: create user
### Method: POST
>```
>http://localhost:3000/v1/users
>```
### Body (**raw**)

```json
{
    "name": "rishabh",
    "email": "rishabh@example.com",
    "roles": [
        "USER"
    ],
    "password": "rishabh",
    "passwordConfirm": "rishabh"
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4ODY2MjEsImlhdCI6MTY4MDg4MzAyMSwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.wqUWPldniYveFk8hdjETi99fAxHnrYdUGnOMt0yj3uw|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get token
### Method: GET
>```
>http://localhost:3000/v1/users/token/1
>```
### 🔑 Authentication basic

|Param|value|Type|
|---|---|---|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: view users
### Method: GET
>```
>http://localhost:3000/v1/users
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4NjIyMjUsImlhdCI6MTY4MDg1ODYyNSwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.PV4JRkWmkQ-z3lyWL2LHV40Bglykokf3xyBGtk2hKs4|


### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4OTEwNDAsImlhdCI6MTY4MDg4NzQ0MCwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.zRptzPqhHKr0gXx8Er0KP7HKXzgqI2CtWLAEb6iEBro|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: suspend
### Method: PATCH
>```
>http://localhost:3000/v1/users/3
>```
### Body (**raw**)

```json
{
    "active": false
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4ODY2MjEsImlhdCI6MTY4MDg4MzAyMSwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.wqUWPldniYveFk8hdjETi99fAxHnrYdUGnOMt0yj3uw|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get user by id
### Method: GET
>```
>http://localhost:3000/v1/users/2
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4ODgyOTYsImlhdCI6MTY4MDg4NDY5Niwicm9sZXMiOlsiVVNFUiJdfQ.Naa9QYvRnAqajZ62BGN2m7AozpI3n2NQdtX0RlTqYLk|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: cart 


## End-point: create cart
### Method: POST
>```
>http://localhost:3000/v1/cart
>```
### Body (**raw**)

```json
{
    "user_id": 1
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4ODkwMzksImlhdCI6MTY4MDg4ODQzOSwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.wlf0AwLJlBGsKOQNMlKu3aWE9VbdDsuTM05VeTqZUKQ|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: create cart item
### Method: POST
>```
>http://localhost:3000/v1/cart-items
>```
### Body (**raw**)

```json
{
    "cart_id": 1,
    "item_id": 1,
    "quantity": 35
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4ODg2OTYsImlhdCI6MTY4MDg4NTA5Niwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.3WGDZS8fnj1ubjT54OSWPcYi51T4-Hci2Dy6wwJD0y0|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: see cart items
### Method: GET
>```
>http://localhost:3000/v1/cart-items/2
>```
### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4NjM4MTAsImlhdCI6MTY4MDg2MDIxMCwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.AGRscERM3QHYJRAsOZDCgg5Ayw8UGLtE_00ZvlJIsIc|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: delete cart item
### Method: DELETE
>```
>http://localhost:3000/v1/cart-items/2
>```
### Body (**raw**)

```json
{
    "cart_id": 1,
    "item_id": 4
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: items 


## End-point: view items
### Method: GET
>```
>http://localhost:3000/v1/items
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create item
### Method: POST
>```
>http://localhost:3000/v1/items
>```
### Body (**raw**)

```json
{
    "name": "test product",
    "quantity": 10,
    "cost": 100
}
```

### 🔑 Authentication bearer

|Param|value|Type|
|---|---|---|
|token|eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJtZXJnZWR1cCIsInN1YiI6IjEiLCJleHAiOjE2ODA4NjM4MTAsImlhdCI6MTY4MDg2MDIxMCwicm9sZXMiOlsiQURNSU4iLCJVU0VSIl19.AGRscERM3QHYJRAsOZDCgg5Ayw8UGLtE_00ZvlJIsIc|string|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
