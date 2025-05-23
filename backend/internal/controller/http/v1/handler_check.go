package v1

import "github.com/gin-gonic/gin"

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

func (h *checkHandler) Handle(c *gin.Context) {

	orgs, err := h.useCase.Run(nil)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, orgs)
}
