package types

import (
	"encoding/json"
	"fmt"
)

type String string

func (s *String) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = String(fmt.Sprintf("%v", v))
	return nil
}

func (s String) Get() string {
	return string(s)
}
