package main

import infrastucture "github.com/cybreem/sequis-test/infrastructure"

//	@title			Sequis - Test
//	@version		1.0
//	@description	Api swagger collection for Sequis - Test.

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@host		127.0.0.1:8080
//	@BasePath	/
//	@schemes	http
func main() {

	component := infrastucture.SetBundlingComponent()
	bundling := infrastucture.Bundling(component.MessageEN, component.MessageID,
		component.SimpleCache, component.DBModel, component.Server)
	bundling.Messaging()
	bundling.Caching()
	bundling.Migration()
	bundling.WebServer()

}
