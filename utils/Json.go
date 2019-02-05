package utils

import (
	"encoding/json"
)

// TODO: Look into using json.NewDecoder because the response is a data stream form a server.

// Decodes JSON data.
func DecodeJson(data []byte, v interface{}) {
	if data != nil {
		var err = json.Unmarshal(data, v)

		if err != nil {
			Error("There was an error parsing the JSON data.")
		}
	} else {
		Warn("The JSON data was nil, so no data was parsed.")
	}
}
