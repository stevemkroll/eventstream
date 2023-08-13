package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevemkroll/eventstream/internal/service"
	"golang.org/x/exp/slog"
)

var msg string

func init() {
	producerCmd.Flags().StringVarP(&msg, "message", "m", "", "message for event streaming producer")
	if err := producerCmd.MarkFlagRequired("message"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(producerCmd)
}

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "event streaming producer",
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()
		slog.InfoCtx(ctx, "reading config")
		viper.SetConfigFile("config/consumer.env")
		if err := viper.ReadInConfig(); err != nil {
			slog.ErrorCtx(ctx, "unable to read config", err)
			panic(err)
		}
		slog.InfoCtx(ctx, "staring producer")
		service.RunConfig()
		service.RunProducer(msg)
	},
}
