package public

import (
	"encoding/json"
	"fmt"
	. "github.com/fanke15/go-zb-v2/log"
	"github.com/fanke15/go-zb-v2/types"
)

func AllTickerHandler(msg []byte) bool {
	var data types.AllTicker
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("AllTicker")
	return false
}
