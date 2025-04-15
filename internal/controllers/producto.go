package controllers

import (
	"labgestor-server/internal/repository"

	"github.com/labstack/echo/v4"
)

type ProductoController interface {
	ObtenerProductoID(c echo.Context) error
	ObtenerProductos(c echo.Context) error
}

type productoController struct {
	Repo repository.ProductoRepository
}

func NewProductoController(repo repository.ProductoRepository) ProductoController {
	return productoController{Repo: repo}
}

// -------------------------------------
// Metodos del controlador
// -------------------------------------
func (controller productoController) ObtenerProductoID(c echo.Context) error {
	// Se lee el cuerpo del request
	Numero_Registro := c.Param("id")

	// Se obtiene el producto haciendo uso de la capa del repositorio
	producto, err := controller.Repo.ObtenerProductoID(Numero_Registro)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, producto)
}

func (controller productoController) ObtenerProductos(c echo.Context) error {
	// Se obtiene el producto haciendo uso de la capa del repositorio
	productos, err := controller.Repo.ObtenerProductos()
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, productos)
}
