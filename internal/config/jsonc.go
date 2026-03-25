package config

import "github.com/tidwall/jsonc"

// StripComments removes JSONC comments and trailing commas from input,
// producing valid JSON. Returns the input unchanged if it contains no
// comments or trailing commas.
func StripComments(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	return jsonc.ToJSON(data)
}
