package router

import "github.com/gin-gonic/gin"

// imports e exports, para qualquer variavel, funcao, tipo de entidade que comece com letra maiuscula esta sendo exportada do seu package
// inicio a funcao com letra maiuscula para que possa ser utilizada em outro lugar
// visibilidade publica
func Initialize() {
	// Inicializa o Router utilizando as configurações Default do Gin
	router := gin.Default()
	
	// Inicializa as rotas
	initializeRoutes(router)

	// estamos rodando nossa API, por padrão roda na porta 8080
	router.Run()
}