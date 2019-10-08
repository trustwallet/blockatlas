package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/redis"
)

var testCmd = cobra.Command{
	Use:   "test",
	Short: "tests",
	Run:   runTest,
}

func runTest(_ *cobra.Command, args []string) {
	db := &redis.Redis{}
	err := db.Init(viper.GetString("observer.redis"))
	if err != nil {
		logger.Error(err)
	}
	db.Entity("test").Query("keybla")
	err = db.Set([]string{"caralho da porra", "bla"})
	if err != nil {
		logger.Error(err)
	}

	var test []string
	err = db.Into(&test)
	if err != nil {
		logger.Error(err)
	}

}

func init() {
	rootCmd.AddCommand(&testCmd)
}
