package handlers

import (
	"NGL/api/request"
	"NGL/db"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Question struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ToUser    string             `bson:"to_user" json:"to_user"`
	Question  string             `bson:"question" json:"question"`
	IsRead    bool               `bson:"is_read" json:"is_read"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// SetQuestion is public : first function name capital
func SetQuestion(c echo.Context) error {
	// bind request
	req := request.QuestionRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	// mongo connection
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		table = db.MgoCollection("questions", client)
	)

	_, err := table.InsertOne(ctx, Question{
		ToUser:    req.ToUser,
		Question:  req.Question,
		IsRead:    false,
		CreatedAt: time.Now().In(time.UTC),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Message: "Data Berhasil Disimpan.",
		Status:  "SUCCESS",
	})
}

// GetQuestion is public : first function name capital
func GetQuestion(c echo.Context) error {
	// bind request
	req := request.QuestionRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	// mongo connection
	client := db.MgoConn()
	defer client.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// == get data question ===
	coll := db.MgoCollection("questions", client)
	result, err := coll.Find(
		ctx,
		bson.M{
			"to_user": req.ToUser,
		},
		options.Find().SetSort(bson.M{"created_at": -1}),
	)
	if err != nil {
		return err
	}

	defer result.Close(ctx)

	var questions = make([]Question, 0)
	for result.Next(ctx) {
		var question Question
		err = result.Decode(&question)
		if err != nil {
			return err
		}

		questions = append(questions, question)
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Data:    questions,
		Message: "Data Berhasil Disimpan.",
		Status:  "SUCCESS",
	})
}
