package platform

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/cosmos"
	"github.com/trustwallet/blockatlas/platform/polkadot"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/platform/aeternity"
	"github.com/trustwallet/blockatlas/platform/aion"
	"github.com/trustwallet/blockatlas/platform/algorand"
	"github.com/trustwallet/blockatlas/platform/binance"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"github.com/trustwallet/blockatlas/platform/ethereum"
	"github.com/trustwallet/blockatlas/platform/fio"
	"github.com/trustwallet/blockatlas/platform/harmony"
	"github.com/trustwallet/blockatlas/platform/icon"
	"github.com/trustwallet/blockatlas/platform/iotex"
	"github.com/trustwallet/blockatlas/platform/nano"
	"github.com/trustwallet/blockatlas/platform/nebulas"
	"github.com/trustwallet/blockatlas/platform/nimiq"
	"github.com/trustwallet/blockatlas/platform/ontology"
	"github.com/trustwallet/blockatlas/platform/ripple"
	"github.com/trustwallet/blockatlas/platform/stellar"
	"github.com/trustwallet/blockatlas/platform/tezos"
	"github.com/trustwallet/blockatlas/platform/theta"
	"github.com/trustwallet/blockatlas/platform/tron"
	"github.com/trustwallet/blockatlas/platform/vechain"
	"github.com/trustwallet/blockatlas/platform/waves"
	"github.com/trustwallet/blockatlas/platform/zilliqa"
)

var (
	allPlatformsList = []blockatlas.Platform{
		&binance.Platform{},
		&nimiq.Platform{},
		&ripple.Platform{},
		&stellar.Platform{CoinIndex: coin.XLM},
		&stellar.Platform{CoinIndex: coin.KIN},
		&ethereum.Platform{CoinIndex: coin.ETH},
		&ethereum.Platform{CoinIndex: coin.ETC},
		&ethereum.Platform{CoinIndex: coin.POA},
		&ethereum.Platform{CoinIndex: coin.CLO},
		&ethereum.Platform{CoinIndex: coin.GO},
		&ethereum.Platform{CoinIndex: coin.WAN},
		&ethereum.Platform{CoinIndex: coin.TOMO},
		&ethereum.Platform{CoinIndex: coin.TT},
		&cosmos.Platform{CoinIndex: coin.ATOM},
		&cosmos.Platform{CoinIndex: coin.KAVA},
		&tezos.Platform{},
		&aion.Platform{},
		&icon.Platform{},
		&iotex.Platform{},
		&ontology.Platform{},
		&theta.Platform{},
		&tron.Platform{},
		&vechain.Platform{},
		&zilliqa.Platform{},
		&waves.Platform{},
		&aeternity.Platform{},
		&bitcoin.Platform{CoinIndex: coin.BTC},
		&bitcoin.Platform{CoinIndex: coin.LTC},
		&bitcoin.Platform{CoinIndex: coin.BCH},
		&bitcoin.Platform{CoinIndex: coin.DASH},
		&bitcoin.Platform{CoinIndex: coin.DOGE},
		&bitcoin.Platform{CoinIndex: coin.ZEC},
		&bitcoin.Platform{CoinIndex: coin.XZC},
		&bitcoin.Platform{CoinIndex: coin.VIA},
		&bitcoin.Platform{CoinIndex: coin.RVN},
		&bitcoin.Platform{CoinIndex: coin.QTUM},
		&bitcoin.Platform{CoinIndex: coin.GRS},
		&bitcoin.Platform{CoinIndex: coin.ZEL},
		&bitcoin.Platform{CoinIndex: coin.DCR},
		&bitcoin.Platform{CoinIndex: coin.DGB},
		&nebulas.Platform{},
		&fio.Platform{},
		&algorand.Platform{},
		&nano.Platform{},
		&harmony.Platform{},
		&polkadot.Platform{CoinIndex: coin.KSM},
	}

	// Platforms contains all registered platforms by handle
	Platforms map[string]blockatlas.Platform

	// BlockAPIs contain platforms with block services
	BlockAPIs map[string]blockatlas.BlockAPI

	// StakeAPIs contain platforms with staking services
	StakeAPIs map[string]blockatlas.StakeAPI

	// CustomAPIs contain platforms with custom HTTP services
	CustomAPIs map[string]blockatlas.CustomAPI

	// NamingAPIs contain platforms which support naming services
	NamingAPIs map[uint64]blockatlas.NamingServiceAPI

	// CollectionAPIs contain platforms which collections services
	CollectionAPIs map[uint]blockatlas.CollectionAPI
)

