package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/model"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/http"
	"strconv"
)

// @Summary Get Collection
// @ID collection_v4
// @Description Get a collection from the address
// @Accept json
// @Produce json
// @Tags Collections
// @Param coin path string true "the coin name" default(ethereum)
// @Param owner path string true "the query address" default(0x0875BCab22dE3d02402bc38aEe4104e1239374a7)
// @Param collection_id path string true "the query collection" default(0x06012c8cf97bead5deae237070f9587f8e7a266d)
// @Success 200 {object} blockatlas.CollectionPage
// @Failure 500 {object} middleware.ApiError
// @Router /v4/{coin}/collections/{owner}/collection/{collection_id} [get]
func GetCollectiblesForSpecificCollectionAndOwner(c *gin.Context, api blockatlas.CollectionAPI) {
	collectibles, err := api.GetCollectibles(c.Param("owner"), c.Param("collection_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	c.JSON(http.StatusOK, &collectibles)
}

// @Description Get collection categories
// @ID collection_categories_v4
// @Summary Get list of collections from a specific coin and addresses
// @Accept json
// @Produce json
// @Tags Collections
// @Param data body string true "Payload" default({"60": ["0xb3624367b1ab37daef42e1a3a2ced012359659b0"]})
// @Success 200 {object} blockatlas.DocsResponse
// @Router /v4/collectibles/categories [post]
func GetCollectionCategoriesFromList(c *gin.Context, apis map[uint]blockatlas.CollectionAPI) {
	var reqs map[string][]string
	if err := c.BindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.Default, err))
		return
	}

	batch := make(blockatlas.CollectionPage, 0)
	for key, addresses := range reqs {
		coinId, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		p, ok := apis[uint(coinId)]
		if !ok {
			continue
		}
		for _, address := range addresses {
			collections, err := p.GetCollections(address)
			if err != nil {
				continue
			}
			batch = append(batch, collections...)
		}
	}
	c.JSON(http.StatusOK, &batch)
}

func GetCollectiblesForOwnerV3(c *gin.Context, api blockatlas.CollectionAPI) {
	collections, err := api.GetCollectionsV3(c.Param("owner"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}

	c.JSON(http.StatusOK, &collections)
}

func GetCollectiblesForSpecificCollectionAndOwnerV3(c *gin.Context, api blockatlas.CollectionAPI) {
	collectibles, err := api.GetCollectiblesV3(c.Param("owner"), c.Param("collection_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateErrorResponse(model.InternalFail, err))
		return
	}
	c.JSON(http.StatusOK, &collectibles)
}

func GetCollectionCategoriesFromListV3(c *gin.Context, apis map[uint]blockatlas.CollectionAPI) {
	var reqs map[string][]string
	if err := c.BindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateErrorResponse(model.Default, err))
		return
	}

	batch := make(blockatlas.CollectionPageV3, 0)
	for key, addresses := range reqs {
		coinId, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		p, ok := apis[uint(coinId)]
		if !ok {
			continue
		}
		for _, address := range addresses {
			collections, err := p.GetCollectionsV3(address)
			if err != nil {
				continue
			}
			batch = append(batch, collections...)
		}
	}
	c.JSON(http.StatusOK, &batch)
}
