package configs

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"wps.ktkt.com/app2017/elastic_query-test/internal/queue/channel"
	"wps.ktkt.com/app2017/elastic_query-test/internal/queue/nsq"

	"wps.ktkt.com/app2017/elastic_query-test/internal/server/selfMiddleware"
	selfTime "wps.ktkt.com/app2017/elastic_query-test/pkg/time"
)

var (
	debug bool

	confPath string

	//lookUpdAddress  []string
	//isStartCustomer bool

	//topic string = "user_event"

	// Conf config
	Conf *Config
)

func init() {

	//defHost, _  = os.Hostname()
	//defAddrs    = os.Getenv("ADDRS")
	var defDebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	flag.BoolVar(&debug, "debug", defDebug, "server debug. or use DEBUG env variable, value: true/false etc.")

	flag.StringVar(&confPath, "conf", "test.toml", "default config path.")

	//flag.BoolVar(&isStartCustomer, "start_customer", false, "是否启动消费服务")
	//
	//// lookup address
	//flag.Var(extend.NewSliceValue([]string{}, &lookUpdAddress), "lookupd_addr", "nsqlookupd 地址 : host1:4171,host2:4171 .")
	////flag.StringVar(&topic, "topic", "user_event", "nsq topic")

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

