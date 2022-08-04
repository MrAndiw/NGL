package handlers

import (
	"NGL/db"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func Home(c echo.Context) error {
	data := M{
		"title":   "Welcome To NGL",
		"message": "Please Share Your link to Instagram",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func Inbox(c echo.Context) error {
	// mongo connection
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// == get data question ===
	toUser := "MRANDIW"
	questions, err := QueryGetQuestion(toUser, client, ctx)
	if err != nil {
		return err
	}

	htmlIbox := ""
	for _, val := range questions {
		htmlIbox += `
			<div class="col-lg-3">
				<a href="https://ngl-overprogram-2022.herokuapp.com/inbox-detail/` + val.ID.Hex() + `"><img src="https://creazilla-store.fra1.digitaloceanspaces.com/emojis/43366/love-letter-emoji-clipart-md.png" alt="` + val.Question + `" width="100" height="100"></a>
			</div>	`
	}

	data := M{
		"questions": questions,
		"htmlIbox":  htmlIbox,
	}

	return c.Render(http.StatusOK, "inbox.html", data)
}

func InboxDetail(c echo.Context) error {

	htmlIbox := ""

	data := M{
		"htmlIbox": htmlIbox,
	}

	return c.Render(http.StatusOK, "inbox-detail.html", data)
}
