package main

import (
    "inventory-backend/config"
    "inventory-backend/routes"
    "github.com/gin-gonic/gin"
	"inventory-backend/seed"
)

func main() {
    // Inicializar o banco de dados
    config.ConnectDatabase()

	seed.LoadSeedData()

    // Inicializar o servidor Gin
    r := gin.Default()

    // Configurar rotas
    routes.SetupRoutes(r)

    // Rodar o servidor
    r.Run(":8080")
}

 

