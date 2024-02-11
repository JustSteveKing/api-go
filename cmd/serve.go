package cmd

import (
	"fmt"

	"github.com/juststeveking/api/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer()

		err := server.ListenAndServe()
		if err != nil {
			panic(fmt.Sprintf("cannot start server: %s", err))
		}

		fmt.Println("Running on port :8080")
	},
}
