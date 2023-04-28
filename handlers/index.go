package handlers

import (
	"net/http"

	actions "github.com/kevincobain2000/go-vercel-template/handlers/actions"
	pkg "github.com/kevincobain2000/go-vercel-template/pkg"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
	validator  pkg.RequestValidator
	userAction *actions.UserAction
}

func NewIndexHandler() IndexHandler {
	return IndexHandler{
		validator:  pkg.NewRequestValidator(),
		userAction: actions.NewUserAction(),
	}
}

// IndexRequest for the /
type IndexRequest struct {
	Query string `json:"q" query:"q" form:"q" validate:"required" message:"No value for the query param. q is required"`
	ID    int    `json:"id" query:"id" form:"id" validate:"required" message:"No value for the query param. id is required"`
}

func (h *IndexHandler) IndexHandler(c echo.Context) error {
	request, err := pkg.ValidateRequest(c, &IndexRequest{})
	if err != nil {
		return err
	}

	response := h.userAction.Get(request.Query)

	return c.JSON(http.StatusOK, response)
}
