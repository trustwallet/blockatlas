module github.com/trustwallet/blockatlas

go 1.15

// +heroku goVersion go1.15
// +heroku install ./cmd/...

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/btcsuite/btcutil v1.0.2
	github.com/deckarep/golang-set v1.7.1
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/itchyny/timefmt-go v0.1.1
	github.com/mitchellh/mapstructure v1.4.1
	github.com/mr-tron/base58 v1.2.0
	github.com/opencontainers/runc v1.0.0-rc9 // indirect
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v0.9.4
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/trustwallet/golibs v0.1.1
	github.com/trustwallet/golibs/network v0.0.0-20210124080535-8638b407c4ab
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210123231150-1d476976d117 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
