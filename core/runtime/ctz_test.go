package runtime

import "testing"

func TestCtz32(t *testing.T) {
	for i := 0; i < 32; i++ {
		t.Log(doCtz32(int32(int32(1) << uint32(i))))
	}
}

func TestCtz64(t *testing.T) {
	for i := 0; i < 64; i++ {
		t.Log(doCtz64(int64(int64(1) << uint64(i))))
	}
}
