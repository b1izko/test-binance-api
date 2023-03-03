package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/b1izko/test-binance-api/pkg/api/handler"
	"github.com/b1izko/test-binance-api/pkg/logger"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func start() {
	logger.Log("Starting the server...")
	router := mux.NewRouter()

	router.HandleFunc("/", handler.DefaultRoute)

	router.HandleFunc("/api/v1/rates", handler.PairRates).
		Queries("pairs", "{pairs}").
		Methods("GET")

	router.HandleFunc("/api/v1/rates", handler.PairRates).
		Methods("POST")

	http.Handle("/", router)

	logger.Log("The server is running")
	logger.Log("Listening at port 3001...")
	err := http.ListenAndServe("localhost:3001", nil)
	if err != nil {
		logger.Error(err, "The server is closed")
		log.Fatal(err)
	}
}
