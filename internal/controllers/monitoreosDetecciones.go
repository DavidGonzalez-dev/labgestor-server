package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type MonitoreosDeteccionesController interface {
	CrearMonitoreosDetecciones(c echo.Context) error
	ObtenerMonitoreosDeteccionesPorDeteccion(c echo.Context) error
	ActualizarMonitoreosDetecciones(c echo.Context) error
	EliminarMonitoreosDetecciones(c echo.Context) error
}

type monitoreosDeteccionesController struct {
	Repo repository.MonitoreosDeteccionRepository
}

func NewMonitoreosDeteccionesRepository(repo repository.MonitoreosDeteccionRepository) MonitoreosDeteccionesController {
	return &monitoreosDeteccionesController{Repo: repo}
}

// CrearMonitoreosDetecciones maneja la creación de una nueva detección de monitoreo
func (controller *monitoreosDeteccionesController) CrearMonitoreosDetecciones(c echo.Context) error {
	var requestBody struct {
		VolumenMuestra            string `json:"volumenMuestra"`
		NombreDiluyente           string `json:"nombreDiluyente"`
		FechayhoraInicio          time.Time `json:"fechayhoraInicio"`
		FechayhoraFinal           time.Time `json:"fechayhoraFinal"`
		IdEtapaDeteccion          int    `json:"idEtapaDeteccion"`
		IdDeteccionMicroorganismo int    `json:"idDeteccionMicroorganismo"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al procesar la solicitud", Error: err.Error()})
	}

	monitoreosDeteccion := &models.MonitoreosDeteccionesMicroorganismo{
		VolumenMuestra:            requestBody.VolumenMuestra,
		NombreDiluyente:           requestBody.NombreDiluyente,
		FechayhoraInicio:          requestBody.FechayhoraInicio,
		FechayhoraFinal:           requestBody.FechayhoraFinal,
		IdEtapaDeteccion:          requestBody.IdEtapaDeteccion,
		IdDeteccionMicroorganismo: requestBody.IdDeteccionMicroorganismo,
	}

	if err := controller.Repo.CrearMonitoreosDetecciones(monitoreosDeteccion); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear la detección de monitoreo", Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Response{Message: "Detección de monitoreo creada exitosamente"})
}

func (controller *monitoreosDeteccionesController) ObtenerMonitoreosDeteccionesPorDeteccion(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: err.Error()})
	}
	detecciones, err := controller.Repo.ObtenerMonitoreosDeteccionesPorDeteccion(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener las detecciones de monitoreo", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Detecciones de monitoreo obtenidas exitosamente", Data: detecciones})
}

func (controller *monitoreosDeteccionesController) ActualizarMonitoreosDetecciones(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido"})
	}

	monitoreosDeteccion, err := controller.Repo.ObtenerMonitoreosDeteccionesID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Caja de bioburden no encontrada"})
	}
	var requestBody struct {
		VolumenMuestra            string `json:"volumenMuestra"`
		NombreDiluyente           string `json:"nombreDiluyente"`
		FechayhoraInicio          time.Time `json:"fechayhoraInicio"`
		FechayhoraFinal           time.Time `json:"fechayhoraFinal"`
		IdEtapaDeteccion          int    `json:"idEtapaDeteccion"`
		IdDeteccionMicroorganismo int    `json:"idDeteccionMicroorganismo"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al procesar la solicitud", Error: err.Error()})
	}

	monitoreosDeteccion.VolumenMuestra = requestBody.VolumenMuestra
	monitoreosDeteccion.NombreDiluyente = requestBody.NombreDiluyente
	monitoreosDeteccion.FechayhoraInicio = requestBody.FechayhoraInicio
	monitoreosDeteccion.FechayhoraFinal = requestBody.FechayhoraFinal
	monitoreosDeteccion.IdEtapaDeteccion = requestBody.IdEtapaDeteccion
	monitoreosDeteccion.IdDeteccionMicroorganismo = requestBody.IdDeteccionMicroorganismo
	
	if err := controller.Repo.ActualizarMonitoreosDetecciones(&monitoreosDeteccion); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar la detección de monitoreo", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Detección de monitoreo actualizada exitosamente", Data: monitoreosDeteccion.ToMap()})
}

func (controller *monitoreosDeteccionesController) EliminarMonitoreosDetecciones(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: err.Error()})
	}

	monitoreosDeteccion := &models.MonitoreosDeteccionesMicroorganismo{ID: id}

	if err := controller.Repo.EliminarMonitoreosDetecciones(monitoreosDeteccion); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al eliminar la detección de monitoreo", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Detección de monitoreo eliminada exitosamente"})
}
