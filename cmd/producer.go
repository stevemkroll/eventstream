package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/stevemkroll/eventstream/internal/service"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var msg string

func init() {
	producerCmd.Flags().StringVarP(&msg, "message", "m", "", "message for event streaming producer")
	producerCmd.MarkFlagRequired("message")
	rootCmd.AddCommand(producerCmd)
}

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "event streaming producer",
	Run: func(_ *cobra.Command, _ []string) {
		log.Println("| reading config")
		viper.SetConfigFile("config/consumer.env")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}

		log.Println("| running producer")
		service.RunConfig()
		service.RunProducer(msg)
	},
}
