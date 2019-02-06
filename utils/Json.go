package utils

import (
	"encoding/json"
	"github.com/landonp1203/goUtils/loggly"
)

// TODO: Look into using json.NewDecoder because the response is a data stream form a server.

// Decodes JSON data.
func DecodeJson(data []byte, v interface{}) {
	if data != nil {
		var err = json.Unmarshal(data, v)

		if err != nil {
			loggly.Error("There was an error parsing the JSON data.")
		}
	} else {
		loggly.Warn("The JSON data was nil, so no data was parsed.")
	}
}
