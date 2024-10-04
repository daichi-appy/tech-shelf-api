package controller

import (
    "net/http"
    "tech-shelf/usecase"

		"github.com/labstack/echo/v4"
)

type IBookController interface {
	SearchBooks(c echo.Context) error
}

type BookController struct{
	bu usecase.IBookUsecase
}

func NewBookController(bu usecase.IBookUsecase) IBookController {
    return &BookController{bu}
}

func (bc *BookController) SearchBooks(c echo.Context) error {
    query := c.QueryParam("q")
    if query == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing query parameter"})
    }

    booksResponse, err := bc.bu.FetchBooksFromRakutenAPI(query)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch books"})
    }

    return c.JSON(http.StatusOK, booksResponse)
}