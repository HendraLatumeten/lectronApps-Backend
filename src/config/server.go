package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HendraLatumeten/lectronApss-Backend/src/routers"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func corsHandler() *cors.Cors {
	t := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	return t
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = ":8080"
		if port := os.Getenv("PORT"); port != "" {
			addrs = ":" + port
		}

		corss := corsHandler()

		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      corss.Handler(mainRoute),
		}

		fmt.Println("Lectronic is running on PORT", addrs, "success")
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
