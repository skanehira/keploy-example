package main

import (
	"encoding/json"
	"io"
)

func serialize[W io.Writer, T any](w W, v T) error {
	return json.NewEncoder(w).Encode(&v)
}

func deserialize[R io.Reader, T any](r R, v T) error {
	return json.NewDecoder(r).Decode(&v)
}
