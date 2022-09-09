package controllers

import (
	"golang-microservices/mvc/services"
	"golang-microservices/mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		//handle the error to the client
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
