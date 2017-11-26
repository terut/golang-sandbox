package controllers

type ErrorResponse struct {
	Status uint
	Body   ErrorBody
}

type ErrorBody struct {
	Message string `json: "message"`
	Type    string `json: "type"`
}
