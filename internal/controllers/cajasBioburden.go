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

type CajasBioburdenController interface {
	CrearCajaBioburden(c echo.Context) error
	ObtenerCajasBioburdenID(c echo.Context) error
	ObtenerCajasPorPruebaRecuento(c echo.Context) error
	ActualizarCajaBioburden(c echo.Context) error
	EliminarCajaBioburden(c echo.Context) error
}

type cajasBioburdenController struct {
	Repo repository.CajasBioburdenRepository
}

func NewCajasBioburdenController(repo repository.CajasBioburdenRepository) CajasBioburdenController {
	return &cajasBioburdenController{Repo: repo}
}

// CrearCajaBioburden maneja la creación de una nueva caja de bioburden
func (controller *cajasBioburdenController) CrearCajaBioburden(c echo.Context) error {
	var requestBody struct {
		Tipo                 string `json:"tipo"`
		MetodoSiembra        string `json:"metodoSiembra"`
		MedidaAritmetica     string `json:"medidaAritmetica"`
		FechayhoraIncubacion string `json:"fechayhoraIncubacion"`
		FechayhoraLectura    string `json:"fechayhoraLectura"`
		FactorDisolucion     string `json:"factorDisolucion"`
		IdPruebaRecuento     int    `json:"idPruebaRecuento"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request"})
	}

	cajaBioburden := models.CajasBioburden{
		Tipo:                 requestBody.Tipo,
		MetodoSiembra:        requestBody.MetodoSiembra,
		MedidaAritmetica:     requestBody.MedidaAritmetica,
		FechayhoraIncubacion: requestBody.FechayhoraIncubacion,
		FechayhoraLectura:    requestBody.FechayhoraLectura,
		FactorDisolucion:     requestBody.FactorDisolucion,
		IdPruebaRecuento:     requestBody.IdPruebaRecuento,
	}

	if err := validation.Validate(cajaBioburden.ToMap(), validation.CajasBioburdenRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	if err := controller.Repo.CrearCajaBioburden(&cajaBioburden); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear la caja de bioburden", Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Response{Message: "Caja de bioburden creada exitosamente"})
}

func (controller *cajasBioburdenController) ObtenerCajasBioburdenID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido"})
	}

	cajaBioburden, err := controller.Repo.ObtenerCajasBioburdenID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Caja de bioburden no encontrada"})
	}

	return c.JSON(http.StatusOK, response.Response{Data: cajaBioburden})
}

func (controller *cajasBioburdenController) ObtenerCajasPorPruebaRecuento(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID de prueba de recuento inválido"})
	}

	cajas, err := controller.Repo.ObtenerCajasPorPruebaRecuento(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Cajas de bioburden no encontradas para la prueba de recuento"})
	}

	return c.JSON(http.StatusOK, response.Response{Data: cajas})
}

func (controller *cajasBioburdenController) ActualizarCajaBioburden(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido"})
	}

	cajaBioburden, err := controller.Repo.ObtenerCajasBioburdenID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Caja de bioburden no encontrada"})
	}

	var requestBody struct {
		Tipo                 string `json:"tipo"`
		Resultado            string `json:"resultado"`
		MetodoSiembra        string `json:"metodoSiembra"`
		MedidaAritmetica     string `json:"medidaAritmetica"`
		FechayhoraIncubacion string `json:"fechayhoraIncubacion"`
		FechayhoraLectura    string `json:"fechayhoraLectura"`
		FactorDisolucion     string `json:"factorDisolucion"`
		IdPruebaRecuento     int    `json:"idPruebaRecuento"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request"})
	}

	cajaBioburden.Tipo = requestBody.Tipo
	cajaBioburden.Resultado = requestBody.Resultado
	cajaBioburden.MetodoSiembra = requestBody.MetodoSiembra
	cajaBioburden.MedidaAritmetica = requestBody.MedidaAritmetica
	cajaBioburden.FechayhoraIncubacion = requestBody.FechayhoraIncubacion
	cajaBioburden.FechayhoraLectura = requestBody.FechayhoraLectura
	cajaBioburden.FactorDisolucion = requestBody.FactorDisolucion
	cajaBioburden.IdPruebaRecuento = requestBody.IdPruebaRecuento

	if err := validation.Validate(cajaBioburden.ToMap(), validation.CajasBioburdenRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	if err := controller.Repo.ActualizarCajaBioburden(cajaBioburden); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al actualizar la caja de bioburden"})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Caja de bioburden actualizada exitosamente"})
}

func (controller *cajasBioburdenController) EliminarCajaBioburden(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido"})
	}

	if err := controller.Repo.EliminarCajaBioburden(id); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al eliminar la caja de bioburden"})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Caja de bioburden eliminada exitosamente"})
}
