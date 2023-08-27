package common

import "encoding/json"

// BinaryMarshalerFunc binary marshal all data types
type BinaryMarshalerFunc func() ([]byte, error)

// MarshalBinary executes the function as long as it complies with the BinaryMarshalerFunc signature
func (f BinaryMarshalerFunc) MarshalBinary() ([]byte, error) {
	return f()
}

// NewMarshalJSON marshal to a binary json the data
func NewMarshalJSON(v any) BinaryMarshalerFunc {
	return func() ([]byte, error) {
		return json.Marshal(v)
	}
}
