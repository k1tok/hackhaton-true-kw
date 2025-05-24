package v1

import (
	apperr "true-kw/errors"
	"true-kw/internal/core/organization"

	"github.com/gin-gonic/gin"
)

type checkHandler struct {
	useCase CheckUseCase
}

func NewCheckHandler(useCase CheckUseCase) Handler {
	return &checkHandler{
		useCase: useCase,
	}
}

func (h *checkHandler) Reg(r *gin.Engine) {
	r.POST("/check", h.Handle)
}

// Check godoc
//	@Summary		Check targets
//	@Description	checks targets and returns prediction
//	@Tags			check
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]organization.Organization	true	"request body"
//	@Success		200		{array}		organization.OrganizationResult
//	@Failure		500		{object}	HTTPError
//	@Router			/check [post]
func (h *checkHandler) Handle(c *gin.Context) {
	op := "CheckHandler.Handle"
	var json []organization.Organization

	if err := c.ShouldBindJSON(&json); err != nil {
		c.Error(apperr.New(err, apperr.ErrInternal, "invalid request", op))
		return
	}
	orgs, err := h.useCase.Run(json)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, orgs)
}