func getPlatform(handle string) blockatlas.Platform {
	switch handle {
	case coin.Binance().Handle:
		return &binance.Platform{}
	case coin.Nimiq().Handle:
		return &nimiq.Platform{}
	case coin.Ripple().Handle:
		return &ripple.Platform{}
	case coin.Stellar().Handle:
		return &stellar.Platform{CoinIndex: coin.XLM}
	case coin.Kin().Handle:
		return &stellar.Platform{CoinIndex: coin.KIN}
	case coin.Ethereum().Handle:
		return &ethereum.Platform{CoinIndex: coin.ETH}
	case coin.Classic().Handle:
		return &ethereum.Platform{CoinIndex: coin.ETC}
	case coin.Poa().Handle:
		return &ethereum.Platform{CoinIndex: coin.POA}
	case coin.Callisto().Handle:
		return &ethereum.Platform{CoinIndex: coin.CLO}
	case coin.Gochain().Handle:
		return &ethereum.Platform{CoinIndex: coin.GO}
	case coin.Wanchain().Handle:
		return &ethereum.Platform{CoinIndex: coin.WAN}
	case coin.Tomochain().Handle:
		return &ethereum.Platform{CoinIndex: coin.TOMO}
	case coin.Thundertoken().Handle:
		return &ethereum.Platform{CoinIndex: coin.TT}
	case coin.Cosmos().Handle:
		return &cosmos.Platform{CoinIndex: coin.ATOM}
	case coin.Kava().Handle:
		return &cosmos.Platform{CoinIndex: coin.KAVA}
	case coin.Tezos().Handle:
		return &tezos.Platform{}
	case coin.Aion().Handle:
		return &aion.Platform{}
	case coin.Icon().Handle:
		return &icon.Platform{}
	case coin.Iotex().Handle:
		return &iotex.Platform{}
	case coin.Ontology().Handle:
		return &ontology.Platform{}
	case coin.Theta().Handle:
		return &theta.Platform{}
	case coin.Tron().Handle:
		return &tron.Platform{}
	case coin.Vechain().Handle:
		return &vechain.Platform{}
	case coin.Zilliqa().Handle:
		return &zilliqa.Platform{}
	case coin.Waves().Handle:
		return &waves.Platform{}
	case coin.Aeternity().Handle:
		return &aeternity.Platform{}
	case coin.Bitcoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.BTC}
	case coin.Litecoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.LTC}
	case coin.Bitcoincash().Handle:
		return &bitcoin.Platform{CoinIndex: coin.BCH}
	case coin.Dash().Handle:
		return &bitcoin.Platform{CoinIndex: coin.DASH}
	case coin.Doge().Handle:
		return &bitcoin.Platform{CoinIndex: coin.DOGE}
	case coin.Zcash().Handle:
		return &bitcoin.Platform{CoinIndex: coin.ZEC}
	case coin.Zcoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.XZC}
	case coin.Viacoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.VIA}
	case coin.Ravencoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.RVN}
	case coin.Qtum().Handle:
		return &bitcoin.Platform{CoinIndex: coin.QTUM}
	case coin.Groestlcoin().Handle:
		return &bitcoin.Platform{CoinIndex: coin.GRS}
	case coin.Zelcash().Handle:
		return &bitcoin.Platform{CoinIndex: coin.ZEL}
	case coin.Decred().Handle:
		return &bitcoin.Platform{CoinIndex: coin.DCR}
	case coin.Digibyte().Handle:
		return &bitcoin.Platform{CoinIndex: coin.DGB}
	case coin.Nebulas().Handle:
		return &nebulas.Platform{}
	case coin.Fio().Handle:
		return &fio.Platform{}
	case coin.Algorand().Handle:
		return &algorand.Platform{}
	case coin.Nano().Handle:
		return &nano.Platform{}
	case coin.Harmony().Handle:
		return &harmony.Platform{}
	case coin.Kusama().Handle:
		return &polkadot.Platform{CoinIndex: coin.KSM}
	default:
		return nil
	}
}

func getActivePlatforms(handle string) []blockatlas.Platform {
	logger.Info("Loaded with: ", logger.Params{"handle": handle})

	platform := getPlatform(handle)
	if platform != nil {
		return []blockatlas.Platform{platform}
	}
	return allPlatformsList
}

func Init(platformHandle string) {
	platformList := getActivePlatforms(platformHandle)

	Platforms = make(map[string]blockatlas.Platform)
	BlockAPIs = make(map[string]blockatlas.BlockAPI)
	StakeAPIs = make(map[string]blockatlas.StakeAPI)
	CustomAPIs = make(map[string]blockatlas.CustomAPI)
	NamingAPIs = make(map[uint64]blockatlas.NamingServiceAPI)
	CollectionAPIs = make(map[uint]blockatlas.CollectionAPI)

	for _, platform := range platformList {
		handle := platform.Coin().Handle
		apiURL := fmt.Sprintf("%s.api", handle)

		if !viper.IsSet(apiURL) {
			continue
		}
		if viper.GetString(apiURL) == "" {
			continue
		}

		p := logger.Params{
			"platform": handle,
			"coin":     platform.Coin(),
		}

		if _, exists := Platforms[handle]; exists {
			logger.Fatal("Duplicate handle", p)
		}

		err := platform.Init()
		if err != nil {
			logger.Error("Failed to initialize API", err, p)
		}

		Platforms[handle] = platform

		if blockAPI, ok := platform.(blockatlas.BlockAPI); ok {
			BlockAPIs[handle] = blockAPI
		}
		if stakeAPI, ok := platform.(blockatlas.StakeAPI); ok {
			StakeAPIs[handle] = stakeAPI
		}
		if customAPI, ok := platform.(blockatlas.CustomAPI); ok {
			CustomAPIs[handle] = customAPI
		}
		if namingAPI, ok := platform.(blockatlas.NamingServiceAPI); ok {
			NamingAPIs[uint64(platform.Coin().ID)] = namingAPI
		}
		if collectionAPI, ok := platform.(blockatlas.CollectionAPI); ok {
			CollectionAPIs[platform.Coin().ID] = collectionAPI
		}
	}
}
