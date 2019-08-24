package messages

import (
	"avito-chat/common"
	"avito-chat/database"
	"avito-chat/models"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/valyala/fasthttp"
)

var db *pgx.ConnPool

func init() {
	db = database.Connect()
}

func Add(ctx *fasthttp.RequestCtx) {
	m := &models.Message{}
	m.UnmarshalJSON(ctx.PostBody())

	err := db.QueryRow(`INSERT INTO messages(chat, author, text)
						VALUES($1, $2, $3) RETURNING message_id;`, m.Chat, m.Author, m.Text).Scan(&m.MessageID)

	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23503" {
			common.WriteResponse(ctx, fasthttp.StatusBadRequest, common.DoesNotExist)

			return
		}
	}

	payload, _ := m.MarshalJSON()
	common.WriteResponse(ctx, fasthttp.StatusOK, payload)

	return
}

func Get(ctx *fasthttp.RequestCtx) {
	c := &models.ChatMessages{}
	c.UnmarshalJSON(ctx.PostBody())

	var rows *pgx.Rows
	rows, _ = db.Query(`SELECT message_id, chat, author, text, messages.created_at
						FROM messages
						JOIN chats
						ON (messages.chat = chats.chat_id)
						WHERE chat_id = $1
						ORDER BY messages.created_at ASC;`, c.ChatID)

	time := &pgtype.Timestamptz{}
	messages := make(models.MessagesArr, 0, 16)
	for rows.Next() {
		message := models.Message{}
		rows.Scan(&message.MessageID, &message.Chat, &message.Author, &message.Text, time)
		message.CreatedAt = time.Time
		messages = append(messages, &message)
	}
	rows.Close()

	payload, _ := messages.MarshalJSON()
	common.WriteResponse(ctx, fasthttp.StatusOK, payload)

	return
}
