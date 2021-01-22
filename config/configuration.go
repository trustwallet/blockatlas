package config

import (
	"reflect"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Gin struct {
		Mode string `mapstructure:"mode"`
	} `mapstructure:"gin"`
	Platform []string `mapstructure:"platform"`
	RestAPI  string   `mapstructure:"rest_api"`
	Observer struct {
		FetchBlocksInterval time.Duration `mapstructure:"fetch_blocks_interval"`
		BlockPoll           struct {
			Min       time.Duration `mapstructure:"min"`
			Max       time.Duration `mapstructure:"max"`
			MaxBlocks int64         `mapstructure:"max_blocks"`
		} `mapstructure:"block_poll"`
		Rabbitmq struct {
			URL string `mapstructure:"url"`
		} `mapstructure:"rabbitmq"`
	} `mapstructure:"observer"`
	Postgres struct {
		URL string `mapstructure:"url"`
		Log bool   `mapstructure:"log"`
	} `mapstructure:"postgres"`
	Ethereum struct {
		API            string `mapstructure:"api"`
		BlockbookAPI   string `mapstructure:"blockbook_api"`
		CollectionsAPI string `mapstructure:"collections_api"`
		CollectionsKey string `mapstructure:"collections_api_key"`
		RPC            string `mapstructure:"rpc"`
	} `mapstructure:"ethereum"`
	Binance struct {
		API string `mapstructure:"api"`
		Key string `mapstructure:"key"`
	} `mapstructure:"binance"`
	Ripple struct {
		API string `mapstructure:"api"`
	} `mapstructure:"ripple"`
	Stellar struct {
		API string `mapstructure:"api"`
	} `mapstructure:"stellar"`
	Nimiq struct {
		API string `mapstructure:"api"`
	} `mapstructure:"nimiq"`
	Tezos struct {
		API   string `mapstructure:"api"`
		RPC   string `mapstructure:"rpc"`
		Baker string `mapstructure:"baker"`
	} `mapstructure:"tezos"`
	Thundertoken struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"thundertoken"`
	Gochain struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"gochain"`
	Classic struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"classic"`
	Smartchain struct {
		API            string `mapstructure:"api"`
		RPC            string `mapstructure:"rpc"`
		CollectionsAPI string `mapstructure:"collections_api"`
	} `mapstructure:"smartchain"`
	BSC struct {
		API            string `mapstructure:"api"`
		RPC            string `mapstructure:"rpc"`
		CollectionsAPI string `mapstructure:"collections_api"`
	} `mapstructure:"bsc"`
	Poa struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"poa"`
	Callisto struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"callisto"`
	Wanchain struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"wanchain"`
	Tomochain struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
	} `mapstructure:"tomochain"`
	Aion struct {
		API string `mapstructure:"api"`
	} `mapstructure:"aion"`
	Icon struct {
		API string `mapstructure:"api"`
	} `mapstructure:"icon"`
	Tron struct {
		API      string `mapstructure:"api"`
		Grid     string `mapstructure:"grid"`
		Explorer string `mapstructure:"explorer"`
	} `mapstructure:"tron"`
	Vechain struct {
		API string `mapstructure:"api"`
	} `mapstructure:"vechain"`
	Theta struct {
		API string `mapstructure:"api"`
	} `mapstructure:"theta"`
	Cosmos struct {
		API string `mapstructure:"api"`
	} `mapstructure:"cosmos"`
	Ontology struct {
		API string `mapstructure:"api"`
	} `mapstructure:"ontology"`
	Zilliqa struct {
		API string `mapstructure:"api"`
		RPC string `mapstructure:"rpc"`
		Key string `mapstructure:"key"`
	} `mapstructure:"zilliqa"`
	Iotex struct {
		API string `mapstructure:"api"`
	} `mapstructure:"iotex"`
	Waves struct {
		API string `mapstructure:"api"`
	} `mapstructure:"waves"`
	Aeternity struct {
		API string `mapstructure:"api"`
	} `mapstructure:"aeternity"`
	Nebulas struct {
		API string `mapstructure:"api"`
	} `mapstructure:"nebulas"`
	Fio struct {
		API string `mapstructure:"api"`
	} `mapstructure:"fio"`
	Bitcoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"bitcoin"`
	Litecoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"litecoin"`
	Bitcoincash struct {
		API string `mapstructure:"api"`
	} `mapstructure:"bitcoincash"`
	Doge struct {
		API string `mapstructure:"api"`
	} `mapstructure:"doge"`
	Dash struct {
		API string `mapstructure:"api"`
	} `mapstructure:"dash"`
	Zcoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"zcoin"`
	Zcash struct {
		API string `mapstructure:"api"`
	} `mapstructure:"zcash"`
	Zelcash struct {
		API string `mapstructure:"api"`
	} `mapstructure:"zelcash"`
	Viacoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"viacoin"`
	Qtum struct {
		API string `mapstructure:"api"`
	} `mapstructure:"qtum"`
	Groestlcoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"groestlcoin"`
	Ravencoin struct {
		API string `mapstructure:"api"`
	} `mapstructure:"ravencoin"`
	Decred struct {
		API string `mapstructure:"api"`
	} `mapstructure:"decred"`
	Algorand struct {
		API string `mapstructure:"api"`
		Key string `mapstructure:"key"`
	} `mapstructure:"algorand"`
	Nano struct {
		API string `mapstructure:"api"`
	} `mapstructure:"nano"`
	Digibyte struct {
		API string `mapstructure:"api"`
	} `mapstructure:"digibyte"`
	Harmony struct {
		API string `mapstructure:"api"`
	} `mapstructure:"harmony"`
	Kava struct {
		API string `mapstructure:"api"`
	} `mapstructure:"kava"`
	Kusama struct {
		API string `mapstructure:"api"`
	} `mapstructure:"kusama"`
	Polkadot struct {
		API string `mapstructure:"api"`
	} `mapstructure:"polkadot"`
	Solana struct {
		API string `mapstructure:"api"`
	} `mapstructure:"solana"`
	Near struct {
		API string `mapstructure:"api"`
	} `mapstructure:"near"`
	Elrond struct {
		API string `mapstructure:"api"`
	} `mapstructure:"elrond"`
	Filecoin struct {
		API      string `mapstructure:"api"`
		Explorer string `mapstructure:"explorer"`
	} `mapstructure:"filecoin"`
	Sentry struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"sentry"`
	Consumer struct {
		Service string `mapstructure:"service"`
		Workers int    `mapstructure:"workers"`
	} `mapstructure:"consumer"`
}

var Default Configuration

func Init(confPath string) {
	c := Configuration{}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if confPath == "" {
		err := viper.ReadInConfig()
		if err != nil {
			log.Panic(err, "Fatal error reading default config")
		} else {
			log.WithFields(log.Fields{"config": viper.ConfigFileUsed()}).Info("Viper using default config")
		}
	} else {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panic(err, "Fatal error reading supplied config")
		} else {
			log.WithFields(log.Fields{"config": viper.ConfigFileUsed()}).Info("Viper using supplied config")
		}
	}

	bindEnvs(c)
	if err := viper.Unmarshal(&c); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config File")
	}
	Default = c
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				log.Fatal(err)
			}
		}
	}
}
