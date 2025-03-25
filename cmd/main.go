package main

import (
	"labgestor-server/infrastructure"
	"labgestor-server/internal/controllers"
	"labgestor-server/internal/repository"
	"labgestor-server/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := infrastructure.NewConexionDB()
	if err != nil {
		e.Logger.Fatal("Error al conectar a la base de datos", err)
	}

	//Repositorios
	usuarioRepo := repository.NewUsuarioRepository(db)
	clienteRepo := repository.NewClienterepository(db)

	// Controladores
	usuarioController := controllers.NewUsuarioController(usuarioRepo)
	clienteController := controllers.NewClienteController(clienteRepo)

	//Handlers para rutas
	routes.NewUsuarioHanlder(e, usuarioController)
	routes.NewClienteHandler(e, clienteController)

	//Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
