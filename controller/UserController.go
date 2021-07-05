package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"ocg.com/model"
	repo "ocg.com/repository"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(repo.GetAllUsers())
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user, err := repo.FindUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(user)
}

func FindUserById(c *fiber.Ctx) int {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(c.Status(400).SendString(err.Error()))
	}
	user, err := repo.FindUserById(int64(id))
	if err != nil {
		fmt.Println(c.Status(404).SendString(err.Error()))
	}
	return (int(user.Id))
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.DeleteUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	userId := repo.CreateNewUser(user)
	return c.SendString(fmt.Sprintf("New user is created successfully with id = %d", userId))

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	updatedUser := new(model.User)

	err = c.BodyParser(&updatedUser)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.UpdateUserById(int64(id), updatedUser)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("User with id = %d is successfully updated", updatedUser.Id))

}

func UpsertUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user := new(model.User)

	err = c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	repo.UpsertUserById(int64(id), user)
	return c.SendString(fmt.Sprintf("User with id = %d is successfully upserted", id))
}
