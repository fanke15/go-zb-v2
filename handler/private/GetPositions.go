package private

import (
	"encoding/json"
	"fmt"
	"github.com/fanke15/go-zb-v2/types"

	. "github.com/fanke15/go-zb-v2/log"
)

func GetPositionsHandler(msg []byte) bool {
	var data []types.GetPositionsResult
	err := json.Unmarshal(msg, &data)
	if err != nil {
		Log.Error().Str("data", string(msg)).Msg(err.Error())
		return true
	}
	Log.Info().Str("data", fmt.Sprint(data)).Msg("GetPositionsHandler")
	return false
}
