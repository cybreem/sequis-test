package infrastucture

import (
	"os"
	"strings"

	"github.com/cybreem/sequis-test/app"
	"github.com/h4lim/go-sdk/config"
	"github.com/h4lim/go-sdk/logging"
	"github.com/spf13/viper"
)

func SetConfigFile() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	configFile := "config.toml"
	if err := config.Initialize(configFile); err != nil {
		app.Log.Fatalf(logging.INTERNAL, "Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}
	setting := viper.AllSettings()
	for h, v := range setting {
		app.Log.Debugf("INTERNAL", " %v %v", h, v)
		if rec, ok := v.(map[string]interface{}); ok {
			for key := range rec {
				err := viper.BindEnv(h + "_" + key)
				if err != nil {
					app.Log.Debugf("INTERNAL", "Failed to bind env \"%s\" into configuration. Got %s", key, err)
				}
			}
		} else {
			app.Log.Debugf("INTERNAL", "record not a map[string]interface{}: %v\n", v)
		}
	}
}
