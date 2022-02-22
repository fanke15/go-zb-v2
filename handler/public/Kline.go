package public

import (
	"encoding/json"
	"fmt"
	. "github.com/fanke15/go-zb-v2/log"
	"github.com/fanke15/go-zb-v2/types"
)

func KLineHandler(msg []byte) bool {
	var data types.KLine
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("KLine")
	return false
}
