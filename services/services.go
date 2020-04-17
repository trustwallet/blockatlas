package services

import (
	"github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/blockatlas/services/domains"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
)

func Init(serviceRepo *servicerepo.ServiceRepo) {
	assets.InitService(serviceRepo)
	domains.InitService(serviceRepo)
}
