package cmd

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/ringmaster/journey/src/backend/pkg/handlers/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run backend service",
	Long:  `run backend service`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		handlers.RegisterRoute(e)

		e.Start(":8181")
	},
}
