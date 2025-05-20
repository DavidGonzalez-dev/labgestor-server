package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PruebaRecuentoController interface {
	CrearPruebaRecuento(c echo.Context) error
}

type pruebaRecuentoController struct {
	repo repository.PruebaRecuentoRepository
}

func NewPruebaRecuentoController(repo repository.PruebaRecuentoRepository) PruebaRecuentoController {
	return &pruebaRecuentoController{repo: repo}
}

// CrearPruebaRecuento crea una nueva prueba de recuento en la base de datos
func (controller pruebaRecuentoController) CrearPruebaRecuento(c echo.Context) error {
	// Se crea el producto y se verifica que no hallan errores
	var requestBody struct {
		MetodoUsado            string `json:"metodoUsado"`
		Concepto               bool   `json:"concepto"`
		Especificacion         string `json:"especificacion"`
		VolumenDiluyente       string `json:"volumenDiluyente"`
		TiempoDisolucion       string `json:"tiempoDisolucion"`
		CantidadMuestra        string `json:"cantidadMuestra"`
		Tratamiento            string `json:"tratamiento"`
		NombreRecuento         string `json:"nombreRecuento"`
		NumeroRegistroProducto string `json:"numeroRegistroProducto"`
	}
	// Se verifica que el cuerpo del request no este vacio
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	pruebasRecuento := models.PruebaRecuento{
		MetodoUsado:            requestBody.MetodoUsado,
		Concepto:               requestBody.Concepto,
		Especificacion:         requestBody.Especificacion,
		VolumenDiluyente:       requestBody.VolumenDiluyente,
		TiempoDisolucion:       requestBody.TiempoDisolucion,
		CantidadMuestra:        requestBody.CantidadMuestra,
		Tratamiento:            requestBody.Tratamiento,
		NombreRecuento:         requestBody.NombreRecuento,
		NumeroRegistroProducto: requestBody.NumeroRegistroProducto,
	}
	// Se crea el producto y se verifica que no hallan errores

	if err := controller.repo.CrearPruebaRecuento(&pruebasRecuento); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear la prueba de recuento", Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Response{Message: "Prueba de recuento creada correctamente", Data: pruebasRecuento})
}
