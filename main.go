package main

import (
	"net/http"
	"os"
	"sigmatech-test/database"
	"sigmatech-test/domain/auth"
	"sigmatech-test/domain/transaction"
	"sigmatech-test/domain/user"
	"sigmatech-test/domain/user_limit"
	"sigmatech-test/handler"
	"sigmatech-test/helper"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	userRepository := user.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	userLimitRepository := user_limit.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userLimitService := user_limit.NewService(userLimitRepository)
	transactionService := transaction.NewService(transactionRepository, userLimitRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	userLimitHandler := handler.NewUserLimitHandler(userLimitService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	route.POST("/email_checkers", userHandler.CheckEmailAvailability)
	route.POST("/register", userHandler.RegisterUser)
	route.POST("/login", userHandler.Login)

	route.GET("/profile", authMiddleware(authService, userService), userHandler.FetchUser)
	route.PUT("/profile", authMiddleware(authService, userService), userHandler.UpdateUser)

	route.GET("/transaction", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	route.POST("/transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	route.GET("/user-limit", authMiddleware(authService, userService), userLimitHandler.GetUserLimits)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	route.Run(":" + port)
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", *user)
	}
}
