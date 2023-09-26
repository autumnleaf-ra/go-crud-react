package repository

import (
	"net/http"

	"github.com/autumnleaf-ra/go-crud-react/database/migrations"
	"github.com/autumnleaf-ra/go-crud-react/database/migrations/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func (r *Repository) GetUsers(c *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Users{})

	// pagination
	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(c.Request()).Response(&[]migrations.Users{})

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})
	return nil
}

func (r *Repository) CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}

	errors := ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "User has been added", "data": user})
	return nil
}

func (r *Repository) UpdateUser(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Request failed",
			})
		return err
	}

	errors := ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id not found",
		})
		return nil
	}

	if db.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get user profiles",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
	})

	return nil
}

func (r *Repository) DeleteUser(c *fiber.Ctx) error {
	userModel := &migrations.Users{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id not found",
		})
		return nil
	}

	err := r.DB.Delete(userModel, id)

	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not delete",
		})
		return err.Error
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User successfully deleted",
	})

	return nil
}

func (r *Repository) GetUserByID(c *fiber.Ctx) error {
	userModel := &migrations.Users{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id not found",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(userModel).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not delete",
		})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User fetched successfully",
		"data":    userModel,
	})

	return nil
}
