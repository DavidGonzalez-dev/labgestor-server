package controllers

import (
	"labgestor-server/internal/models"
	"labgestor-server/internal/repository"
	"labgestor-server/utils/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ControlesNegativosController interface {
	CrearControlesNegativos(c echo.Context) error
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

	if err := controller.repo.CrearControlesNegativos(&controlesNegativos); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro creado correctamente"})

	// Se verifica que el producto no exista
}
