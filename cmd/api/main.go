package main

import (
	_ "github.com/edwinhuish/go-rest-template/docs"
	"github.com/edwinhuish/go-rest-template/internal/api"
)

//	@Golang			API REST
//	@version		1.0
//	@description	API REST in Golang with Gin Framework

//	@contact.name	Antonio Paya Gonzalez
//	@contact.url	http://antoniopg.tk
//	@contact.email	antonioalfa22@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/edwinhuish/go-rest-template/blob/master/LICENSE

//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	api.Run("")
}
