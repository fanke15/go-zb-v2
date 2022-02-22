package private

import (
	"encoding/json"
	"fmt"
	"github.com/fanke15/go-zb-v2/types"

	. "github.com/fanke15/go-zb-v2/log"
)

func GetUndoneOrderHandler(msg []byte) bool {
	var data types.GetAllOrderResult
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("GetUndoneOrderHandler")
	return false
}
