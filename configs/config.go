package configs

import (
	"flag"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"

)

var (
	debug bool

	confPath string

	Conf *Config
)

func init() {

	//defHost, _  = os.Hostname()
	//defAddrs    = os.Getenv("ADDRS")
	var defDebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	flag.BoolVar(&debug, "debug", defDebug, "server debug. or use DEBUG env variable, value: true/false etc.")

	flag.StringVar(&confPath, "conf", "test.toml", "default config path.")

}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)

	glog.Infof(" config Web %#v \n", Conf.Web)

	return
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{
		Debug:           debug,
		StartHttp:       true,

		Web: &WebConfig{
			Listen:        ":1323",
			AccessLogPath: "./log/access.log",
		},



	}
}

// Config is grpc config.
type Config struct {
	Debug bool

	StartHttp bool // 开启http 和接收服务


	Web *WebConfig

}

type WebConfig struct {
	Listen        string
	AccessLogPath string
}

