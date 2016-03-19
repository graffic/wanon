package bot

import (
	"net/http"
	"time"

	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.bot")

// Context bot context
type Context struct {
	Storage Storage
	Conf    *ConfService
	API     telegram.API
}

func createAPI(conf *ConfService) (telegram.API, error) {
	var apiConf telegram.Configuration
	conf.Get(&apiConf)
	api := telegram.NewAPI(&http.Client{Timeout: time.Second * 15}, apiConf)
	result, err := api.GetMe()
	if err != nil {
		return nil, err
	}
	log.Info("%s online", result.Username)

	return api, nil
}

// CreateContext creates a bot context from a configuration file
func CreateContext(configurationFile string) (*Context, error) {
	conf, err := LoadConf(configurationFile)
	if err != nil {
		return nil, err
	}

	storage, err := NewStorage(conf)
	if err != nil {
		return nil, err
	}

	api, err := createAPI(conf)
	if err != nil {
		return nil, err
	}

	return &Context{storage, conf, api}, nil
}
