package des

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	bignumbers "github.com/danielost/big-numbers/src"
)

const numberOfRounds int = 16

func initialPermutation(b []string) []string {
	return []string{
		b[57], b[49], b[41], b[33], b[25], b[17], b[9], b[1],
		b[59], b[51], b[43], b[35], b[27], b[19], b[11], b[3],
		b[61], b[53], b[45], b[37], b[29], b[21], b[13], b[5],
		b[63], b[55], b[47], b[39], b[31], b[23], b[15], b[7],
		b[56], b[48], b[40], b[32], b[24], b[16], b[8], b[0],
		b[58], b[50], b[42], b[34], b[26], b[18], b[10], b[2],
		b[60], b[52], b[44], b[36], b[28], b[20], b[12], b[4],
		b[62], b[54], b[46], b[38], b[30], b[22], b[14], b[6],
	}
}

func reversedInitialPermutation(s []string) []string {
	return []string{
		s[39], s[7], s[47], s[15], s[55], s[23], s[63], s[31],
		s[38], s[6], s[46], s[14], s[54], s[22], s[62], s[30],
		s[37], s[5], s[45], s[13], s[53], s[21], s[61], s[29],
		s[36], s[4], s[44], s[12], s[52], s[20], s[60], s[28],
		s[35], s[3], s[43], s[11], s[51], s[19], s[59], s[27],
		s[34], s[2], s[42], s[10], s[50], s[18], s[58], s[26],
		s[33], s[1], s[41], s[9], s[49], s[17], s[57], s[25],
		s[32], s[0], s[40], s[8], s[48], s[16], s[56], s[24],
	}
}

func Encrypt(message, key string) (string, error) {
	if err := validateKey(key); err != nil {
		return "", err
	}

	blocks := blocksFromMessage(message, 8)
	lastBlock := blocks[len(blocks)-1]
	if len(lastBlock) < 8 {
		lastBlock += "8"
		for len(lastBlock) < 8 {
			lastBlock += "0"
		}
		blocks[len(blocks)-1] = lastBlock
	} else {
		blocks = append(blocks, "80000000")
	}

	keys := expandKey(key)
	encodedMessage := ""

	for i, block := range blocks {
		fmt.Printf("Block %d:\n", i+1)
		binary := strings.Split(stringToBinary(block), "")
		inititalPermutation := initialPermutation(binary)
		encodedBlockLeft, encodedBlockRight := feistel(inititalPermutation, keys, true)
		reversedInititalPermutation := reversedInitialPermutation(append(encodedBlockRight, encodedBlockLeft...))
		encodedMessage += strings.Join(reversedInititalPermutation, "")
	}
	var bn bignumbers.BigNumber
	bn.SetBinary(encodedMessage)
	return bn.GetHex(), nil
}

func Decrypt(message, key string) (string, error) {
	if err := validateKey(key); err != nil {
		return "", err
	}
	blocks := blocksFromMessage(message, 16)
	keys := expandKey(key)
	encodedMessage := ""

	for i, block := range blocks {
		fmt.Printf("Block %d:\n", i+1)
		i, _ := strconv.ParseUint(block, 16, 64)
		binary := fmt.Sprintf("%064b", i)
		inititalPermutation := initialPermutation(strings.Split(binary, ""))
		encodedBlockLeft, encodedBlockRight := feistel(inititalPermutation, keys, false)
		reversedInititalPermutation := reversedInitialPermutation(append(encodedBlockRight, encodedBlockLeft...))
		encodedMessage += strings.Join(reversedInititalPermutation, "")
	}

	var bn bignumbers.BigNumber
	bn.SetBinary(encodedMessage)
	bs, _ := hex.DecodeString(bn.GetHex())
	asciiMessage := string(bs)
	paddingIndex := strings.LastIndex(asciiMessage, "8")
	if paddingIndex != -1 {
		return asciiMessage[:paddingIndex], nil
	}
	return asciiMessage, nil
}
