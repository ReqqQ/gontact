
# Gontact

Simple API in GO with DDD where user can keep his contacts with groups 

## API Reference

#### Get user

```http
  GET user/:userId
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `userId` | `int` |  |
| `apiKey` | `string` | **Required**. Your API key |

#### Get user contacts

```http
  GET user/:userId/contacts
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `userId`      | `int` | **Required**. |

#### Get user contacts group types

```http
  GET user/:userId/contacts/group/types
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `userId`      | `int` | **Required**. |

#### Get user contacts with filter on group

```http
  GET user/:userId/contacts/group/:groupId
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `userId`      | `int` | **Required**. |
| `groupId`      | `int` | **Required**. |


## Features

- create user 
- create contact
- create type 
- middleware for API Token