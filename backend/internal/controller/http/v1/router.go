package v1

import "github.com/gin-gonic/gin"

func NewRouter(useCases *UseCases) *gin.Engine {
	r := gin.Default()
	r.Use(ErrorMiddleware())
	NewHealthHandler().Reg(r)
	NewCheckHandler(useCases.CheckUsecase).Reg(r)
	return r
}
