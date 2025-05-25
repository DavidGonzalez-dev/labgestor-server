package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"labgestor-server/utils/validation"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ControlesNegativosController interface {
	CrearControlesNegativos(c echo.Context) error
	ObtenerControlesNegativosID(c echo.Context) error
	ObtenerControlesPorProducto(c echo.Context) error
	ActualizarControlesNegativos(c echo.Context) error
	EliminarControlesNegativos(c echo.Context) error
}

type controlesNegativosController struct {
	repo repository.ControlesNegativosRepository
}

// Funcion para instanciar un controlador
func NewControlesNegativosController(repo repository.ControlesNegativosRepository) ControlesNegativosController {
	return &controlesNegativosController{repo: repo}
}

// ? ------------------------------------------------
// ? CONTROLADORES CRUD
// ? ------------------------------------------------
func (controller controlesNegativosController) CrearControlesNegativos(c echo.Context) error {
	// Se crea el producto y se verifica que no hallan errores
	var requestBody struct {
		MedioCultivo         string    `json:"medioCultivo"`
		FechayhoraIncubacion time.Time `json:"fechayhoraIncubacion"`
		FechayhoraLectura    time.Time `json:"fechayhoraLectura"`
		Resultado            string    `json:"resultado"`
		NumeroRegistro       string    `json:"numeroRegistro"`
	}

	// Se verifica que el cuerpo del request no este vacio
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	controlesNegativos := models.ControlesNegativosMedio{
		MedioCultivo:           requestBody.MedioCultivo,
		FechayhoraIncubacion:   requestBody.FechayhoraIncubacion,
		FechayhoraLectura:      requestBody.FechayhoraLectura,
		Resultado:              requestBody.Resultado,
		NumeroRegistroProducto: requestBody.NumeroRegistro,
	}

	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
	if err := validation.Validate(controlesNegativos.ToMap(), validation.ControlesNegativosRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}
	if err := controller.repo.CrearControlesNegativos(&controlesNegativos); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro creado correctamente"})

}
func (controller controlesNegativosController) ObtenerControlesNegativosID(c echo.Context) error {
	// Se obtiene el id del producto
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: err.Error()})
	}

	// Se obtiene el producto y se verifica que no hallan errores
	controlesNegativos, err := controller.repo.ObtenerControlesNegativosID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro obtenido correctamente", Data: controlesNegativos.ToMap()})
}

func (controller controlesNegativosController) ObtenerControlesPorProducto(c echo.Context) error {
	// Se obtiene el id del producto
	id := c.Param("id")

	// Se obtiene el producto y se verifica que no hallan errores
	controlesNegativos, err := controller.repo.ObtenerControlesPorProducto(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener el registro", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Registro obtenido correctamente", Data: controlesNegativos})
}

func (controller controlesNegativosController) ActualizarControlesNegativos(c echo.Context) error {

	// Se crea el producto y se verifica que no hallan errores
	var requestBody struct {
		ID                     int       `json:"id"`
		MedioCultivo           string    `json:"medioCultivo"`
		FechayhoraIncubacion   time.Time `json:"fechayhoraIncubacion"`
		FechayhoraLectura      time.Time `json:"fechayhoraLectura"`
		Resultado              string    `json:"resultado"`
		NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}
	controlesNegativos := models.ControlesNegativosMedio{
		ID:                     requestBody.ID,
		MedioCultivo:           requestBody.MedioCultivo,
		FechayhoraIncubacion:   requestBody.FechayhoraIncubacion,
		FechayhoraLectura:      requestBody.FechayhoraLectura,
		Resultado:              requestBody.Resultado,
		NumeroRegistroProducto: requestBody.NumeroRegistroProducto,
	}

	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
	if err := validation.Validate(controlesNegativos.ToMap(), validation.ControlesNegativosRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	if err := controller.repo.ActualizarControlesNegativos(&controlesNegativos); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro actualizado correctamente"})
}

func (controller controlesNegativosController) EliminarControlesNegativos(c echo.Context) error {
	// Se obtiene el id del producto
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "ID inválido", Error: err.Error()})
	}

	// Se obtiene el producto y se verifica que no hallan errores
	if err := controller.repo.EliminarControlesNegativos(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al eliminar el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro eliminado correctamente"})
}
