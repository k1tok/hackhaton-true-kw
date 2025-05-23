package v1

import (
	"true-kw/internal/core/organization"

	"github.com/gin-gonic/gin"
)

// handler
type Handler interface {
	Reg(r *gin.Engine)
}

// UseCases
type CheckUseCase interface {
	Run(dta []organization.Organization) ([]organization.Organization, error)
}
