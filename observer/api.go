package observer

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/util"
)

func Setup(router gin.IRouter) {
	router.Use(util.RequireConfig("storage"))

	router.GET("/", listObservers)
	router.POST("/", addObserver)
	router.DELETE("/", removeObserver)
}

func listObservers(c *gin.Context) {
	s := storage.GetInstance()
	if list := s.List(); list != nil {
		c.JSON(200, list)
		return
	}

	c.JSON(200, make([]string, 0))
}

func addObserver(c *gin.Context) {
	var r AddRequest
	if c.Bind(&r) != nil {
		return
	}

	s := storage.GetInstance()
	observer := s.Add(r.Coin, r.Address, r.Webhook)
	c.JSON(200, observer)
}

func removeObserver(c *gin.Context) {
	var r RemoveRequest
	if c.BindQuery(&r) != nil {
		return
	}

	s := storage.GetInstance()
	if !s.Contains(r.Coin, r.Address) {
		c.JSON(404, gin.H{"message": "observer not found"})
		return
	}

	s.Remove(r.Coin, r.Address)
	c.JSON(200, gin.H{})
}

type AddRequest struct {
	Coin    uint   `form:"coin" binding:"required"`
	Address string `form:"address" binding:"required"`
	Webhook string `form:"webhook" binding:"required"`
}

type RemoveRequest struct {
	Coin    uint   `form:"coin" binding:"required"`
	Address string `form:"address" binding:"required"`
}
