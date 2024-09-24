package controllers

import (
	requestmodels "github.com/ZiplEix/pixel-espion/request_models"
	"github.com/ZiplEix/pixel-espion/services"
	"github.com/ZiplEix/pixel-espion/validation"
	"github.com/gofiber/fiber/v2"
)

// Pixel1 godoc
// @Summary Get Spy Image
// @Description Retrieve an image associated with a spy by their ID and log the visit record.
// @Tags spy
// @Accept  json
// @Produce  png
// @Param id query string true "Spy ID"
// @Success 200 {file} file "Returns the spy image"
// @Failure 400 {object} errorResponse "Bad Request: Spy ID is required"
// @Failure 404 {object} errorResponse "Not Found: Spy not found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /spy/pixel1 [get]
func Pixel1(c *fiber.Ctx) error {
	spyId := c.Query("id")
	if spyId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "Spy ID is required",
		})
	}

	imagePath, err := services.Pixel1(spyId, c.IP())
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	c.Set("Content-Type", "image/png")
	return c.SendFile(imagePath)
}

// NewSpy godoc
// @Summary Create a new spy
// @Description Creates a new spy for the authenticated user
// @Tags spies
// @Accept json
// @Produce json
// @Param NewSpyRequest body requestmodels.NewSpyRequest true "Spy information"
// @Success 201 {object} fiber.Map{spy_id=int} "Created"
// @Failure 400 {object} fiber.Map{error=string} "Bad Request"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /spy/new [post]
func NewSpy(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var req requestmodels.NewSpyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	spyId, err := services.NewSpy(req, userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"spy_id": spyId,
	})
}

// GetAllSpies godoc
// @Summary Retrieve all spies for the authenticated user
// @Description Returns a list of spies associated with the authenticated user
// @Tags spies
// @Produce json
// @Success 200 {object} fiber.Map{spies=[]models.Spy} "List of spies"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /spy/all [get]
func GetAllSpies(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	spies, err := services.GetAllSpies(userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"spies": spies,
	})
}

// GetSpy godoc
// @Summary Retrieve a spy by ID
// @Description Returns the details of a specific spy based on the provided ID
// @Tags spies
// @Produce json
// @Param id path string true "Spy ID"
// @Success 200 {object} models.Spy "Spy details"
// @Failure 404 {object} fiber.Map{error=string} "Spy not found"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /spy/{id} [get]
func GetSpy(c *fiber.Ctx) error {
	spyId := c.Params("id")

	spy, err := services.GetSpy(spyId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(spy)
}

// GetSpyRecords godoc
// @Summary Retrieve records of a specific spy
// @Description Returns all records associated with a specific spy based on the provided spy ID
// @Tags records
// @Produce json
// @Param id path string true "Spy ID"
// @Success 200 {object} fiber.Map{records=[]models.Record} "List of records for the spy"
// @Failure 404 {object} fiber.Map{error=string} "Spy not found"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /record/spy/{id} [get]
func GetSpyRecords(c *fiber.Ctx) error {
	spyId := c.Params("id")

	records, err := services.GetSpyRecords(spyId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"records": records,
	})
}

// GetAllRecords godoc
// @Summary Retrieve all records for the authenticated user
// @Description Returns all records associated with the user's spies
// @Tags records
// @Produce json
// @Success 200 {object} fiber.Map{records=[]models.Record} "List of records for the user"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /record/all [get]
func GetAllRecords(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	records, err := services.GetAllRecords(userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"records": records,
	})
}

// UpdateSpy godoc
// @Summary Update a spy's name and color
// @Description Updates the name and color of a spy specified by ID, only if the user is the owner
// @Tags spies
// @Accept json
// @Produce json
// @Param id path string true "Spy ID"
// @Param spy body requestmodels.NewSpyRequest true "Spy update details"
// @Success 204 "No Content"
// @Failure 400 {object} fiber.Map{error=string} "Bad Request"
// @Failure 403 {object} fiber.Map{error=string} "Unauthorized"
// @Failure 404 {object} fiber.Map{error=string} "Spy Not Found"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /spy/{id} [put]
func UpdateSpy(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	spyId := c.Params("id")

	var req requestmodels.NewSpyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := validation.NewSpy(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	err = services.UpdateSpy(spyId, req, userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteSpy godoc
// @Summary Delete a spy
// @Description Deletes a spy specified by ID, only if the user is the owner
// @Tags spies
// @Param id path string true "Spy ID"
// @Success 204 "No Content"
// @Failure 403 {object} fiber.Map{error=string} "Unauthorized"
// @Failure 404 {object} fiber.Map{error=string} "Spy Not Found"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /spy/{id} [delete]
func DeleteSpy(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	spyId := c.Params("id")

	err := services.DeleteSpy(spyId, userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteRecord godoc
// @Summary Delete a record
// @Description Deletes a record specified by ID, only if the associated spy belongs to the user
// @Tags records
// @Param id path string true "Record ID"
// @Success 204 "No Content"
// @Failure 403 {object} fiber.Map{error=string} "Unauthorized"
// @Failure 404 {object} fiber.Map{error=string} "Record Not Found"
// @Failure 500 {object} fiber.Map{error=string} "Internal Server Error"
// @Router /record/{id} [delete]
func DeleteRecord(c *fiber.Ctx) error {
	recordId := c.Params("id")
	userId := c.Locals("userId").(uint)

	err := services.DeleteRecord(recordId, userId)
	if err != nil {
		return c.Status(err.(services.ServiceError).Code).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
