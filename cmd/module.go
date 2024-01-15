package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewServerCmd), fx.Invoke(runCmd))

func runCmd(lc fx.Lifecycle, shutdowner fx.Shutdowner, server *ServerCmd) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var serverCmd = &cobra.Command{
				Use:   "start",
				Short: "Starts the server",
				Run:   server.Start(),
			}

			var rootCmd = &cobra.Command{
				Use:   "drink",
				Short: "",
			}

			config, err := loadConfig("$HOME", ".")
			if err != nil {
				fmt.Println(err)
			}

			serverCmd.Flags().StringVarP(&server.port, "port", "p", config.Port, "Port for the web server")

			rootCmd.AddCommand((serverCmd))

			if err := rootCmd.Execute(); err != nil {
				shutdowner.Shutdown(fx.ExitCode(1))
			}

			return shutdowner.Shutdown(fx.ExitCode(0))
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
