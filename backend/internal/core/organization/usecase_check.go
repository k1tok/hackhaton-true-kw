package organization

import (
	"fmt"
	apperr "true-kw/errors"
)

type checkUseCase struct {
	listOrgAPI ListOrgAPI
	osmAPI     OsmAPI
}

func NewCheckUseCase(listOrgAPI ListOrgAPI, osmAPI OsmAPI) *checkUseCase {
	return &checkUseCase{
		listOrgAPI: listOrgAPI,
		osmAPI:     osmAPI,
	}
}

func (u *checkUseCase) Run(dta []Organization) ([]OrganizationResult, error) {
	op := "CheckUseCase.Run"
	if dta == nil {
		return nil, apperr.New(nil, apperr.ErrInternal, "data is empty", op)
	}
	for _, org := range dta {
		osmOrgs, err := u.osmAPI.Search(org.Address)
		if err != nil {
			if !apperr.IsErrorCode(err, apperr.ErrNotFound) {
				return nil, err
			}
		}
		fmt.Println(len(osmOrgs))
		listOrgs, err := u.listOrgAPI.Search(org.Address)
		if err != nil {
			return nil, err
		}
		fmt.Println(len(listOrgs))
	}
	return nil, nil
}
