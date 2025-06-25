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

// Interfaz que define los metodos que debe implementar el controlador de controles negativos
type ControlesNegativosController interface {
	CrearControlesNegativos(c echo.Context) error
	ObtenerControlesNegativosID(c echo.Context) error
	ObtenerControlesPorProducto(c echo.Context) error
	ActualizarControlesNegativos(c echo.Context) error
	EliminarControlesNegativos(c echo.Context) error
}

type controlesNegativosController struct {
	repo         repository.ControlesNegativosRepository
	ProductoRepo repository.ProductoRepository
}

// Funcion para instanciar un controlador
func NewControlesNegativosController(repo repository.ControlesNegativosRepository, productoRepo repository.ProductoRepository) ControlesNegativosController {
	return &controlesNegativosController{repo: repo, ProductoRepo: productoRepo}
}

// ? ------------------------------------------------
// ? CONTROLADORES CRUD
// ? ------------------------------------------------
// Este controlador nos permite crear un registro de los controles negativos en la base de datos
func (controller controlesNegativosController) CrearControlesNegativos(c echo.Context) error {
	// Se crea el producto y se verifica que no hallan errores
	var requestBody struct {
		MedioCultivo           string    `json:"medioCultivo"`
		FechayhoraIncubacion   time.Time `json:"fechayhoraIncubacion"`
		FechayhoraLectura      time.Time `json:"fechayhoraLectura"`
		Resultado              string    `json:"resultado"`
		NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
	}

	// Se verifica que el cuerpo del request no este vacio
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	// Se verifica que el las fechas de incubacion y lectura sean validas
	if requestBody.FechayhoraIncubacion.IsZero() || requestBody.FechayhoraLectura.IsZero() {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "Las fechas de incubacion y lectura son requeridas"})
	}

	controlesNegativos := models.ControlesNegativosMedio{
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
	// Se verifica que el producto exista
	producto, _ := controller.ProductoRepo.ObtenerInfoProducto(requestBody.NumeroRegistroProducto)
	if producto == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Error al crear la prueba re recuento", Error: "El producto al que le estas asignando la prueba de recuento no existe"})
	}

	// Se crea la prueba de recuento si el producto aun tiene un estado diferente a "terminado"(3)
	if producto.IDEstado != 3 {

		// Se crea la prueba de recuento
		if err := controller.repo.CrearControlesNegativos(&controlesNegativos); err != nil {
			return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear la control negativo", Error: err.Error()})
		}

		if err := controller.ProductoRepo.ActualizarEstadoProducto(2, requestBody.NumeroRegistroProducto); err != nil {
			return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al crear la control negativo", Error: err.Error()})
		}

		return c.JSON(http.StatusCreated, response.Response{Message: "Control negativo creada correctamente"})
	}

	return c.JSON(http.StatusBadRequest, response.Response{Message: "Error al crear el control negativo", Error: "El producto ya ha sido terminado, no se pueden crear controles negativos para este producto"})
}

// Este controlador nos permite obtener un registro de los controles negativos en la base de datos
func (controller controlesNegativosController) ObtenerControlesNegativosID(c echo.Context) error {

	// Se obtiene el id del producto y se convierte a entero
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id tiene que ser un numero entero", Error: err.Error()})
	}

	// Se obtiene el producto y se verifica que no hallan errores
	controlesNegativos, err := controller.repo.ObtenerControlesNegativosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado"})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro obtenido correctamente", Data: controlesNegativos})
}

// Este controlador nos permite obtener los registro de los controles negativos por producto en la base de datos
func (controller controlesNegativosController) ObtenerControlesPorProducto(c echo.Context) error {

	// Se obtiene el id del producto
	id := c.Param("id")

	// Se verifica que exista el producto
	if producto, _ := controller.ProductoRepo.ObtenerProductoID(id); producto == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Este producto no existe"})
	}


	// Se obtiene el producto y se verifica que no hallan errores
	controlesNegativos, err := controller.repo.ObtenerControlesPorProducto(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al obtener los registros", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Registros obtenidos correctamente", Data: controlesNegativos})
}

// Este controlador nos permite actualizar un registro de los controles negativos en la base de datos
func (controller controlesNegativosController) ActualizarControlesNegativos(c echo.Context) error {
	// Se obtiene el id del producto y se convierte a entero
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id tiene que ser un numero entero", Error: err.Error()})
	}

	// Se verifica que el registro exista
	controlesNegativos, err := controller.repo.ObtenerControlesNegativosID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado", Error: err.Error()})
	}

	// Se lee el cuerpo del request
	var requestBody struct {
		MedioCultivo           string    `json:"medioCultivo"`
		FechayhoraIncubacion   time.Time `json:"fechayhoraIncubacion"`
		FechayhoraLectura      time.Time `json:"fechayhoraLectura"`
		Resultado              string    `json:"resultado"`
		NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
	}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Error al leer el cuerpo del request", Error: err.Error()})
	}

	controlesNegativos.MedioCultivo = requestBody.MedioCultivo
	controlesNegativos.FechayhoraIncubacion = requestBody.FechayhoraIncubacion
	controlesNegativos.FechayhoraLectura = requestBody.FechayhoraLectura
	controlesNegativos.Resultado = requestBody.Resultado


	//? ------------------------------------------------
	//? Se hace la validacion de los campos
	//? ------------------------------------------------
	if err := validation.Validate(controlesNegativos.ToMap(), validation.ControlesNegativosRules); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{Message: "Informacion con formato erroneo", Error: err.Error()})
	}

	// Se actualiza el registro con los nuevos datos
	if err := controller.repo.ActualizarControlesNegativos(controlesNegativos); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al actualizar el registro", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Response{Message: "Registro actualizado correctamente"})
}

// Este controlador nos permite eliminar un registro de los controles negativos en la base de datos
func (controller controlesNegativosController) EliminarControlesNegativos(c echo.Context) error {
	// Se obtiene el id del producto
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Message: "El id debe ser un numero entero", Error: err.Error()})
	}

	// Se verifica que el registro exista
	if controlesNegativos, _ := controller.repo.ObtenerControlesNegativosID(id); controlesNegativos == nil {
		return c.JSON(http.StatusNotFound, response.Response{Message: "Registro no encontrado"})
	}

	// Se hace la eliminacion del registro
	if err := controller.repo.EliminarControlesNegativos(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{Message: "Error al eliminar el registro", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Response{Message: "Registro eliminado correctamente"})
}
