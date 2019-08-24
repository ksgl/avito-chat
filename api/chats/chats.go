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
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	var rows *pgx.Rows
	rows, _ = db.Query(`SELECT c.chat_id, c.name, c.created_at
						FROM chats c
						JOIN userchats uc
						ON (c.chat_id = uc.userchat_id)
						JOIN messages m
						ON (c.chat_id = m.chat)
						WHERE uc.user_id = $1
						ORDER BY m.created_at ASC;`, u.UserID)

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
}
