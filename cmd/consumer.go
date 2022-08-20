package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/stevemkroll/eventstream/internal/service"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var run bool

func init() {
	consumerCmd.Flags().BoolVarP(&run, "run", "r", true, "run the consumer")
	consumerCmd.MarkFlagRequired("run")
	rootCmd.AddCommand(consumerCmd)
}

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "event streaming consumer",
	Run: func(_ *cobra.Command, _ []string) {
		log.Println("| reading config")
		viper.SetConfigFile("config/consumer.env")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}

		log.Println("| running consumer")
		service.RunConfig()
		service.RunConsumer()
	},
}
