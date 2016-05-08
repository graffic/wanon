package bot

import (
	"github.com/graffic/goejdb"
	"github.com/graffic/wanon/telegram"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("wanon.bot")

// Context is the global bot context
type Context struct {
	Storage *goejdb.Ejdb
	Conf    *ConfService
	API     telegram.API
}

// CreateContext for the bot from a configuration file
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
