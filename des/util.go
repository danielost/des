package des

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func validateKey(key string) error {
	alphabet := "1234567890abcdef"
	if len(key) > 16 {
		return fmt.Errorf("key is too long")
	}
	for _, r := range key {
		if !strings.ContainsRune(alphabet, r) {
			return fmt.Errorf("key contains a prohibited symbol - [%s]", string(r))
		}
	}
	binary, _ := hexToBinary(key)
	for len(binary) < 64 {
		binary = "0" + binary
	}
	fmt.Printf("Key entropy is %.4f\n", calculateEntropy(binary, "01"))
	if entropy := calculateEntropy(binary, "01"); entropy < 0.9 {
		return fmt.Errorf("key is too weak, entropy is %.4f", entropy)
	}
	return nil
}

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

func stringToBinary(s string) (binary string) {
	for _, r := range s {
		binary = fmt.Sprintf("%s%.8b", binary, r)
	}
	return
}

func hexToBinary(hexString string) (string, error) {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	var binaryString strings.Builder
	for _, b := range hexBytes {
		binaryString.WriteString(fmt.Sprintf("%08b", b))
	}

	return binaryString.String(), nil
}

func calculateEntropy(s string, alphabet string) float64 {
	sum := 0.0
	for _, r := range alphabet {
		c := strings.Count(s, string(r))
		p := float64(c) / float64(len(s))
		sum += -p * math.Log2(p)
	}
	return sum
}
