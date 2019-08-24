package messages

import (
	"avito-chat/common"
	"avito-chat/database"
	"avito-chat/models"

	"github.com/jackc/pgx"
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

}
