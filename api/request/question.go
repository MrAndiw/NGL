package request

type ResponseInsert struct {
	Data    interface{}
	Message string `json:"message"`
	Status  string `json:"status"`
}

type QuestionRequest struct {
	ToUser   string `json:"to_user"`
	Question string `json:"question"`
}

type DeleteQuestionRequest struct {
	Id string `json:"id"`
}

type ReadQuestionRequest struct {
	Id string `json:"id"`
}
