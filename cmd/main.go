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

	// Controladores
	usuarioController := controllers.NewUsuarioController(usuarioRepo)

	//Handlers para rutas
	routes.NewUsuarioHanlder(e, usuarioController)

	//Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}