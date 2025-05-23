package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() Handler {
	return &HealthHandler{}
}

func (h *HealthHandler) Reg(r *gin.Engine) {
	r.GET("/health", h.Handle)
}

func (h *HealthHandler) Handle(c *gin.Context) {
	c.Status(http.StatusOK)
}
