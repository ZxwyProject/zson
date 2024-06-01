//go:build sonic && avx && (linux || windows || darwin) && amd64

package zson

import "github.com/bytedance/sonic"

var (
	Json = sonic.ConfigDefault
	// Marshal returns the JSON encoding bytes of v.
	Marshal = Json.Marshal
	// Unmarshal parses the JSON-encoded string and stores the result in the value pointed to by v.
	Unmarshal = Json.Unmarshal
	// MarshalIndent returns the JSON encoding bytes with indent and prefix.
	MarshalIndent = Json.MarshalIndent
	// NewDecoder create a Decoder holding reader
	NewDecoder = Json.NewDecoder
	// NewEncoder create a Encoder holding writer
	NewEncoder = Json.NewEncoder
)
