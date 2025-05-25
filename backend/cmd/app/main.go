package main

import (
	_ "true-kw/docs"
	"true-kw/internal/app"
)

//	@title			Hakaton true kw
//	@version		1.0
//	@description	This is a project for the true kw library.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Flastor

//	@host		localhost:8080
//	@BasePath	/

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	c := app.NewContext()
	c.Run()
}
