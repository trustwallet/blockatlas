package platform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

var loaders = map[string]func(gin.IRouter){}

func Add(name string, loader func(gin.IRouter)) {
	loaders[name] = loader
}

func Register(root gin.IRouter) {
	for _, ns := range viper.GetStringSlice("platforms") {
		loader := loaders[ns]
		if loader == nil {
			fmt.Fprintf(os.Stderr, "Failed to load platform %s\n", ns)
			os.Exit(1)
		}
		loader(root.Group(ns))
		fmt.Printf("Loaded /%s\n", ns)
	}
}
