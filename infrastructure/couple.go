package infrastucture

import (
	"net/http"
	"time"

	"github.com/cybreem/sequis-test/app"
	"github.com/h4lim/go-sdk/app/types"
	"github.com/h4lim/go-sdk/caching"
	"github.com/h4lim/go-sdk/config"
	"github.com/h4lim/go-sdk/database"
	"github.com/h4lim/go-sdk/logging"
	"github.com/h4lim/go-sdk/utils"
)

func SetBundlingComponent() ComponentContext {
	SetConfigFile()
	component := ComponentContext{
		MessageEN:   setMessageEN(),
		MessageID:   setMessageID(),
		SimpleCache: setCaching(),
		DBModel:     setDBModel(),
		Server:      setServer(),
	}
	return component
}

func setMessageEN() types.MessageMap {
	messageMapEN := logging.InitializeMaps(config.MustGetString("server.maps"), app.MESSAGE_EN)
	return *messageMapEN
}

func setMessageID() types.MessageMap {
	messageMapID := logging.InitializeMaps(config.MustGetString("server.maps"), app.MESSAGE_ID)
	return *messageMapID
}

func setCaching() caching.SimpleCache {
	simpleCache := caching.SimpleCache{
		ExpiredTime: config.MustGetInt("server.cache_expired"),
		PurgeTime:   config.MustGetInt("server.cache_purged"),
	}
	return simpleCache
}

func setDBModel() database.DBModel {
	dbModel := database.DBModel{
		ServerMode: config.MustGetString("server.mode"),
		Driver:     config.MustGetString(utils.GetRunMode() + ".db_driver"),
		Host:       config.MustGetString(utils.GetRunMode() + ".db_host"),
		Port:       config.MustGetString(utils.GetRunMode() + ".db_port"),
		Name:       config.MustGetString(utils.GetRunMode() + ".db_name"),
		Username:   config.MustGetString(utils.GetRunMode() + ".db_username"),
		Password:   config.MustGetString(utils.GetRunMode() + ".db_password"),
	}
	return dbModel
}

func setServer() http.Server {
	server := http.Server{
		Addr:         ":" + config.MustGetString("server.port"),
		Handler:      initRoutes(config.MustGetBool("server.debug")),
		ReadTimeout:  time.Minute * time.Duration(config.MustGetInt("server.http_timeout")),
		WriteTimeout: time.Minute * time.Duration(config.MustGetInt("server.http_timeout")),
	}
	return server
}
