package API

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

var secretKey = []byte("your_secret_key")

func createToken(username string) (string, int64, error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expiresAt,
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}
	return tokenString, expiresAt, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func CreateUserHandler(c *gin.Context) {
	var user User
	fmt.Println("user:", user)

	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(400, gin.H{"error": "Content-Type must be application/json"})
		return
	}

	// Check if there's actually content
	if c.Request.ContentLength == 0 {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	data, err := os.ReadFile("./users/users.json")
	if err != nil {
		fmt.Println("File not found, creating new one")
		data = []byte("[]")
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, u := range users {
		if u.UserName == user.UserName {
			c.JSON(400, gin.H{"error": "User already exists"})
			return
		}
	}

	users = append(users, user)
	updated, _ := json.MarshalIndent(users, "", " ")

	os.WriteFile("./users/users.json", updated, 0644)

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	data, err := os.ReadFile("./users/users.json")
	if err != nil {
		fmt.Println("File not found")
		data = []byte("[]")
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, u := range users {
		if u.UserName == user.UserName {
			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
				c.JSON(401, gin.H{"error": "Invalid password"})
				return
			}
			token, expiresAt, err := createToken(u.UserName)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, gin.H{"message": "Login successful", "token": token, "expiresAt": expiresAt})
			return
		}
	}

}

func ValidateTokenHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	fmt.Println("tokenString:", tokenString)
	if err := ValidateTokenFunc(tokenString); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Token is valid"})
}

func ValidateTokenFunc(tokenString string) error {
	if tokenString == "" {
		return fmt.Errorf("no token provided")
	}
	tokenString = tokenString[len("Bearer "):]

	if err := verifyToken(tokenString); err != nil {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func TestAuthHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	fmt.Println("tokenString:", tokenString)
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "No token provided"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	if err := verifyToken(tokenString); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Token is valid"})
}
