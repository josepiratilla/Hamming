package hamming

import (
	"math/rand"
	"testing"
)

const nCheck = 200

func TestDecoderNoError(t *testing.T) {

	rand.Seed(1)

	for i := 0; i < nCheck; i++ {
		v := rand.Uint64() & inputMask
		encoded, err := encode(v)
		if err != nil {
			t.Errorf("Error at encoding value 0x%0x\n", v)
		}
		decoded, err := decode(encoded)
		if err != nil {
			t.Errorf("Error at decodingvalue 0x%0x\n", v)
		}
		if decoded != v {
			t.Errorf("Mistmach at encoding, decoding.\nInit Value:0x%0x\nEncoded:0x%0x\nDecoded:0x%0x",
				v, encoded, decoded)
		}
	}
}

func TestDecoderOneErrorCorrected(t *testing.T) {

	rand.Seed(1)

	for i := 0; i < nCheck; i++ {
		v := rand.Uint64() & inputMask
		noise := uint64(1) << rand.Intn(64)
		encoded, err := encode(v)
		if err != nil {
			t.Errorf("Error at encoding value 0x%0x\n", v)
		}
		noiseEncoded := encoded ^ noise
		decoded, err := decode(noiseEncoded)
		if err != nil {
			t.Errorf("Error at decodingvalue 0x%0x\n", v)
		}
		if decoded != v {
			t.Errorf("Mistmach at encoding, decoding with one bit error.\nInit Value:0x%0x\nEncoded:0x%0x\nAfter noise:0x%0x\nDecoded:0x%0x",
				v, encoded, noiseEncoded, decoded)
		}

	}

}

func TestDecoderTwoErrorsDetected(t *testing.T) {
	rand.Seed(1)

	for i := 0; i < nCheck; i++ {
		v := rand.Uint64() & inputMask
		noise := uint64(1) << rand.Intn(64)
		noise ^= uint64(1) << rand.Intn(64)
		if noise == 0 {
			continue
		}
		encoded, err := encode(v)
		if err != nil {
			t.Errorf("Error at encoding value 0x%0x\n", v)
		}
		noiseEncoded := encoded ^ noise
		decoded, err := decode(noiseEncoded)
		if err != errTwoOrMoreErrorsDetected {
			t.Errorf("Error not raised with two errors at transmisssion\nInit value: 0x%0x\nNoise: 0x%0x\nEncoded: 0x%0x\nAfter noise: 0x%0x\nDecoded: 0x%0x\n",
				v, noise, encoded, noiseEncoded, decoded)
		}

	}
}
