# Swagger Petstore - OpenAPI 3.0

Version: 1.0.27

Base URL: /api/v3

OpenAPI document: https://docs.example.com/docs/petstore-v3/openapi.json

This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about
Swagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!
You can now help us improve the API whether it's by making changes to the definition itself or to the code.
That way, with time, we can improve the API in general, and expose some of the new features in OAS3.

Some useful links:
- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)

## pet

Everything about your Pets

### POST /pet

**Add a new pet to the store.**

Add a new pet to the store.

Operation ID: `addPet`

#### Request body (application/json)

Create a new pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

#### Request body (application/x-www-form-urlencoded)

Create a new pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

#### Request body (application/xml)

Create a new pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

#### Responses

**200** — Successful operation

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid input

**422** — Validation exception

**default** — Unexpected error

### PUT /pet

**Update an existing pet.**

Update an existing pet by Id.

Operation ID: `updatePet`

#### Request body (application/json)

Update an existent pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

#### Request body (application/x-www-form-urlencoded)

Update an existent pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

#### Request body (application/xml)

Update an existent pet in the store

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

#### Responses

**200** — Successful operation

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid ID supplied

**404** — Pet not found

**422** — Validation exception

**default** — Unexpected error

### GET /pet/findByStatus

**Finds Pets by status.**

Multiple status values can be provided with comma separated strings.

Operation ID: `findPetsByStatus`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `status` | query | yes | string | Status values that need to be considered for filter |

#### Responses

**200** — successful operation

Array of objects with fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
[
  {
    "category": {
      "id": 1,
      "name": "Dogs"
    },
    "id": 10,
    "name": "doggie",
    "photoUrls": [
      "string"
    ],
    "status": "available"
  }
]
```

Array of objects with fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid status value

**default** — Unexpected error

### GET /pet/findByTags

**Finds Pets by tags.**

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

Operation ID: `findPetsByTags`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `tags` | query | yes | array of string | Tags to filter by |

#### Responses

**200** — successful operation

Array of objects with fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
[
  {
    "category": {
      "id": 1,
      "name": "Dogs"
    },
    "id": 10,
    "name": "doggie",
    "photoUrls": [
      "string"
    ],
    "status": "available"
  }
]
```

Array of objects with fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid tag value

**default** — Unexpected error

### GET /pet/{petId}

**Find pet by ID.**

Returns a single pet.

Operation ID: `getPetById`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `petId` | path | yes | integer (int64) | ID of pet to return |

#### Responses

**200** — successful operation

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid ID supplied

**404** — Pet not found

**default** — Unexpected error

### POST /pet/{petId}

**Updates a pet in the store with form data.**

Updates a pet resource based on the form data.

Operation ID: `updatePetWithForm`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `petId` | path | yes | integer (int64) | ID of pet that needs to be updated |
| `name` | query | no | string | Name of pet that needs to be updated |
| `status` | query | no | string | Status of pet that needs to be updated |

#### Responses

**200** — successful operation

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

Example (generated from schema):

```json
{
  "category": {
    "id": 1,
    "name": "Dogs"
  },
  "id": 10,
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "status": "available",
  "tags": [
    {
      "id": 0,
      "name": "string"
    }
  ]
}
```

Fields:

- `category` (object)
  - `id` (integer (int64))
  - `name` (string)
- `id` (integer (int64))
- `name` (string, required)
- `photoUrls` (array of string, required)
- `status` (string): pet status in the store — one of `available`, `pending`, `sold`
- `tags` (array of object)
  - `id` (integer (int64))
  - `name` (string)

**400** — Invalid input

**default** — Unexpected error

### DELETE /pet/{petId}

**Deletes a pet.**

Delete a pet.

Operation ID: `deletePet`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `api_key` | header | no | string |  |
| `petId` | path | yes | integer (int64) | Pet id to delete |

#### Responses

**200** — Pet deleted

**400** — Invalid pet value

**default** — Unexpected error

### POST /pet/{petId}/uploadImage

**Uploads an image.**

Upload image of the pet.

Operation ID: `uploadFile`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `petId` | path | yes | integer (int64) | ID of pet to update |
| `additionalMetadata` | query | no | string | Additional Metadata |

#### Request body (application/octet-stream)

Type: string (binary)

#### Responses

**200** — successful operation

Fields:

- `code` (integer (int32))
- `message` (string)
- `type` (string)

Example (generated from schema):

```json
{
  "code": 0,
  "message": "string",
  "type": "string"
}
```

**400** — No file uploaded

**404** — Pet not found

**default** — Unexpected error

## store

Access to Petstore orders

### GET /store/inventory

**Returns pet inventories by status.**

Returns a map of status codes to quantities.

Operation ID: `getInventory`

#### Responses

**200** — successful operation

**default** — Unexpected error

### POST /store/order

**Place an order for a pet.**

Place a new order in the store.

Operation ID: `placeOrder`

#### Request body (application/json)

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

Example (generated from schema):

```json
{
  "complete": true,
  "id": 10,
  "petId": 198772,
  "quantity": 7,
  "shipDate": "date-time",
  "status": "approved"
}
```

#### Request body (application/x-www-form-urlencoded)

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

#### Request body (application/xml)

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

#### Responses

**200** — successful operation

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

Example (generated from schema):

