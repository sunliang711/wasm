package runtime

import (
	"math/rand"
	"testing"
)

func TestPopcnt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rd := rand.Int31()
		t.Logf("random value: %032b\n", rd)
		ret := doPopcnt32(rd)
		t.Logf("after doPopcnt32() : %v\n", ret)
	}
}

func TestPopcnt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rd := rand.Int63()
		t.Logf("random value: %064b\n", rd)
		ret := doPopcnt64(rd)
		t.Logf("after doPopcnt64() : %v\n", ret)
	}
}
