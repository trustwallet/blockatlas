package config

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"reflect"
	"strings"
)

type Configuration struct {
	Gin struct {
		Mode         string `json:"mode"`
		ReverseProxy bool   `json:"reverse_proxy"`
	} `json:"gin"`
	Platform   []string `json:"platform"`
	RestAPI    string   `json:"rest_api"`
	SpamWords  []string `json:"spam_words"`
	Subscriber string   `json:"subscriber"`
	Observer   struct {
		Backlog                     string `json:"backlog"`
		FetchBlocksInterval         string `json:"fetch_blocks_interval"`
		BacklogMaxBlocks            int    `json:"backlog_max_blocks"`
		TxsBatchLimit               int    `json:"txs_batch_limit"`
		PushNotificationsBatchLimit int    `json:"push_notifications_batch_limit"`
		BlockPoll                   struct {
			Min string `json:"min"`
			Max string `json:"max"`
		} `json:"block_poll"`
		Rabbitmq struct {
			URL      string `json:"url"`
			Consumer struct {
				PrefetchCount int `json:"prefetch_count"`
			} `json:"consumer"`
		} `json:"rabbitmq"`
	} `json:"observer"`
	Postgres struct {
		URL  string `json:"url"`
		Read struct {
			URL string `json:"url"`
		} `json:"read"`
		Log bool `json:"log"`
	} `json:"postgres"`
	Binance struct {
		API      string `json:"api"`
		Explorer string `json:"explorer"`
	} `json:"binance"`
	Ripple struct {
		API string `json:"api"`
	} `json:"ripple"`
	Stellar struct {
		API string `json:"api"`
	} `json:"stellar"`
	Kin struct {
		API string `json:"api"`
	} `json:"kin"`
	Nimiq struct {
		API string `json:"api"`
	} `json:"nimiq"`
	Tezos struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"tezos"`
	Thundertoken struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"thundertoken"`
	Gochain struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"gochain"`
	Classic struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"classic"`
	Smartchain struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"smartchain"`
	BSC struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"bsc"`
	Poa struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"poa"`
	Callisto struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"callisto"`
	Wanchain struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"wanchain"`
	Tomochain struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
	} `json:"tomochain"`
	Ethereum struct {
		API            string `json:"api"`
		BlockbookAPI   string `json:"blockbook_api"`
		CollectionsAPI string `json:"collections_api"`
		CollectionsKey string `json:"collections_api_key"`
		RPC            string `json:"rpc"`
	} `json:"ethereum"`
	Aion struct {
		API string `json:"api"`
	} `json:"aion"`
	Icon struct {
		API string `json:"api"`
	} `json:"icon"`
	Tron struct {
		API      string `json:"api"`
		Explorer string `json:"explorer"`
	} `json:"tron"`
	Vechain struct {
		API string `json:"api"`
	} `json:"vechain"`
	Theta struct {
		API string `json:"api"`
	} `json:"theta"`
	Cosmos struct {
		API string `json:"api"`
	} `json:"cosmos"`
	Ontology struct {
		API string `json:"api"`
	} `json:"ontology"`
	Zilliqa struct {
		API string `json:"api"`
		RPC string `json:"rpc"`
		Key string `json:"key"`
	} `json:"zilliqa"`
	Iotex struct {
		API string `json:"api"`
	} `json:"iotex"`
	Waves struct {
		API string `json:"api"`
	} `json:"waves"`
	Aeternity struct {
		API string `json:"api"`
	} `json:"aeternity"`
	Nebulas struct {
		API string `json:"api"`
	} `json:"nebulas"`
	Fio struct {
		API string `json:"api"`
	} `json:"fio"`
	Bitcoin struct {
		API string `json:"api"`
	} `json:"bitcoin"`
	Litecoin struct {
		API string `json:"api"`
	} `json:"litecoin"`
	Bitcoincash struct {
		API string `json:"api"`
	} `json:"bitcoincash"`
	Doge struct {
		API string `json:"api"`
	} `json:"doge"`
	Dash struct {
		API string `json:"api"`
	} `json:"dash"`
	Zcoin struct {
		API string `json:"api"`
	} `json:"zcoin"`
	Zcash struct {
		API string `json:"api"`
	} `json:"zcash"`
	Zelcash struct {
		API string `json:"api"`
	} `json:"zelcash"`
	Viacoin struct {
		API string `json:"api"`
	} `json:"viacoin"`
	Qtum struct {
		API string `json:"api"`
	} `json:"qtum"`
	Groestlcoin struct {
		API string `json:"api"`
	} `json:"groestlcoin"`
	Ravencoin struct {
		API string `json:"api"`
	} `json:"ravencoin"`
	Decred struct {
		API string `json:"api"`
	} `json:"decred"`
	Algorand struct {
		API string `json:"api"`
	} `json:"algorand"`
	Nano struct {
		API string `json:"api"`
	} `json:"nano"`
	Digibyte struct {
		API string `json:"api"`
	} `json:"digibyte"`
	Harmony struct {
		API string `json:"api"`
	} `json:"harmony"`
	Kava struct {
		API string `json:"api"`
	} `json:"kava"`
	Kusama struct {
		API string `json:"api"`
	} `json:"kusama"`
	Polkadot struct {
		API string `json:"api"`
	} `json:"polkadot"`
	Solana struct {
		API string `json:"api"`
	} `json:"solana"`
	Near struct {
		API string `json:"api"`
	} `json:"near"`
	Elrond struct {
		API string `json:"api"`
	} `json:"elrond"`
	Filecoin struct {
		API string `json:"api"`
	} `json:"filecoin"`
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
			logger.Panic(err, "Fatal error reading default config")
		} else {
			logger.Info("Viper using default config", logger.Params{"config": viper.ConfigFileUsed()})
		}
	} else {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			logger.Panic(err, "Fatal error reading supplied config")
		} else {
			logger.Info("Viper using supplied config", logger.Params{"config": viper.ConfigFileUsed()})
		}
	}

	bindEnvs(c)
	if err := viper.Unmarshal(&c); err != nil {
		logger.Panic(err, "Error Unmarshal Viper Config File")
	}
	Default = c
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("json")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				logger.Fatal(err)
			}
		}
	}
}
