package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/CrowderSoup/drinkingaroundthe.world/database"
	"github.com/CrowderSoup/drinkingaroundthe.world/web"
)

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("port", "p", viper.GetString("port"), "The port that the web server should run on")
	serverCmd.Flags().StringP("db_connection_string", "d", viper.GetString("db_connection_string"), "The connection string for our postgres db")
	serverCmd.Flags().StringP("secret", "s", viper.GetString("secret"), "Some secret key used for sessions etc.")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the web server",
	Long:  `Starts the web server for running Drink Around the World`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDatabase(cmd.Flag("db_connection_string").Value.String())
		e := echo.New()

		server := web.NewServer(e, db, cmd.Flag("secret").Value.String())

		server.Start(cmd.Flag("port").Value.String())
	},
}
