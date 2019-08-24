# A Chat API Server (Golang + PostgreSQL)

### Deployment
API lives here.

### Build Instructions
Assuming you've cloned this repo:
```
docker build . -t chat-api-server
docker run -p 9000:9000 --name chat-api-server -t chat-api-server
```
Now you have your server available at `http://localhost:9000`

### API
Method   | Path     | Body                                          | Response                               |
-------- | -------- | --------------------------------------------- | -------------------------------------- |
POST     | /users/add | {"username": \<string\>}                                      | 201 Created, 409 Conflict (if user already exists) |
POST     | /chats/add | {"name": \<string\>, "users": ["<USER_ID_1>", "<USER_ID_2>"]} | 201 Created (assuming all chat users exist) |
POST     | /chats/get | {"user": "<USER_ID>"}                                        | 200 OK |
POST     | /messages/add | {"chat": "<CHAT_ID>", "author": "<USER_ID>", "text": \<string\>} | 200 OK, 401 Bad Request (if chat or/and author doesn't exist) |
POST     | /messages/get | none                                     |  |