package main

import (
	"labgestor-server/infrastructure"
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/internal/routes"
	utils "labgestor-server/utils/initialization"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Cargar Variables de entorno
	utils.LoadEnvVariables()

	// Servidor Echo
	e := echo.New()

	// Configuracion para evitar errores CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4321"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// Conexion a la base de datos
	db, err := infrastructure.NewConexionDB()
	if err != nil {
		e.Logger.Fatal("Error al conectar a la base de datos", err)
	}

	//Repositorios
	usuarioRepo := repository.NewUsuarioRepository(db)
	clienteRepo := repository.NewClienterepository(db)
	fabricanteRepo := repository.NewFabricanterepository(db)
	productoRepo := repository.NewProductoRepository(db)
	pruebaRecuentoRepo := repository.NewPruebaRecuentoRepository(db)

	// Controladores
	usuarioController := controllers.NewUsuarioController(usuarioRepo)
	clienteController := controllers.NewClienteController(clienteRepo)
	fabricanteController := controllers.NewFabricanteController(fabricanteRepo)
	productoController := controllers.NewProductoController(productoRepo)
	pruebaRecuentoController := controllers.NewPruebaRecuentoController(pruebaRecuentoRepo)

	//Handlers para rutas
	routes.NewUsuarioHanlder(e, usuarioController, usuarioRepo)
	routes.NewClienteHandler(e, clienteController)
	routes.NewFabricanteHandler(e, fabricanteController)
	routes.NewProductoHandler(e, productoController)
	routes.NewPruebaRecuentoHandler(e, pruebaRecuentoController)

	//Iniciar servidor
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
