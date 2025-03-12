package cmd

import (
	"github.com/spf13/cobra"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/routes"
)

func init() {
	rootCmd.AddCommand(serveBackOfficeAPICmd)
}

var serveBackOfficeAPICmd = &cobra.Command{
	Use:   "serve-http-api",
	Short: "Start HTTP API server",
	RunE: func(cmd *cobra.Command, args []string) error {

		logger, err := getLogger()
		if err != nil {
			return err
		}

		app, err := app.New(logger)
		if err != nil {
			return err
		}

		config, err := routes.InitConfig()
		if err != nil {
			return err
		}

		routes.NewRouter(config, logger, app)

		return nil
	},
}
