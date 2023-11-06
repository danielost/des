package des

import (
	"strings"
)

func pc1(key []string) ([]string, []string) {
	return []string{
			key[56], key[48], key[40], key[32], key[24], key[16], key[8],
			key[0], key[57], key[49], key[41], key[33], key[25], key[17],
			key[9], key[1], key[58], key[50], key[42], key[34], key[26],
			key[18], key[10], key[2], key[59], key[51], key[43], key[35],
		},
		[]string{
			key[62], key[54], key[46], key[38], key[30], key[22], key[14],
			key[6], key[61], key[53], key[45], key[37], key[29], key[21],
			key[13], key[5], key[60], key[52], key[44], key[36], key[28],
			key[20], key[12], key[4], key[27], key[19], key[11], key[3],
		}
}

func pc2(key []string) []string {
	return []string{
		key[13], key[16], key[10], key[23], key[0], key[4], key[2], key[27],
		key[14], key[5], key[20], key[9], key[22], key[18], key[11], key[3],
		key[25], key[7], key[15], key[6], key[26], key[19], key[12], key[1],
		key[40], key[51], key[30], key[36], key[46], key[54], key[29], key[39],
		key[50], key[44], key[32], key[47], key[43], key[48], key[38], key[55],
		key[33], key[52], key[45], key[41], key[49], key[35], key[28], key[31],
	}
}

func leftShift(key []string, shiftBy int) []string {
	shiftedKey := key[shiftBy:]
	return append(shiftedKey, key[:shiftBy]...)
}

func expandKey(key string) []string {
	binaryKey, _ := hexToBinary(key) //stringToBinary(key)
	for len(binaryKey) < 64 {
		binaryKey = "0" + binaryKey
	}
	left, right := pc1(strings.Split(binaryKey, ""))

	leftShiftIndex := []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}
	keys := make([]string, 16)

	for i := 0; i < 16; i++ {
		left = leftShift(left, leftShiftIndex[i])
		right = leftShift(right, leftShiftIndex[i])
		key := append(left, right...)
		keys[i] = strings.Join(pc2(key), "")
	}

	return keys
}
