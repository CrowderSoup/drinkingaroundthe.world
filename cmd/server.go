package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/CrowderSoup/drinkingaroundthe.world/web"
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("port", "p", viper.GetString("port"), "The port that the web server should run on")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the web server",
	Long:  `Starts the web server for running Drink Around the World`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		server := web.NewServer(e)

		server.Start(cmd.Flag("port").Value.String())
	},
}
