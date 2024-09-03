package controller

import (
	"net/http"

	"hello-go-deploy/response"

	"github.com/labstack/echo/v4"
)

func List(ctx echo.Context) error {
	return response.ListRes(ctx, http.StatusOK, "text")
}