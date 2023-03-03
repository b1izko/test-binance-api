package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/b1izko/test-binance-api/pkg/api"
	"github.com/b1izko/test-binance-api/pkg/api/errors"
	"github.com/b1izko/test-binance-api/pkg/logger"
	"github.com/b1izko/test-binance-api/pkg/utils"
	"github.com/gorilla/mux"
)

// DefaultRoute is default route
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	logger.Request(r.RequestURI, r.RemoteAddr)
	api.WriteError(w, errors.CodeInvalidRequest, "wrong request")
	return
}

// PairRates is route for pairs handler
func PairRates(w http.ResponseWriter, r *http.Request) {
	logger.Request(r.RequestURI, r.RemoteAddr)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.WriteError(w, errors.CodeInvalidRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var buffer struct {
		Pairs []string `json:"pairs"`
	}

	if len(body) != 0 {
		err = json.Unmarshal(body, &buffer)
		if err != nil {
			api.WriteError(w, errors.CodeParseError, err.Error())
			return
		}
	}

	var pairs []string
	if len(buffer.Pairs) == 0 {
		values := mux.Vars(r)
		pairs = strings.Split(values["pairs"], ",")
	} else {
		pairs = buffer.Pairs
	}

	result, err := utils.SendRequestToBinance(pairs)
	if err != nil {
		api.WriteError(w, errors.CodeSendRequestError, err.Error())
		return
	}

	response := api.NewResponse(result)
	response.WriteResponse(w)
	return
}
