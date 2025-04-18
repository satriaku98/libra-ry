package handler

import (
	"strconv"
	"strings"

	"libra-ry/internal/domain"
	"libra-ry/internal/usecase"
	"libra-ry/pkg"

	"github.com/gofiber/fiber/v2"
)

type BukuHandler struct {
	useCase usecase.BukuUseCase
}

func NewBukuHandler(uc usecase.BukuUseCase) *BukuHandler {
	return &BukuHandler{useCase: uc}
}

// GetBuku godoc
// @Summary Get paginated books
// @Tags books
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param judul query string false "Search by book title"
// @Param penulis query string false "Search by author"
// @Param penerbit query string false "Search by publisher"
// @Param tahun_terbit query int false "Search by year published"
// @Param tags query string false "Comma-separated list of tags (e.g., drama,adventure)"
// @Param sortBy query string false "Comma-separated sort fields (e.g., tahun_terbit,-stok)"
// @Security BearerAuth
// @Success 200 {array} domain.BukuSwagger
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /buku [get]
func (h *BukuHandler) GetBuku(c *fiber.Ctx) error {
	// Ambil parameter halaman
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return pkg.ErrorResponse(c, 400, "Invalid page number")
	}

	// Ambil parameter pencarian
	title := c.Query("judul", "")
	author := c.Query("penulis", "")
	publisher := c.Query("penerbit", "")
	yearStr := c.Query("tahun_terbit", "0")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid year")
	}
	tagsParam := c.Query("tags")
	var tags []string
	if tagsParam != "" {
		for _, p := range strings.Split(tagsParam, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				tags = append(tags, p)
			}
		}
	}
	sortByParam := c.Query("sortBy")
	var sortBy []string
	if sortByParam != "" {
		for _, p := range strings.Split(sortByParam, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				sortBy = append(sortBy, p)
			}
		}
	}

	books, totalBooks, err := h.useCase.GetAllBooks(page, title, author, publisher, year, tags, sortBy)
	if err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to fetch books")
	}
	return pkg.SuccessResponse(c, "Books retrieved successfully", books, totalBooks)
}

// GetBukuByID godoc
// @Summary Get a book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Security BearerAuth
// @Success 200 {object} domain.BukuSwagger
// @Router /buku/{id} [get]
func (h *BukuHandler) GetBukuByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid book ID")
	}

	book, err := h.useCase.GetBookByID(uint(id))
	if err != nil {
		return pkg.ErrorResponse(c, 404, "Book not found")
	}

	return pkg.SuccessResponse(c, "Book retrieved successfully", book, 1)
}

// CreateBuku godoc
// @Summary Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.BukuSwagger true "Book Data"
// @Security BearerAuth
// @Success 201 {string} string "Book created"
// @Router /buku [post]
func (h *BukuHandler) CreateBuku(c *fiber.Ctx) error {
	var buku domain.Buku
	if err := c.BodyParser(&buku); err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid input")
	}

	if err := h.useCase.CreateBook(&buku); err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to create book")
	}

	return pkg.SuccessResponse(c, "Book created successfully", buku, 1)
}

// UpdateBuku godoc
// @Summary Update book details
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body domain.BukuSwagger true "Updated Book Data"
// @Security BearerAuth
// @Success 200 {string} string "Book updated"
// @Router /buku/{id} [put]
func (h *BukuHandler) UpdateBuku(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid book ID")
	}

	var buku domain.Buku
	if err := c.BodyParser(&buku); err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid input")
	}

	buku.ID = uint(id)
	if err := h.useCase.UpdateBook(&buku); err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to update book")
	}
	return pkg.SuccessResponse(c, "Book updated successfully", buku, 1)
}

// DeleteBuku godoc
// @Summary Delete a book
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Security BearerAuth
// @Success 200 {string} string "Book deleted"
// @Router /buku/{id} [delete]
func (h *BukuHandler) DeleteBuku(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid book ID")
	}

	if err := h.useCase.DeleteBook(uint(id)); err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to delete book")
	}

	return pkg.SuccessResponse(c, "Book deleted successfully", nil, 1)
}
