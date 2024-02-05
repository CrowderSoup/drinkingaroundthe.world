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
	serverCmd.Flags().String("mailgun_api_key", viper.GetString("mailgun_api_key"), "API Key for Mailgun")
	serverCmd.Flags().String("mailgun_domain", viper.GetString("mailgun_domain"), "Domain for Mailgun")
	serverCmd.Flags().String("mailgun_sending_address", viper.GetString("mailgun_sending_address"), "Sending Email address for Mailgun")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the web server",
	Long:  `Starts the web server for running Drink Around the World`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDatabase(cmd.Flag("db_connection_string").Value.String())
		e := echo.New()

		// Wire up flags
		secret := cmd.Flag("secret").Value.String()

		server := web.NewServer(e, db, secret)

		server.Start(cmd.Flag("port").Value.String())
	},
}
