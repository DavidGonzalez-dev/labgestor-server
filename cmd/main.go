package main

import (
	"labgestor-server/infrastructure"
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/internal/routes"
	"labgestor-server/utils/initialization"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	// Cargar Variables de entorno
	utils.LoadEnvVariables()
	
	// Servidor Echo
	e := echo.New()

	// Conexion a la base de datos
	db, err := infrastructure.NewConexionDB()
	if err != nil {
		e.Logger.Fatal("Error al conectar a la base de datos", err)
	}

	//Repositorios
	usuarioRepo := repository.NewUsuarioRepository(db)
	clienteRepo := repository.NewClienterepository(db)
	fabricanteRepo := repository.NewFabricanterepository(db)

	// Controladores
	usuarioController := controllers.NewUsuarioController(usuarioRepo)
	clienteController := controllers.NewClienteController(clienteRepo)
	fabricanteController := controllers.NewFabricanteController(fabricanteRepo)

	//Handlers para rutas
	routes.NewUsuarioHanlder(e, usuarioController, usuarioRepo)
	routes.NewClienteHandler(e, clienteController)
	routes.NewFabricanteHandler(e, fabricanteController)

	//Iniciar servidor
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
