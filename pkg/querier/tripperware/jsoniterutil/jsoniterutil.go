package jsoniterutil

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	Shared = jsoniter.Config{
		EscapeHTML:             false, // No HTML in our responses.
		SortMapKeys:            true,
		ValidateJsonRawMessage: false,
	}.Froze()
)

func Marshal(v any) ([]byte, error) {
	return Shared.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return Shared.Unmarshal(data, v)
}
