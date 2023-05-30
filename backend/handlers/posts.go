package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	return c.SendString("CreatePost")
}

func GetPosts(c *fiber.Ctx) error {
	return c.SendString("GetPosts")
}

func DeletePost(c *fiber.Ctx) error {
	return c.SendString("DeletePost")
}

func RatePost(c *fiber.Ctx) error {
	return c.SendString("RatePost")
}

func GetPost(c *fiber.Ctx) error {
	return c.SendString("GetPost")
}

func CommentOnPost(c *fiber.Ctx) error {
	return c.SendString("CommentOnPost")
}

func DeleteComment(c *fiber.Ctx) error {
	return c.SendString("DeleteComment")
}
