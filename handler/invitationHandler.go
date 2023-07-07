package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"p3/utils"

	"gorm.io/gorm"
)

type InvitationHandler struct {
	dbInstance *gorm.DB
}

func NewInvitationHandler(router *echo.Echo, db *gorm.DB) {
	handler := &InvitationHandler{
		dbInstance: db,
	}
	router.GET("/invite", handler.SendInvitation)
}

func (h InvitationHandler) SendInvitation(ctx echo.Context) error {
	org := ctx.Get("OrganisationName").(string)
	user := ctx.Get("LoggeddInUserEmail").(string)
	receiver := ctx.QueryParam("receiver")
	utils.InviteUser(utils.InviteEmployee{Link: "www.google.com", Organisation: org}, utils.EmailOptions{Sender: user, Receiver: receiver})
	ctx.String(http.StatusOK, "invitation sent")

}
