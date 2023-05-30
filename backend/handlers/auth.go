package handlers

import (
	"github.com/JingusJohn/go-angular-twiddit/backend/api"
	"github.com/JingusJohn/go-angular-twiddit/backend/storage"
	"github.com/JingusJohn/go-angular-twiddit/backend/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	u := new(UserCreateRequest)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "Cannot parse JSON",
			"msg":      err.Error(),
			"received": u,
		})
	}

	if u.Password != u.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Passwords do not match",
		})
	}

	// validate that the user doesn't already exist
	// query the database for the user
	// if the user exists, return an error
	// if the user doesn't exist, create the user
	// return the user (excluding the password/hash)

	oldRows, err := storage.DB.Queryx("SELECT * FROM users WHERE email = $1 OR username = $2", u.Email, u.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not query database",
			"msg":   err.Error(),
		})
	}
	// scan the rows into a slice
	var users []types.User
	for oldRows.Next() {
		var user types.User
		if err := oldRows.StructScan(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not scan database rows",
				"msg":   err.Error(),
			})
		}
		users = append(users, user)
	}
	defer oldRows.Close()
	// if the slice is not empty, return an error
	if len(users) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	// create the user
	user := types.NewUser(u.Email, u.Username, string(hash))

	// insert the user into the database
	_, err = storage.DB.NamedExec("INSERT INTO users (id, email, username, hash, date_created, date_updated) VALUES (:id, :email, :username, :hash, :date_created, :date_updated)", user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not insert user into database",
			"msg":   err.Error(),
		})
	}

	// create a profile for the user
	profile := types.NewProfile(user.ID, user.Username)

	// insert the profile into the database
	_, err = storage.DB.NamedExec("INSERT INTO profiles (id, user_id, profile_name, date_created, date_updated) VALUES (:id, :user_id, :profile_name, :date_created, :date_updated)", profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not insert profile into database",
			"msg":   err.Error(),
		})
	}

	// return the user and the profile
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    user,
		"profile": profile,
	})
}

func Login(c *fiber.Ctx) error {
	// parse the request body
	u := new(UserLoginRequest)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "Cannot parse JSON",
			"msg":      err.Error(),
			"received": u,
		})
	}

	// query the database for the user
	// if the user exists, check the password
	// if the password is correct, return the user (excluding the password/hash)
	// if the password is incorrect, return an error
	// if the user doesn't exist, return an error

	oldRows, err := storage.DB.Queryx("SELECT * FROM users WHERE email = $1", u.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not query database",
			"msg":   err.Error(),
		})
	}
	// scan the rows into a slice
	var users []types.User
	for oldRows.Next() {
		var user types.User
		if err := oldRows.StructScan(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not scan database rows",
				"msg":   err.Error(),
			})
		}
		users = append(users, user)
	}
	defer oldRows.Close()
	// if the slice is empty, return an error
	if len(users) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User does not exist",
		})
	}

	// check the password
	if err := bcrypt.CompareHashAndPassword([]byte(users[0].Hash), []byte(u.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	// create a session
	session, err := api.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}

	session.Set("authenticated", true)
	session.Set("user_id", users[0].ID)

	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}

func Logout(c *fiber.Ctx) error {
	session, err := api.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}

	err = session.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

/*
Validates that the user is currently has a session
*/
func HealthCheck(c *fiber.Ctx) error {
	session, err := api.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}

	auth := session.Get("authenticated")
	if auth == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Authorized",
		})
	}
}

func GetUser(c *fiber.Ctx) error {
	session, err := api.SessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
			"msg":   err.Error(),
		})
	}

	if session.Get("authenticated") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	userID := session.Get("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// query the database for the user
	var user types.User
	rows, err := storage.DB.Queryx("SELECT * FROM users WHERE id = $1", userID.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not query database",
			"msg":   err.Error(),
		})
	}
	for rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not scan database rows",
				"msg":   err.Error(),
			})
		}
	}
	defer rows.Close()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("DeleteUser")
}
