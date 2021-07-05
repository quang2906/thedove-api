package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"ocg.com/model"
	repo "ocg.com/repository"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.GetAllReviews())
}

func GetReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review, err := repo.FindReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(review)
}

func FindReviewById(c *fiber.Ctx) int {
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(c.Status(400).SendString(err.Error()))
	}
	review, err := repo.FindReviewById(int64(id))
	if err != nil {
		fmt.Println(c.Status(404).SendString(err.Error()))
	}
	return (int(review.Id))
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	reviewId := repo.CreateNewReview(review)
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func UpdateReview(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	updatedReview := new(model.Review)

	err = c.BodyParser(&updatedReview)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.UpdateReviewById(int64(id), updatedReview)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Review with id = %d is successfully updated", id))

}

func UpsertReview(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	review := new(model.Review)

	err = c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	repo.UpsertReviewById(int64(id), review)
	return c.SendString(fmt.Sprintf("Review with id = %d is successfully upserted", id))
}
