package hamming

import "errors"

const (
	//	packetSize = 64
	//	parityBit  = 1 //(corrensponds to the 0b000000 index)
	//	//parityGroupBits =
	inputMask = (1 << (64 - 7)) - 1
)

//Encoder offer the Hamming encoding functionality for 64 bits packets
//type Encoder struct {
//}

// func encode(input uint64) (uint64,error) {
// 	if input != input&inputMask {
// 		return
// 	}

// }

var parityMask = map[int]uint64{
	0:  0b1111111111111111111111111111111111111111111111111111111111111111,
	1:  0b1010101010101010101010101010101010101010101010101010101010101010,
	2:  0b1100110011001100110011001100110011001100110011001100110011001100,
	4:  0b1111000011110000111100001111000011110000111100001111000011110000,
	8:  0b1111111100000000111111110000000011111111000000001111111100000000,
	16: 0b1111111111111111000000000000000011111111111111110000000000000000,
	32: 0b1111111111111111111111111111111100000000000000000000000000000000,
}

var errInputTooBig = errors.New("The input number is bigger than the maximum size that can be processed")

func encode(input uint64) (uint64, error) {
	out, err := reallocateBits(input)
	if err != nil {
		return 0, err
	}
	for i := 1; i < 64; i <<= 1 {
		out = addParityBit(out, i)
	}
	out = addParityBit(out, 0)
	return out, nil
}

func addParityBit(input uint64, bit int) uint64 {

	process := input & parityMask[bit]

	for i := 1; i < 64; i <<= 1 {
		process ^= process >> i
	}
	process &= 1
	if process == 1 {
		input |= 1 << bit
	}
	return input
}

func reallocateBits(input uint64) (uint64, error) {
	if input != input&inputMask {
		return 0, errInputTooBig
	}

	output := partialDisplacement(input, 0)

	for i := 1; i < 64; i <<= 1 {
		output = partialDisplacement(output, i)
	}

	return output, nil
}

func partialDisplacement(input uint64, position int) uint64 {
	//Example position 3
	// input  0b0000000000000000000000000000000000000000000000000000000001101101
	//                                                                      ^
	//                                                                      |
	// result 0b0000000000000000000000000000000000000000000000000000000011010101

	right := (input >> position) << (position + 1)
	left := (input << (64 - position)) >> (64 - position)

	return right | left
}

// 0000 | 000000 Parity bit
// with just one 1 Parity group
