# A Chat API Server (Golang + PostgreSQL)

### Deployment
API lives [here](http://167.71.252.215:9000).

### Build Instructions
Assuming you've cloned this repo:
```
docker-compose up
```
Now you have your server available at `http://localhost:9000`

### API
Method   | Path     | Body                                          | Response                               | Response Body |
-------- | -------- | --------------------------------------------- | -------------------------------------- | ----- |
POST     | /users/add | {"username": \<string\>}                                      | 201 Created, 409 Conflict (if user already exists) | {"username": \<string\>, "user_id": <USER_ID>} |
POST     | /chats/add | {"name": \<string\>, "users": ["<USER_ID_1>", "<USER_ID_2>"]} | 201 Created (assuming all chat users exist) | {"chat_id": <CHAT_ID>, "name": \<string\>, "created_at": \<timestamp\>} |
POST     | /chats/get | {"user": "<USER_ID>"}                                        | 200 OK | [ {"chat_id": <CHAT_ID>, "name": \<string\>, "created_at": \<timestamp\>} ] or [] |
POST     | /messages/add | {"chat": "<CHAT_ID>", "author": "<USER_ID>", "text": \<string\>} | 200 OK, 401 Bad Request (if chat or/and author doesn't exist) | {"message_id": <MESSAGE_ID>, "chat": <CHAT_ID>, "author": \<string\>, "text": \<string\>} |
POST     | /messages/get | {"chat": <CHAT_ID>}  | 200 OK | [ {"message_id": <MESSAGE_ID>, "chat": <CHAT_ID>, "author": \<string\>, "text": \<string\>} ] or [] |