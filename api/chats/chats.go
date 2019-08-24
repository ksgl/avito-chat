package chats

import (
	"avito-chat/common"
	"avito-chat/database"
	"avito-chat/models"
	"fmt"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/valyala/fasthttp"
)

var db *pgx.ConnPool

func init() {
	db = database.Connect()
}

func Add(ctx *fasthttp.RequestCtx) {
	c := &models.Chat{}
	c.UnmarshalJSON(ctx.PostBody())

	time := &pgtype.Timestamptz{}

	db.QueryRow(`INSERT INTO chats(name) VALUES($1) RETURNING chat_id, created_at;`, c.Name).Scan(&c.ChatID, time)
	c.CreatedAt = time.Time

	var insertUserchats strings.Builder
	fmt.Fprintf(&insertUserchats, `INSERT INTO userchats(userchat_id,user_id) VALUES `)
	for idx, u := range c.Users {
		if idx == len(c.Users)-1 {
			fmt.Fprintf(&insertUserchats, `('%d', '%d')`, c.ChatID, u)
		} else {
			fmt.Fprintf(&insertUserchats, `('%d', '%d'),`, c.ChatID, u)
		}
	}
	fmt.Fprintf(&insertUserchats, ` ON CONFLICT DO NOTHING;`)
	db.Exec(insertUserchats.String())

	payload, _ := c.MarshalJSON()
	common.WriteResponse(ctx, fasthttp.StatusCreated, payload)

	return
}

func Get(ctx *fasthttp.RequestCtx) {
	u := &models.UserChat{}
	u.UnmarshalJSON(ctx.PostBody())

	var rows *pgx.Rows
	rows, _ = db.Query(`SELECT chat_id, name, created_at
						FROM chats
						JOIN userchats
						ON (chats.chat_id = userchats.userchat_id)
						LEFT JOIN (
							SELECT chat, max(created_at) AS newest
							FROM messages
							GROUP BY chat
						) m
						ON m.chat = chats.chat_id
						WHERE user_id = $1
						ORDER BY newest DESC;`, u.UserID)

	time := &pgtype.Timestamptz{}
	chats := make(models.ChatsArr, 0, 16)
	for rows.Next() {
		chat := models.Chat{}
		rows.Scan(&chat.ChatID, &chat.Name, time)
		chat.CreatedAt = time.Time
		chats = append(chats, &chat)
	}
	rows.Close()

	payload, _ := chats.MarshalJSON()
	common.WriteResponse(ctx, fasthttp.StatusOK, payload)

	return
}
