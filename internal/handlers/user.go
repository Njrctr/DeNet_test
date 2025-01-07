package handlers

import (
	"net/http"
	"strconv"

	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get User
// @Security ApiKeyAuth
// @Tags Users
// @Description get user
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/status [get]
func (h *Handler) userInfo(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user id param")
		return
	}

	userInfo, err := h.services.GetUserInfo(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct {
		User models.User `json:"user"`
	}{
		User: userInfo,
	})
}

type leaderboardResponce struct {
	Data []models.User `json:"data"`
}

// @Summary Get Users Leaderboad
// @Security ApiKeyAuth
// @Tags Users
// @Description get users leaderboad
// @ID get-users-leaderboad
// @Accept  json
// @Produce  json
// @Success 200 {object} leaderboardResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/leaderboard [get]
func (h *Handler) usersLeaderboard(c *gin.Context) {

	users, err := h.services.GetUsersLeaderboard()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, leaderboardResponce{
		Data: users,
	})
}

type referalCode struct {
	ReferrerCode string `json:"referal_code"`
}

// @Summary Refer code
// @Security ApiKeyAuth
// @Tags Users
// @Description input refer code
// @ID refer-code
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param refer_code body referalCode true "refer code"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/referrer [post]
func (h *Handler) userReferrerCode(c *gin.Context) {
	var taskReferalCode referalCode

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param")
		return
	}

	if err := c.BindJSON(&taskReferalCode); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return

	}

	err = h.services.ReferrerCode(userId, taskReferalCode.ReferrerCode)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"succes": "ok",
	})
}
