package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PruebaRecuentoController interface {
	CrearPruebaRecuento(c echo.Context) error
	ObtenerPruebaRecuentoID(c echo.Context) error
	ActualizarPruebaRecuento(c echo.Context) error
	ObtenerPruebasPorProducto(c echo.Context) error
	EliminarPruebaRecuento(c echo.Context) error
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

	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
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
	return c.JSON(http.StatusCreated, response.Response{Message: "Prueba de recuento creada correctamente"})
}

func (controller pruebaRecuentoController) ObtenerPruebaRecuentoID(c echo.Context) error {
	idPrueba, err := strconv.Atoi(c.Param("id")) // Este es el ID único de la prueba
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: "El id tiene que ser un numero entero"})
	}

	pruebaRecuento, err := controller.repo.ObtenerPruebaRecuentoID(idPrueba)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{
			Message: "No se encontró la prueba de recuento",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Message: "Prueba de recuento obtenida correctamente",
		Data:    pruebaRecuento,
	})
}

func (controller pruebaRecuentoController) ActualizarPruebaRecuento(c echo.Context) error {
	//? ------------------------------------------------
	//? Se lee el cuerpo del request
	//? ------------------------------------------------
	idParam, err := strconv.Atoi(c.Param("id")) // Este es el ID único de la prueba
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: err.Error()})
	}

	// Se verifica que la prueba exista exista
	pruebaRecuento, err := controller.repo.ObtenerPruebaRecuentoID(idParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "No se encontró la prueba de recuento", Error: err.Error()})
	}

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
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al leer el cuerpo de request", Error: err.Error()})
	}

	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
	pruebaRecuento.MetodoUsado = requestBody.MetodoUsado
	pruebaRecuento.Concepto = requestBody.Concepto
	pruebaRecuento.Especificacion = requestBody.Especificacion
	pruebaRecuento.VolumenDiluyente = requestBody.VolumenDiluyente
	pruebaRecuento.TiempoDisolucion = requestBody.TiempoDisolucion
	pruebaRecuento.CantidadMuestra = requestBody.CantidadMuestra
	pruebaRecuento.Tratamiento = requestBody.Tratamiento
	pruebaRecuento.NombreRecuento = requestBody.NombreRecuento
	pruebaRecuento.NumeroRegistroProducto = requestBody.NumeroRegistroProducto

	if err := validation.Validate(pruebaRecuento.ToMap(), validation.PruebaRecuentoRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}
	//? ------------------------------------------------
	//? Se actualiza el producto
	//? ------------------------------------------------
	if err := controller.repo.ActualizarPruebaRecuento(pruebaRecuento); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al actualizar la prueba de recuento", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Prueba de recuento actualizada correctamente"})
}

func (controller pruebaRecuentoController) ObtenerPruebasPorProducto(c echo.Context) error {
	numeroRegistro := c.Param("numeroRegistro") // Este es el ID único de la prueba

	pruebaRecuento, err := controller.repo.ObtenerPruebasPorProducto(numeroRegistro)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "No se encontró la prueba de recuento", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Prueba de recuento obtenida correctamente", Data: pruebaRecuento})
}

func (controller pruebaRecuentoController) EliminarPruebaRecuento(c echo.Context) error {

	// Se verifica si el producto existe
	idPrueba, err := strconv.Atoi(c.Param("id")) // Este es el ID único de la prueba
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: "El id tiene que ser un numero entero"})
	}

	// Se elimina la prueba de recuento
	err = controller.repo.EliminarPruebaRecuento(idPrueba)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "No se encontró la prueba de recuento", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Prueba de recuento eliminada correctamente"})
}
