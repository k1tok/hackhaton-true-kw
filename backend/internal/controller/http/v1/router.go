package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(useCases *UseCases) *gin.Engine {
	r := gin.Default()
	r.Use(ErrorMiddleware())
	NewHealthHandler().Reg(r)
	NewCheckHandler(useCases.CheckUsecase).Reg(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
