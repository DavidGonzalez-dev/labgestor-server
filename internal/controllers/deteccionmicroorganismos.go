package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DeteccionMicroorganismosController interface {
	CrearDeteccionMicroorganismos(c echo.Context) error
	ObtenerDeteccionMicroorganismosID(c echo.Context) error
	ActualizarDeteccionMicroorganismos(c echo.Context) error
	ObtenerDeteccionMicroorganismosPorProducto(c echo.Context) error
	EliminarDeteccionMicroorganismos(c echo.Context) error
}

type deteccionMicroorganismosController struct {
	repo repository.DeteccionMicroorganismosRepository
}

func NewDeteccionMicroorganismosController(repo repository.DeteccionMicroorganismosRepository) DeteccionMicroorganismosController {
	return &deteccionMicroorganismosController{repo: repo}
}

func (controller *deteccionMicroorganismosController) CrearDeteccionMicroorganismos(c echo.Context) error {
	var requestBody struct {
		NombreMicroorganismo   string `json:"nombreMicroorganismo"`
		Especificacion         string `json:"especificacion"`
		Concepto               bool   `json:"concepto"`
		Tratamiento            string `json:"tratamiento"`
		MetodoUsado            string `json:"metodoUsado"`
		CantidadMuestra        string `json:"cantidadMuestra"`
		VolumenDiluyente       string `json:"volumenDiluyente"`
		Resultado              string `json:"resultado"`
		NumeroRegistroProducto string `json:"numeroRegistroProducto"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	deteccionesMicroorganismos := models.DeteccionesMicroorganismos{
		NombreMicroorganismo:   requestBody.NombreMicroorganismo,
		Especificacion:         requestBody.Especificacion,
		Concepto:               requestBody.Concepto,
		Tratamiento:            requestBody.Tratamiento,
		MetodoUsado:            requestBody.MetodoUsado,
		CantidadMuestra:        requestBody.CantidadMuestra,
		VolumenDiluyente:       requestBody.VolumenDiluyente,
		Resultado:              requestBody.Resultado,
		NumeroRegistroProducto: requestBody.NumeroRegistroProducto,
	}
	if err := controller.repo.CrearDeteccionMicroorganismos(&deteccionesMicroorganismos); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear la deteccion de microorganismos", Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Response{Message: "Deteccion de microorganismos creada exitosamente"})
}

func (controller *deteccionMicroorganismosController) ObtenerDeteccionMicroorganismosID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

	deteccion, err := controller.repo.ObtenerDeteccionMicroorganismosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Deteccion de microorganismos no encontrada", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos encontrada", Data: deteccion.ToMap()})
}

func (controller *deteccionMicroorganismosController) ActualizarDeteccionMicroorganismos(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

	deteccion, err := controller.repo.ObtenerDeteccionMicroorganismosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Deteccion de microorganismos no encontrada", Error: err.Error()})
	}

	var requestBody struct {
		NombreMicroorganismo   string `json:"nombreMicroorganismo"`
		Especificacion         string `json:"especificacion"`
		Concepto               bool   `json:"concepto"`
		Tratamiento            string `json:"tratamiento"`
		MetodoUsado            string `json:"metodoUsado"`
		CantidadMuestra        string `json:"cantidadMuestra"`
		VolumenDiluyente       string `json:"volumenDiluyente"`
		Resultado              string `json:"resultado"`
		NumeroRegistroProducto string `json:"numeroRegistroProducto"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	deteccion.NombreMicroorganismo = requestBody.NombreMicroorganismo
	deteccion.Especificacion = requestBody.Especificacion
	deteccion.Concepto = requestBody.Concepto
	deteccion.Tratamiento = requestBody.Tratamiento
	deteccion.MetodoUsado = requestBody.MetodoUsado
	deteccion.CantidadMuestra = requestBody.CantidadMuestra
	deteccion.VolumenDiluyente = requestBody.VolumenDiluyente
	deteccion.Resultado = requestBody.Resultado
	deteccion.NumeroRegistroProducto = requestBody.NumeroRegistroProducto

	if err := controller.repo.ActualizarDeteccionMicroorganismos(deteccion); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar la deteccion de microorganismos", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos actualizada exitosamente"})
}

func (controller *deteccionMicroorganismosController) ObtenerDeteccionMicroorganismosPorProducto(c echo.Context) error {
	// Se obtiene el id del producto
	id := c.Param("id")

	// Se obtiene el producto y se verifica que no hallan errores
	deteccion, err := controller.repo.ObtenerDeteccionMicroorganismosPorProducto(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener los registros", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Registros obtenidos correctamente", Data: deteccion})
}

func (controller *deteccionMicroorganismosController) EliminarDeteccionMicroorganismos(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

	if err := controller.repo.EliminarDeteccionMicroorganismos(id); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Deteccion de microorganismos no encontrada", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos eliminada exitosamente"})
}
