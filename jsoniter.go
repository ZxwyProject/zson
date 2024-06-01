//go:build jsoniter

package zson

import jsoniter "github.com/json-iterator/go"

var (
	Json = jsoniter.Config{
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
	// Marshal adapts to json/encoding Marshal API
	Marshal = Json.Marshal
	// Unmarshal adapts to json/encoding Unmarshal API
	Unmarshal = Json.Unmarshal
	// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
	MarshalIndent = Json.MarshalIndent
	// NewDecoder adapts to json/stream NewDecoder API.
	NewDecoder = Json.NewDecoder
	// NewEncoder same as json.NewEncoder
	NewEncoder = Json.NewEncoder
)
