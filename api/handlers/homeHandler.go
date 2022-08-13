package handlers

import (
	"NGL/db"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type QuestionParam struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

// ========
// PAGE HOME
// ========
func Home(c echo.Context) error {
	data := M{
		"title":   "Welcome To NGL",
		"message": "Please Share Your link to Instagram",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// ========
// PAGE INBOX
// ========
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

		// encode to base64
		questionParam := QuestionParam{
			ID:       val.ID.String(),
			Question: val.Question,
		}

		b64Question, err := encodeToBase64(questionParam)
		if err != nil {
			return err
		}

		htmlIbox += `
			<div class="col-lg-3">
				<a href="https://ngl-overprogram-2022.herokuapp.com/inbox-detail/` + b64Question + `"><img src="https://creazilla-store.fra1.digitaloceanspaces.com/emojis/43366/love-letter-emoji-clipart-md.png" alt="` + val.Question + `" width="100" height="100"></a>
			</div>	`
	}

	data := M{
		"questions": questions,
		"htmlIbox":  htmlIbox,
	}

	return c.Render(http.StatusOK, "inbox.html", data)
}

// ========
// PAGE IBOX DETAIL
// ========
func InboxDetail(c echo.Context) error {

	id := c.Param("id")

	// decode from base64
	var val QuestionParam
	if err := decodeFromBase64(&val, id); err != nil {
		return err
	}

	htmlIbox := ""

	data := M{
		"question": val.Question,
		"htmlIbox": htmlIbox,
	}

	return c.Render(http.StatusOK, "inbox-detail.html", data)
}

func encodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func decodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}
