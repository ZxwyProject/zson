//go:build go_json

package zson

import json "github.com/goccy/go-json"

var (
	// Marshal returns the JSON encoding of v.
	Marshal = json.MarshalNoEscape
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	Unmarshal = json.UnmarshalNoEscape
	// MarshalIndent is like Marshal but applies Indent to format the output.
	MarshalIndent = json.MarshalIndent
	// NewDecoder returns a new decoder that reads from r.
	NewDecoder = json.NewDecoder
	// NewEncoder returns a new encoder that writes to w.
	NewEncoder = json.NewEncoder
)
