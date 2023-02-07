package handler

import (
	"net/http"
	"sigmatech-test/domain/user_limit"
	"sigmatech-test/helper"
	"sigmatech-test/model"

	"github.com/gin-gonic/gin"
)

type userLimitHandler struct {
	service user_limit.Service
}

func NewUserLimitHandler(userLimitService user_limit.Service) *userLimitHandler {
	return &userLimitHandler{userLimitService}
}

func (h *userLimitHandler) GetUserLimits(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)

	userLimits, err := h.service.GetUserLimitsByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Error to get transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of transactions", http.StatusOK, "success", user_limit.FormatUserLimits(userLimits))
	c.JSON(http.StatusOK, response)
}
