package definitionparser

import "github.com/pelletier/go-toml/v2"

// ParseFile is an abstraction of a TOML parser.
//
// Internally, it uses the go-toml package.
func ParseFile(data []byte, v interface{}) error {
	return toml.Unmarshal(data, &v)
}
