package infrastucture

import (
	"net/http"
	"os"

	"github.com/cybreem/sequis-test/app"
	// "github.com/cybreem/sequis-test/app/deliveries"
	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/h4lim/go-sdk/app/types"
	"github.com/h4lim/go-sdk/caching"
	"github.com/h4lim/go-sdk/config"
	"github.com/h4lim/go-sdk/database"
	"github.com/h4lim/go-sdk/logging"
	"github.com/h4lim/go-sdk/utils"
	"golang.org/x/sync/errgroup"
)

type ComponentContext struct {
	MessageEN   types.MessageMap
	MessageID   types.MessageMap
	SimpleCache caching.SimpleCache
	DBModel     database.DBModel
	Server      http.Server
}

type ComponentBundling interface {
	Messaging()
	Caching()
	Migration()
	WebServer()
}

func Bundling(MessageEN types.MessageMap,
	MessageID types.MessageMap,
	SimpleCache caching.SimpleCache, DBModel database.DBModel,
	Server http.Server) ComponentBundling {

	return ComponentContext{
		MessageEN:   MessageEN,
		MessageID:   MessageID,
		SimpleCache: SimpleCache,
		DBModel:     DBModel,
		Server:      Server,
	}
}

func (c ComponentContext) Messaging() {
	utils.MMEN = &c.MessageEN
	utils.MMID = &c.MessageID
}

func (c ComponentContext) Caching() {

	caching.GlobalCache = c.SimpleCache.CreateNewCache()
	if caching.GlobalCache == nil {
		app.Log.Fatalf(logging.INTERNAL, "Error create new cache: %s\n")
		os.Exit(1)
	}
	app.SimpleCache = c.SimpleCache

}

func (c ComponentContext) Migration() {
	db, err := c.DBModel.InitDB()
	if err != nil {
		app.Log.Fatalf(logging.INTERNAL, "Error reading database configuration: %s\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.FriendRequests{})
	db.AutoMigrate(&models.Friendlists{})
	utils.DBModel = &c.DBModel

	// orderAppsDelivery := deliveries.NewOrderAppsDelivery()
	// orderAppsDelivery.ImportRestaurant("migration/restaurants.json")
}

func (c ComponentContext) WebServer() {
	var groupRouter errgroup.Group
	app.Log.Infof(logging.INTERNAL, "Running in mode %s", config.MustGetString("server.mode"))
	app.Log.Infof(logging.INTERNAL, "Running on port %d", config.MustGetInt("server.port"))
	app.Log.Infof(logging.INTERNAL, "Serving document directory %s", config.MustGetString("server.docs"))
	app.Log.Infof(logging.INTERNAL, "Debug logging: %t", config.MustGetBool("server.debug"))

	groupRouter.Go(func() error {
		return c.Server.ListenAndServe()
	})

	if err := groupRouter.Wait(); err != nil {
		app.Log.Fatalf(logging.INTERNAL, "Error: %s", err.Error())
	}

}
