SET SYNCHRONOUS_COMMIT = 'off';

-------------------------------------- cleanup

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS messages CASCADE;
DROP TABLE IF EXISTS userchats CASCADE;


-------------------------------------- users

CREATE TABLE users (
  user_id       SERIAL PRIMARY KEY,
  username      CITEXT NOT NULL UNIQUE,
  created_at    TIMESTAMPTZ DEFAULT now()
);

-------------------------------------- chats

CREATE TABLE chats (
  chat_id       SERIAL PRIMARY KEY,
  name          CITEXT NOT NULL,
  created_at    TIMESTAMPTZ DEFAULT now()
);

-------------------------------------- userchats

CREATE TABLE userchats (
  userchat_id       INT REFERENCES chats(chat_id),
  user_id           INT REFERENCES users(user_id),

  CONSTRAINT userchats_pkey PRIMARY KEY (userchat_id,user_id)
);

-------------------------------------- messages

CREATE TABLE messages (
  message_id    SERIAL PRIMARY KEY,
  chat          INT NOT NULL REFERENCES chats(chat_id),
  author        INT REFERENCES users(user_id),
  text          TEXT NOT NULL,
  created_at    TIMESTAMPTZ DEFAULT now()
);