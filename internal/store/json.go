package store

import (
	"encoding/json"
	"fmt"
)

func encode(value any) ([]byte, error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("encode record: %w", err)
	}
	return payload, nil
}

func decode(payload []byte, target any) error {
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf("decode record: %w", err)
	}
	return nil
}
