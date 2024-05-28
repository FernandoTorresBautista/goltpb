package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OutputResponse response of the get api
type OutputResponse struct {
	Ltp []struct {
		Pair   string `json:"pair"`
		Amount string `json:"amount"`
	} `json:"ltp"`
}

// apiURL example: https://api.kraken.com/0/public/Ticker?pair=XBTUSD
var apiURL string = "https://api.kraken.com/0/public/Ticker?pair=%s"

// GetInfo ...
// @Summary GetInfo
// @Description GetInfo return the Last Traded Price of Bitcoin for the following currency pairs: BTC/USD, BTC/CHF, BTC/EUR
// @Tags LTPB
// @Produce json
// @Success 200 {object} OutputResponse "return the list of pairs"
// @Success 400 {object} map[string]string "return the error of the bad request"
// @Success 500 {object} map[string]string "return the error of the failure in the API"
// @Router /api/v1/ltp [get]
func (api *Apiv1) GetInfo(c *gin.Context) {
	// works with both
	currencyPairs := []string{"BTCCHF", "BTCEUR", "BTCUSD"} // []string{"XBTCHF", "XBTEUR", "XBTUSD"}
	responses := [3]interface{}{}
	for i, v := range currencyPairs {
		// build url
		iURL := fmt.Sprintf(apiURL, v)
		// fmt.Println(iURL)
		// Get Ticker Information
		resp, err := http.Get(iURL)
		if err != nil {
			api.logger.Printf("error getting response from the api: %+v\n", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			api.logger.Printf("error getting response from the api, code: %d\n", resp.StatusCode)
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("error getting response from the api, code: %d", resp.StatusCode))
			return
		}
		//
		err = json.NewDecoder(resp.Body).Decode(&responses[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	// tranform the information
	currencyPairsResults := []string{"XBTCHF", "XXBTZEUR", "XXBTZUSD"}
	currencyPairsNames := []string{"BTC/CHF", "BTC/EUR", "BTC/USD"}
	resp := OutputResponse{}
	for i, v := range responses {
		// fmt.Printf("%d: %+v\n", i, v)
		// convert to a map
		responseMap, ok := v.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, "error accesing to the result")
			return
		}
		// access to result
		result, ok := responseMap["result"].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, "error accesing to the map result")
			return
		}
		// fmt.Printf("result: %+v\n", result)
		// access to the submap
		list, ok := result[currencyPairsResults[i]].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, "error accesing to the currency pair result")
			return
		}
		// fmt.Printf("result list: %+v\n", list)
		// access to the c array
		ltc, ok := list["c"].([]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, "error accesing to the Last trade closed array")
			return
		}
		// fmt.Printf("ltc: %s\n", ltc[0])
		resp.Ltp = append(resp.Ltp, struct {
			Pair   string "json:\"pair\""
			Amount string "json:\"amount\""
		}{Pair: currencyPairsNames[i], Amount: ltc[0].(string)})
	}
	// return the information
	c.JSON(http.StatusOK, resp)
}
