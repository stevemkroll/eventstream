package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevemkroll/eventstream/internal/service"
	"golang.org/x/exp/slog"
)

var run bool

func init() {
	consumerCmd.Flags().BoolVarP(&run, "run", "r", true, "run the consumer")
	if err := consumerCmd.MarkFlagRequired("run"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(consumerCmd)
}

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "event streaming consumer",
	Run: func(_ *cobra.Command, _ []string) {
		ctx := context.Background()
		slog.InfoCtx(ctx, "reading config")
		viper.SetConfigFile("config/consumer.env")
		if err := viper.ReadInConfig(); err != nil {
			slog.ErrorCtx(ctx, "unable to read config", err)
			panic(err)
		}
		slog.InfoCtx(ctx, "staring consumer")
		service.RunConfig()
		service.RunConsumer(ctx)
	},
}
