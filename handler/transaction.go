package handler

import (
	"net/http"
	"sigmatech-test/domain/transaction"
	"sigmatech-test/helper"
	"sigmatech-test/model"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)

	transactions, err := h.service.GetTransactions(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Error to get transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of transactions", http.StatusOK, "success", transaction.FormatTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
