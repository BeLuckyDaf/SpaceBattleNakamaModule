package sjson

import (
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

// Marshal is a json wrapper
func Marshal(v interface{}, logger runtime.Logger) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		if logger != nil {
			logger.Error("Could not marshal %v to json string.", v)
		}
		return nil
	}
	return b
}

// Unmarshal is a json wrapper
func Unmarshal(b []byte, v interface{}, logger runtime.Logger) bool {
	err := json.Unmarshal(b, v)
	if err != nil {
		if logger != nil {
			logger.Error("Could not unmarshal %s to and object.", string(b))
		}
		return false
	}
	return true
}