```json
{
  "complete": true,
  "id": 10,
  "petId": 198772,
  "quantity": 7,
  "shipDate": "date-time",
  "status": "approved"
}
```

**400** — Invalid input

**422** — Validation exception

**default** — Unexpected error

### GET /store/order/{orderId}

**Find purchase order by ID.**

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

Operation ID: `getOrderById`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `orderId` | path | yes | integer (int64) | ID of order that needs to be fetched |

#### Responses

**200** — successful operation

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

Example (generated from schema):

```json
{
  "complete": true,
  "id": 10,
  "petId": 198772,
  "quantity": 7,
  "shipDate": "date-time",
  "status": "approved"
}
```

Fields:

- `complete` (boolean)
- `id` (integer (int64))
- `petId` (integer (int64))
- `quantity` (integer (int32))
- `shipDate` (string (date-time))
- `status` (string): Order Status — one of `placed`, `approved`, `delivered`

**400** — Invalid ID supplied

**404** — Order not found

**default** — Unexpected error

### DELETE /store/order/{orderId}

**Delete purchase order by identifier.**

For valid response try integer IDs with value < 1000. Anything above 1000 or non-integers will generate API errors.

Operation ID: `deleteOrder`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `orderId` | path | yes | integer (int64) | ID of the order that needs to be deleted |

#### Responses

**200** — order deleted

**400** — Invalid ID supplied

**404** — Order not found

**default** — Unexpected error

## user

Operations about user

### POST /user

**Create user.**

This can only be done by the logged in user.

Operation ID: `createUser`

#### Request body (application/json)

Created user object

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
{
  "email": "john@email.com",
  "firstName": "John",
  "id": 10,
  "lastName": "James",
  "password": "12345",
  "phone": "12345",
  "userStatus": 1,
  "username": "theUser"
}
```

#### Request body (application/x-www-form-urlencoded)

Created user object

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

#### Request body (application/xml)

Created user object

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

#### Responses

**200** — successful operation

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
{
  "email": "john@email.com",
  "firstName": "John",
  "id": 10,
  "lastName": "James",
  "password": "12345",
  "phone": "12345",
  "userStatus": 1,
  "username": "theUser"
}
```

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

**default** — Unexpected error

### POST /user/createWithList

**Creates list of users with given input array.**

Creates list of users with given input array.

Operation ID: `createUsersWithListInput`

#### Request body (application/json)

Array of objects with fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
[
  {
    "email": "john@email.com",
    "firstName": "John",
    "id": 10,
    "lastName": "James",
    "password": "12345",
    "phone": "12345",
    "userStatus": 1,
    "username": "theUser"
  }
]
```

#### Responses

**200** — Successful operation

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
{
  "email": "john@email.com",
  "firstName": "John",
  "id": 10,
  "lastName": "James",
  "password": "12345",
  "phone": "12345",
  "userStatus": 1,
  "username": "theUser"
}
```

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

**default** — Unexpected error

### GET /user/login

**Logs user into the system.**

Log into the system.

Operation ID: `loginUser`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `username` | query | no | string | The user name for login |
| `password` | query | no | string | The password for login in clear text |

#### Responses

**200** — successful operation

Type: string

Example (generated from schema):

```json
string
```

Type: string

**400** — Invalid username/password supplied

**default** — Unexpected error

### GET /user/logout

**Logs out current logged in user session.**

Log user out of the system.

Operation ID: `logoutUser`

#### Responses

**200** — successful operation

**default** — Unexpected error

### GET /user/{username}

**Get user by user name.**

Get user detail based on username.

Operation ID: `getUserByName`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `username` | path | yes | string | The name that needs to be fetched. Use user1 for testing |

#### Responses

**200** — successful operation

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
{
  "email": "john@email.com",
  "firstName": "John",
  "id": 10,
  "lastName": "James",
  "password": "12345",
  "phone": "12345",
  "userStatus": 1,
  "username": "theUser"
}
```

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

**400** — Invalid username supplied

**404** — User not found

**default** — Unexpected error

### PUT /user/{username}

**Update user resource.**

This can only be done by the logged in user.

Operation ID: `updateUser`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `username` | path | yes | string | name that need to be deleted |

#### Request body (application/json)

Update an existent user in the store

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

Example (generated from schema):

```json
{
  "email": "john@email.com",
  "firstName": "John",
  "id": 10,
  "lastName": "James",
  "password": "12345",
  "phone": "12345",
  "userStatus": 1,
  "username": "theUser"
}
```

#### Request body (application/x-www-form-urlencoded)

Update an existent user in the store

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

#### Request body (application/xml)

Update an existent user in the store

Fields:

- `email` (string)
- `firstName` (string)
- `id` (integer (int64))
- `lastName` (string)
- `password` (string)
- `phone` (string)
- `userStatus` (integer (int32)): User Status
- `username` (string)

#### Responses

**200** — successful operation

**400** — bad request

**404** — user not found

**default** — Unexpected error

### DELETE /user/{username}

**Delete user resource.**

This can only be done by the logged in user.

Operation ID: `deleteUser`

#### Parameters

| Name | In | Required | Type | Description |
|---|---|---|---|---|
| `username` | path | yes | string | The name that needs to be deleted |

#### Responses

**200** — User deleted

**400** — Invalid username supplied

**404** — User not found

**default** — Unexpected error
