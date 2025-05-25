package app

import (
	"true-kw/config"
	listorg "true-kw/internal/adapters/list-org/api"
	"true-kw/internal/adapters/osm"
	v1 "true-kw/internal/controller/http/v1"
	"true-kw/internal/core/organization"
)

type Context struct {
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) Run() {
	c.Server().Serve()
}

func (c *Context) Config() *config.Config {
	return config.Get()
}

func (c *Context) Server() *v1.Server {
	return v1.NewServer(c.UseCases(), c.Config().App)
}

// UseCases
func (c *Context) UseCases() *v1.UseCases {
	return &v1.UseCases{
		CheckUsecase: c.CheckUseCase(),
	}
}
func (c *Context) CheckUseCase() v1.CheckUseCase {
	return organization.NewCheckUseCase(c.ListOrgAPI(), c.OsmAPI())
}

// Adapters
func (c *Context) OsmAPI() organization.OsmAPI {
	return osm.NewOsmAPI(c.Config().API.OsmAPI)
}

func (c *Context) ListOrgAPI() organization.ListOrgAPI {
	return listorg.NewListOrgAPI(c.Config().API.ListOrgAPI)
}
