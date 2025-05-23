package organization

import apperr "true-kw/errors"

type checkUseCase struct {
}

func NewCheckUseCase() *checkUseCase {
	return &checkUseCase{}
}

func (u *checkUseCase) Run(dta []Organization) ([]Organization, error) {
	op := "CheckUseCase.Run"
	if dta == nil {
		return nil, apperr.New(nil, apperr.ErrInternal, "data is empty", op)
	}
	return nil, nil
}
