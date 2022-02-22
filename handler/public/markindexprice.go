package public

import (
	"encoding/json"
	. "github.com/fanke15/go-zb-v2/log"
	"strconv"
)

func MarkIndexPriceHandler(msg []byte) bool {
	var data map[string]string
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	if data == nil {
		return false
	}
	for k, v := range data {
		price, err := strconv.ParseFloat(v, 10)
		if err != nil {
			Log.Error().Str("data", string(msg)).Msg(err.Error())
			return true
		}
		Log.Info().Str("symbol", k).Float64("price", price).Msg("MarkIndexPriceHandler")
	}

	return false
}
