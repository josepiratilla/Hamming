package hamming

import "errors"

var errTwoOrMoreErrorsDetected = errors.New("Detected more than one error")

func decode(input uint64) (uint64, error) {

	clean, err := solveErrors(input)
	if err != nil {
		return 0, err
	}
	out := removeAllBits(clean)
	return out, nil
}

func solveErrors(input uint64) (uint64, error) {
	parityErrors := computeParities(input)
	if len(parityErrors) == 0 {
		return input, nil
	}
	if parityErrors[0] == 0 {
		return 0, errTwoOrMoreErrorsDetected
	}
	delete(parityErrors, 0)
	index := 0
	for i := range parityErrors {
		index += i
	}
	out := input ^ (1 << index)
	return out, nil
}

func computeParities(input uint64) map[int]int {
	out := make(map[int]int)
	for i := 64; i > 0; {
		i >>= 1
		if parityFail(input, i) {
			out[i] = 1
		}
	}
	return out
}

//parityFail returns true if there is an error in the parity
func parityFail(input uint64, bit int) bool {
	process := input & parityMask[bit]

	for i := 1; i < 64; i <<= 1 {
		process ^= process >> i
	}
	process &= 1
	return process == 1
}

func removeAllBits(input uint64) uint64 {
	out := input
	for i := 64; i > 0; {
		i >>= 1
		out = bitRemoval(out, i)
	}
	return out
}

func bitRemoval(input uint64, pos int) uint64 {
	left := (input >> (pos + 1)) << pos
	right := (input << (64 - pos)) >> (64 - pos)
	return left | right
}
