package private

import (
	"encoding/json"
	"fmt"
	"github.com/fanke15/go-zb-v2/types"

	. "github.com/fanke15/go-zb-v2/log"
)

func PositionsChangeHandler(msg []byte) bool {
	var data types.PositionsChangeResult
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("PositionsChangeHandler")
	return false
}
