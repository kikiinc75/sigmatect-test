package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sigmatech-test/database"
	"sigmatech-test/domain/auth"
	"sigmatech-test/domain/user"
	"sigmatech-test/handler"
	"sigmatech-test/payload"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	newUser := payload.RegisterUserInput{
		FullName: "budi doremi",
		Email:    "budi2@example.com",
		Password: "test",
	}
	writer := makeRequest("POST", "/register", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
	user := payload.LoginInput{
		Email:    "budi@example.com",
		Password: "password",
	}
	writer := makeRequest("POST", "/login", user, false)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := payload.LoginInput{
		Email:    "budi@example.com",
		Password: "password",
	}

	writer := makeRequest("POST", "/login", user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["token"]
}

func router() *gin.Engine {
	router := gin.Default()
	db := database.ConnectDB()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.RegisterUser)

	// protectedRoutes := router.Group("/api")
	// protectedRoutes.Use(middleware.JWTAuthMiddleware())
	// protectedRoutes.POST("/entry", AddEntry)
	// protectedRoutes.GET("/entry", GetAllEntries)

	return router
}
