package json

import "encoding/json"

func ToStruct[T any](value []byte) (*T, error) {
	var resp T
	if err := json.Unmarshal(value, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
