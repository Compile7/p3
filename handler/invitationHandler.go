package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/viper"
	"net/http"
	"p3/datamodels"
	cErr "p3/err"
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
	router.GET("/invite", handler.AcceptInvitation)
	router.POST("/invite", utils.AddRoles(handler.SendInvitation))

}

func (h InvitationHandler) SendInvitation(ctx echo.Context) error {
	//org := ctx.Get("OrganisationName").(string)
	//user := ctx.Get("LoggeddInUserEmail").(string)
	var input datamodels.SendInvitationReq
	roles := ctx.Get("Roles").(utils.EmployeeRoles)
	if roles.Role != "Admin" {
		return utils.GetError(cErr.NotAllowed)
	}
	ctx.Bind(&input)
	domain, _ := viper.Get("domain").(string)

	user := "ankit"
	ids := []string{}
	links := []string{}

	for range input.Receivers {
		id, _ := gonanoid.New()
		ids = append(ids, id)
		links = append(links, fmt.Sprintf("%s/invitation/%s", domain, id))
	}
	for i, receiver := range input.Receivers {
		utils.InviteUser(utils.InviteEmployee{Link: links[i], Organisation: roles.OrgName}, utils.EmailOptions{Sender: user, Receiver: receiver})

	}
	return ctx.String(http.StatusOK, "invitation sent successfully")

}
func (h InvitationHandler) AcceptInvitation(ctx echo.Context) error {

	//redirect
	return ctx.Redirect(http.StatusOK, "Accepted")
}
