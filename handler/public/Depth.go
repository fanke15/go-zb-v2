package public

import (
	"encoding/json"
	"fmt"
	. "github.com/fanke15/go-zb-v2/log"
	"github.com/fanke15/go-zb-v2/types"
)

func DepthHandler(msg []byte) bool {
	var data types.Depth
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}

	Log.Info().Str("asks", fmt.Sprint(data.Asks)).Msg("Depth")
	Log.Info().Str("bids", fmt.Sprint(data.Bids)).Msg("Depth")
	return false
}
