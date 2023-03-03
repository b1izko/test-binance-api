package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/b1izko/test-binance-api/pkg/logger"
	"github.com/b1izko/test-binance-api/pkg/utils"
)

var isPair bool

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Return the current value of the pair(s)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pairs := strings.Split(cmd.Flags().Arg(0), ",")
		resp, err := utils.SendRequestToBinance(pairs)
		if err != nil {
			logger.Error(err, "Wrong request to Binance")
			os.Exit(1)
		}

		var result string
		for _, pair := range pairs {
			tokens := strings.Split(pair, "-")
			if len(tokens) != 2 {
				logger.Error(errors.New("pair parsing error"), "Wrong responce from Binance")
				os.Exit(1)
			}
			key := tokens[0] + tokens[1]
			result = result + tokens[0] + tokens[1] + ": " + resp[key] + "; "
		}
		fmt.Println(result[:len(result)-2])
	},
}

func init() {
	rateCmd.Flags().BoolVarP(&isPair, "pair", "p", false, "return the current value of the pair(s)")
	rootCmd.AddCommand(rateCmd)
}
