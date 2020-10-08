package hamming

import "testing"

func TestPartialDisplacement(t *testing.T) {
	testValues := []struct {
		input    uint64
		expected uint64
		position int
	}{
		{
			0b0000000000000000000000000000000000000000000000000000000001101101,
			0b0000000000000000000000000000000000000000000000000000000011010101,
			3,
		},
		{
			0b0000000000000000011111111111111111111111111111111111111111111111,
			0b0000000000000000111111111111111111111111111111111111111111111110,
			0,
		},
		{
			0b0000000000000000000000000000000000000000000000000000000001101101,
			0b0000000000000000000000000000000000000000000000000000000001101101,
			10,
		},
	}
	for _, testValue := range testValues {
		actual := partialDisplacement(testValue.input, testValue.position)
		if actual != testValue.expected {
			t.Errorf("partialDisplacement(0x%016x,%d)\nReturned:0x%016x\nExpected:0x%016x\n",
				testValue.input, testValue.position, actual, testValue.expected)
		}
	}

}

func TestReallocateBits(t *testing.T) {
	vs := []struct {
		input    uint64
		expected uint64
	}{
		{
			0b0000000000000000000000000000000000000000000000000000000001101101,
			0b0000000000000000000000000000000000000000000000000000110011001000,
		},
		{
			0b0000000111111111111111111111111111111111111111111111111111111111,
			0b1111111111111111111111111111111011111111111111101111111011101000,
		},
	}
	for _, v := range vs {
		actual, err := reallocateBits(v.input)
		if err != nil {
			t.Errorf("reallocateBits(0x%016x) is returning an error\n", v.input)
		}
		if actual != v.expected {
			t.Errorf("reallocateBits(0x%016x)\nReturned:0x%016x\nExpected:0x%016x\n",
				v.input, actual, v.expected)
		}
	}
}

func TestReallocateBitsError(t *testing.T) {
	input := uint64(0b0000001111111111111111111111111111111111111111111111111111111111)
	_, err := reallocateBits(input)
	if err != errInputTooBig {
		t.Error("Reallocate bits is not returning the correct error when receiving a number too big")
	}
}

func TestAddPartityBit(t *testing.T) {
	vs := []struct {
		input    uint64
		pos      int
		expected uint64
	}{
		{
			0b1111111111111111111111111111111011111111111111101111111011101000,
			1,
			0b1111111111111111111111111111111011111111111111101111111011101010,
		},
		{
			0b0000000000000000000000000000000000000000000000000000000000000000,
			3,
			0b0000000000000000000000000000000000000000000000000000000000000000,
		},
	}
	for _, v := range vs {
		actual := addParityBit(v.input, v.pos)

		if actual != v.expected {
			t.Errorf("addParityBit(0x%016x,%d)\nReturned:0x%016x\nExpected:0x%016x\n",
				v.input, v.pos, actual, v.expected)
		}

	}
}

func TestEncode(t *testing.T) {
	vs := []struct {
		input    uint64
		expected uint64
	}{
		{
			0b0000000111111111111111111111111111111111111111111111111111111111,
			0b1111111111111111111111111111111111111111111111111111111111111111,
		},
		{
			0b0000000000000000000000000000000000000000000000000000000000000000,
			0b0000000000000000000000000000000000000000000000000000000000000000,
		},
		{
			0b0000000100100101110010010011001100100110101001100001110011000010,
			0b1001001011100100100110011001001010101001100001101001100000110101,
		},
	}
	for _, v := range vs {
		actual, _ := encode(v.input)

		if actual != v.expected {
			t.Errorf("encode(0x%016x)\nReturned:0x%016x\nExpected:0x%016x\n",
				v.input, actual, v.expected)
		}

	}
}

func TestEncodeError(t *testing.T) {
	input := uint64(0b0000001111111111111111111111111111111111111111111111111111111111)
	_, err := encode(input)
	if err != errInputTooBig {
		t.Error("Reallocate bits is not returning the correct error when receiving a number too big")
	}
}
