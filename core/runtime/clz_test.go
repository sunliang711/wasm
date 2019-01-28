package runtime

import "testing"

func TestClz32(t *testing.T) {
	for i := 0; i < 32; i++ {
		t.Log(doClz32(int32(int32(1) << uint32(i))))
	}
}

func TestClz64(t *testing.T) {
	for i := 0; i < 64; i++ {
		t.Log(doClz64(int64(int64(1) << uint64(i))))
	}
}
