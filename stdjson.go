//go:build !jsoniter && !go_json && !(sonic && avx && (linux || windows || darwin) && amd64)

package zson

import "encoding/json"

var (
	// Marshal returns the JSON encoding of v.
	Marshal = json.Marshal
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	Unmarshal = json.Unmarshal
	// MarshalIndent is like Marshal but applies Indent to format the output.
	MarshalIndent = json.MarshalIndent
	// NewDecoder returns a new decoder that reads from r.
	NewDecoder = json.NewDecoder
	// NewEncoder returns a new encoder that writes to w.
	NewEncoder = json.NewEncoder
)
