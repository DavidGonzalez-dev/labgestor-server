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

type DeteccionMicroorganismosController interface {
	CrearDeteccionMicroorganismos(c echo.Context) error
	ObtenerDeteccionMicroorganismosID(c echo.Context) error
	ActualizarDeteccionMicroorganismos(c echo.Context) error
	ObtenerDeteccionMicroorganismosPorProducto(c echo.Context) error
	EliminarDeteccionMicroorganismos(c echo.Context) error
}

type deteccionMicroorganismosController struct {
	repo               repository.DeteccionMicroorganismosRepository
	ProductoRepository repository.ProductoRepository
}

func NewDeteccionMicroorganismosController(repo repository.DeteccionMicroorganismosRepository, productRepo repository.ProductoRepository) DeteccionMicroorganismosController {
	return &deteccionMicroorganismosController{repo: repo, ProductoRepository: productRepo}
}

// Este hadnler nos permite crear un registro d deteccion de microorganismo en la base de datos
func (controller *deteccionMicroorganismosController) CrearDeteccionMicroorganismos(c echo.Context) error {

	// Se lee el cuerpo del request
	var requestBody struct {
		NombreMicroorganismo   string `json:"nombreMicroorganismo"`
		Especificacion         string `json:"especificacion"`
		Tratamiento            string `json:"tratamiento"`
		MetodoUsado            string `json:"metodoUsado"`
		CantidadMuestra        string `json:"cantidadMuestra"`
		VolumenDiluyente       string `json:"volumenDiluyente"`
		NumeroRegistroProducto string `json:"numeroRegistroProducto"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Se hace la validacion de campos
	deteccionesMicroorganismos := models.DeteccionesMicroorganismos{
		NombreMicroorganismo:   requestBody.NombreMicroorganismo,
		Especificacion:         requestBody.Especificacion,
		Tratamiento:            requestBody.Tratamiento,
		MetodoUsado:            requestBody.MetodoUsado,
		CantidadMuestra:        requestBody.CantidadMuestra,
		VolumenDiluyente:       requestBody.VolumenDiluyente,
		NumeroRegistroProducto: requestBody.NumeroRegistroProducto,
		Estado:                 "pendiente",
	}

	if err := validation.Validate(deteccionesMicroorganismos.ToMap(), validation.DetecccionMicroorganismosRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	// Se verifica que el producto exista
	producto, _ := controller.ProductoRepository.ObtenerInfoProducto(requestBody.NumeroRegistroProducto)
	if producto == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al crear la prueba re recuento", Error: "El producto al que le estas asignando la prueba de recuento no existe"})
	}

	// Se crea la prueba de recuento si el producto aun tiene un estado diferente a "terminado"(3)
	if producto.IDEstado != 3 {

		// Se crea la prueba de recuento
		if err := controller.repo.CrearDeteccionMicroorganismos(&deteccionesMicroorganismos); err != nil {
			return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear la deteccion de microorganismo", Error: err.Error()})
		}

		if err := controller.ProductoRepository.ActualizarEstadoProducto(2, requestBody.NumeroRegistroProducto); err != nil {
			return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear la deteccion de microorganismo", Error: err.Error()})
		}

		return c.JSON(http.StatusCreated, response.Response{Message: "Deteccion de microorganismo creada correctamente"})
	}

	return c.JSON(http.StatusCreated, response.Response{Message: "Deteccion de microorganismo creada exitosamente"})
}

// Este handler nos permite obtener un registro de deteccion de microorganismo basado en un id
func (controller *deteccionMicroorganismosController) ObtenerDeteccionMicroorganismosID(c echo.Context) error {

	// Se obtiene el id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

	// Se obtiene el registro
	deteccion, err := controller.repo.ObtenerDeteccionMicroorganismosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Hubo un error al obtener el registro de deteccion de microorganismo", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos encontrada", Data: deteccion})
}

// Este handler nos permite actualizar un registro de deteccion de microorganismo
func (controller *deteccionMicroorganismosController) ActualizarDeteccionMicroorganismos(c echo.Context) error {

	// Obtenemos el id del registro que se desea actualizar
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

	// Obtenemos el registro desde la base de datos
	deteccion, err := controller.repo.ObtenerDeteccionMicroorganismosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Deteccion de microorganismos no encontrada", Error: err.Error()})
	}

	// Leemos el cuerpo del request
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

	// Realizamos la actualizacion de datos y validamos los datos
	deteccion.NombreMicroorganismo = requestBody.NombreMicroorganismo
	deteccion.Especificacion = requestBody.Especificacion
	deteccion.Concepto = requestBody.Concepto
	deteccion.Tratamiento = requestBody.Tratamiento
	deteccion.MetodoUsado = requestBody.MetodoUsado
	deteccion.CantidadMuestra = requestBody.CantidadMuestra
	deteccion.VolumenDiluyente = requestBody.VolumenDiluyente
	deteccion.Resultado = requestBody.Resultado
	deteccion.NumeroRegistroProducto = requestBody.NumeroRegistroProducto

	if err := validation.Validate(deteccion.ToMap(), validation.DetecccionMicroorganismosRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	if err := controller.repo.ActualizarDeteccionMicroorganismos(deteccion); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar la deteccion de microorganismos", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos actualizada exitosamente"})
}

// Este handler nos permite obtener los registros de detecciones de microorganismos de nu producto
func (controller *deteccionMicroorganismosController) ObtenerDeteccionMicroorganismosPorProducto(c echo.Context) error {
	// Se obtiene el id del producto
	numeroRegistroProducto := c.Param("id")

	// Se verifica que el producto exista
	if _, err := controller.ProductoRepository.ObtenerInfoProducto(numeroRegistroProducto); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al obtener los datos", Error: err.Error()})
	}

	// Se obtienen los registros
	detecciones, err := controller.repo.ObtenerDeteccionMicroorganismosPorProducto(numeroRegistroProducto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener los registros", Error: err.Error()})
	}
	if len(detecciones) == 0 {
		return c.JSON(http.StatusNotFound, response.Response{Message: "No existen registros de detecciones de microorganismos para este producto"})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Registros obtenidos correctamente", Data: detecciones})
}

// Este handle nos permite eliminar un registro de deteccion de microorganismo
func (controller *deteccionMicroorganismosController) EliminarDeteccionMicroorganismos(c echo.Context) error {

	// Obtenemos el id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID invalido", Error: err.Error()})
	}

<<<<<<< HEAD
=======
	// Eliminamos el registro
>>>>>>> 0e5738afd739d5819881177926e021e06b6ed071
	if err := controller.repo.EliminarDeteccionMicroorganismos(id); err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Deteccion de microorganismos no encontrada", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Deteccion de microorganismos eliminada exitosamente"})
}
