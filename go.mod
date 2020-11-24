module github.com/trustwallet/blockatlas

go 1.15

// +heroku goVersion go1.15
// +heroku install ./cmd/...

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/btcsuite/btcutil v1.0.2
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/chenjiandongx/ginprom v0.0.0-20200410120253-7cfb22707fa6
	github.com/containerd/continuity v0.0.0-20200413184840-d3ef23f19fbb // indirect
	github.com/deckarep/golang-set v1.7.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/elastic/go-sysinfo v1.3.0 // indirect
	github.com/elastic/go-windows v1.0.1 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/imroc/req v0.3.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/mr-tron/base58 v1.2.0
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/ory/dockertest v3.3.5+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/trustwallet/golibs v0.0.15
	go.elastic.co/apm v1.8.0
	go.elastic.co/apm/module/apmgin v1.8.0
	go.elastic.co/apm/module/apmhttp v1.8.0
	go.elastic.co/fastjson v1.1.0 // indirect
	go.uber.org/atomic v1.7.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/tools v0.0.0-20200513175351-0951661448da // indirect
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/postgres v1.0.0
	gorm.io/gorm v1.20.0
	gorm.io/plugin/dbresolver v1.0.0
	gotest.tools v2.2.0+incompatible // indirect
	howett.net/plist v0.0.0-20200419221736-3b63eb3a43b5 // indirect
)
