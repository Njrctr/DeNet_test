package handlers

import (
	"net/http"
	"strconv"

	"github.com/Njrctr/DeNet_test/models"
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

// type referrerInput struct {
// 	ReferrerCode string `json:"referrer_code"`
// }

// @Summary Input REF code
// @Security ApiKeyAuth
// @Tags Users
// @Description input ref code
// @ID post-ref
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param input body referrerInput true "Referrer Code"
// @Success 200 {object} models.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id}/referrer [post]
// func (h *Handler) usersReferrer(c *gin.Context) {

// 	userId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "Invalid user id param")
// 		return
// 	}

// 	var refCode string
// 	if err := c.BindJSON(&refCode); err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "Invalid ref code")
// 		return

// 	}

// 	users, err := h.services.GetUsersLeaderboard()
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, leaderboardResponce{
// 		Data: users,
// 	})
// }
