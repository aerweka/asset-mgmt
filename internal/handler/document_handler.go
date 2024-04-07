package handler

import (
	"log"

	"asset-management.com/database"
	"asset-management.com/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Index(c *fiber.Ctx) error {
	db := database.DB

	var documents []model.Document

	db.Find(&documents)
	if len(documents) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No documents found", "data": nil})
	}

	return c.JSON(fiber.Map{"message": "Documents retrieved successfully", "data": documents})
}

func CreateDocument(c *fiber.Ctx) error {
	db := database.DB
	document := new(model.Document)
	log.Print(document)
	// Store the body in the note and return error if encountered
	err := c.BodyParser(document)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	document.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&document).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": document})
}

func GetDocument(c *fiber.Ctx) error {
	db := database.DB
	var document model.Document

	// Read the param noteId
	id := c.Params("id")

	// Find the note with the given Id
	db.Find(&document, "id = ?", id)

	// If no such note present return an error
	if document.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No document present", "data": nil})
	}

	// Return the note with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Document Found", "data": document})
}

func UpdateDocument(c *fiber.Ctx) error {
	type updateDocument struct {
		Name         string `json:"name"`
		DueDate      string `json:"due_date"`
		ActivaCode   string `json:"activa_code"`
		DocumentType string `json:"document_type"`
	}
	db := database.DB
	var document model.Document

	// Read the param noteId
	id := c.Params("id")

	// Find the note with the given Id
	db.Find(&document, "id = ?", id)

	// If no such note present return an error
	if document.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateDocumentData updateDocument
	err := c.BodyParser(&updateDocumentData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the note
	document.Name = updateDocumentData.Name
	document.ActivaCode = updateDocumentData.ActivaCode
	document.DocumentType = updateDocumentData.DocumentType
	document.DueDate = updateDocumentData.DueDate

	// Save the Changes
	db.Save(&document)

	// Return the updated note
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": document})
}
