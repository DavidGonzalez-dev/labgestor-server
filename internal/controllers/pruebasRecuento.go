package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Interfaz que define los metodos del controlador
type PruebaRecuentoController interface {
	CrearPruebaRecuento(c echo.Context) error
}

type pruebaRecuentoController struct {
	repo repository.PruebaRecuentoRepository
}

// Funcion para instanciar un controlador
func NewPruebaRecuentoController(repo repository.PruebaRecuentoRepository) PruebaRecuentoController {
	return &pruebaRecuentoController{repo: repo}
}

// ? ------------------------------------------------
// ? CONTROLADORES CRUD
// ? ------------------------------------------------
func (controller pruebaRecuentoController) CrearPruebaRecuento(c echo.Context) error {

	//? ------------------------------------------------
	//? Se lee el cuerpo del request
	//? ------------------------------------------------
	var requestBody struct {
		MetodoUsado            string `json:"metodoUsado"`
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
		Especificacion:         requestBody.Especificacion,
		VolumenDiluyente:       requestBody.VolumenDiluyente,
		TiempoDisolucion:       requestBody.TiempoDisolucion,
		CantidadMuestra:        requestBody.CantidadMuestra,
		Tratamiento:            requestBody.Tratamiento,
		NombreRecuento:         requestBody.NombreRecuento,
		NumeroRegistroProducto: requestBody.NumeroRegistroProducto,
	}

	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
	if err := validation.Validate(pruebasRecuento.ToMap(), validation.PruebaRecuentoRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	//? ------------------------------------------------
	//? Se crea el producto
	//? ------------------------------------------------
	if err := controller.repo.CrearPruebaRecuento(&pruebasRecuento); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear la prueba de recuento", Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Response{Message: "Prueba de recuento creada correctamente", Data: pruebasRecuento})
}
