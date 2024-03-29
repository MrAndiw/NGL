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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Question struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ToUser    string             `bson:"to_user" json:"to_user"`
	Question  string             `bson:"question" json:"question"`
	IsRead    bool               `bson:"is_read" json:"is_read"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// CreateQuestion is public : first function name capital
func CreateQuestion(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "Data Gagal Simpan.",
			Status:  "FAILED",
		})
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Message: "Data Berhasil Simpan.",
		Status:  "SUCCESS",
	})
}

func QueryGetQuestion(toUser string, client *mongo.Client, ctx context.Context) ([]Question, error) {
	var questions = make([]Question, 0)

	// == get data question ===
	coll := db.MgoCollection("questions", client)
	result, err := coll.Find(
		ctx,
		bson.M{
			"to_user": toUser,
		},
		options.Find().SetSort(bson.M{"created_at": -1}),
	)
	if err != nil {
		return questions, err
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var question Question
		err = result.Decode(&question)
		if err != nil {
			return questions, err
		}

		questions = append(questions, question)
	}

	return questions, nil
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
	toUser := req.ToUser
	questions, err := QueryGetQuestion(toUser, client, ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Data:    questions,
		Message: "Data Berhasil Diambil.",
		Status:  "SUCCESS",
	})
}

// DeleteQuestion is public : first function name capital
func DeleteQuestion(c echo.Context) error {
	// bind request
	req := request.DeleteQuestionRequest{}
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

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "[Delete Question] ID not found",
			Status:  "FAILED",
		})
	}

	_, err = table.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "[Delete Question] Gagal delete question",
			Status:  "FAILED",
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "Data Gagal Delete.",
			Status:  "FAILED",
		})
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Message: "Data Berhasil Delete.",
		Status:  "SUCCESS",
	})
}

// ReadQuestion is public : first function name capital
func ReadQuestion(c echo.Context) error {
	// bind request
	req := request.ReadQuestionRequest{}
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

	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "[Set Read Question] ID not found",
			Status:  "FAILED",
		})
	}

	_, err = table.UpdateOne(ctx, bson.D{{Key: "_id", Value: objID}}, bson.D{
		{Key: "$set", Value: bson.M{
			"is_read": true,
		}},
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "[Set Read Question] Gagal set read question",
			Status:  "FAILED",
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, request.ResponseInsert{
			Message: "Data Gagal Set Read.",
			Status:  "FAILED",
		})
	}

	return c.JSON(http.StatusOK, request.ResponseInsert{
		Message: "Data Berhasil Set Read.",
		Status:  "SUCCESS",
	})
}
