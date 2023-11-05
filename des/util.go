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

// Breaks a message into blocks of size N.
func blocksFromMessage(message string, blockSize int) []string {
	if blockSize >= len(message) {
		return []string{message}
	}

	blocks := make([]string, 0)
	for i := 0; i < len(message); i += blockSize {
		block := message[i:min(i+blockSize, len(message))]
		blocks = append(blocks, block)
	}
	return blocks
}
