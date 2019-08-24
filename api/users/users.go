package users

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
	u := &models.User{}
	u.UnmarshalJSON(ctx.PostBody())

	err := db.QueryRow(`INSERT INTO users(username) VALUES($1) RETURNING user_id;`, u.Username).Scan(&u.UserID)

	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			common.WriteResponse(ctx, fasthttp.StatusConflict, common.AlreadyExists)

			return
		}
	}

	payload, _ := u.MarshalJSON()
	common.WriteResponse(ctx, fasthttp.StatusCreated, payload)

	return
}
