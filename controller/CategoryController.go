package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"ocg.com/model"
	repo "ocg.com/repository"
)

func GetAllCategory(c *fiber.Ctx) error {
	return c.JSON(repo.GetAllCategories())
}

func GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	category, err := repo.FindCategoryById(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(category)
}

func FindCategoryById(c *fiber.Ctx) int {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(c.Status(400).SendString(err.Error()))
	}
	category, err := repo.FindCategoryById(id)
	if err != nil {
		fmt.Println(c.Status(404).SendString(err.Error()))
	}
	return (int(category.Id))
}

func DeleteCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.DeleteCategoryById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(model.Category)

	err := c.BodyParser(&category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	categoryId := repo.CreateNewCategory(category)
	return c.SendString(fmt.Sprintf("New category is created successfully with id = %d", categoryId))

}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	updatedCategory := new(model.Category)

	err = c.BodyParser(&updatedCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.UpdateCategoryById(int64(id), updatedCategory)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Category with id = %d is successfully updated", id))

}

func UpsertCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	category := new(model.Category)

	err = c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	repo.UpsertCategoryById(int64(id), category)
	return c.SendString(fmt.Sprintf("Category with id = %d is successfully upserted", id))
}
