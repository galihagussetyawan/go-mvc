package handler

import (
	"go-mvc/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func checkValidEmailPassword(users *[]model.User, email string, password string) bool {
	for _, data := range *users {
		if data.Email == email && data.Password == password {
			return true
		}
	}
	return false
}

func createAccessToken(email string, role string) string {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, _ := token.SignedString([]byte("secret-token-key-hardcoded")) //harcoded secret-key
	return jwt
}

func Login(c *fiber.Ctx) error {
	// init harcoded data users
	users := make([]model.User, 2)
	users = append(users, model.User{Email: "agent@gmail.com", Password: "agent", Username: "agentusername", Role: "agent"}, model.User{Email: "admin@gmail.com", Password: "admin", Username: "adminusername", Role: "admin"})

	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Render("register", fiber.Map{
			"STATUS": "email dan password harus diisi",
		})
	}

	isValid := checkValidEmailPassword(&users, email, password)

	var user model.User
	for _, v := range users {
		if v.Email == email {
			user = v
		}
	}

	if !isValid {
		return c.Render("register", fiber.Map{
			"STATUS": "email atau password salah",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    createAccessToken(user.Email, user.Role),
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 168),
	})

	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Redirect("/")
}
