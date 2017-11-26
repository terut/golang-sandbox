package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	NoContent(int) error
	JSON(int, interface{}) error
}
