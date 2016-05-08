package bot

import (
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.bot")

// BotContext is the global bot context
type BotContext struct {
	Storage Storage
	Conf    *ConfService
	API     telegram.API
}

// CreateBotContext from a configuration file
func CreateBotContext(configurationFile string) (*BotContext, error) {
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

	return &BotContext{storage, conf, api}, nil
}
