package des

import "fmt"

// Validates hex key.
// TODO: check key strength.
func validateKey(key string) error {
	if len(key) > 16 {
		return fmt.Errorf("key is too long")
	}
	return nil
}
