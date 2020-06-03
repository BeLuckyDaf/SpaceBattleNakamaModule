package serialization

import (
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

// Serialize is a json wrapper
func Serialize(v interface{}, logger runtime.Logger) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		if logger != nil {
			logger.Error("Could not marshal %v to json string.", v)
		}
		return nil
	}
	return b
}

// Deserialize is a json wrapper
func Deserialize(b []byte, v interface{}, logger runtime.Logger) bool {
	err := json.Unmarshal(b, v)
	if err != nil {
		if logger != nil {
			logger.Error("Could not unmarshal %s to and object.", string(b))
		}
		return false
	}
	return true
}
