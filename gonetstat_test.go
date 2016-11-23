package GOnetstat

import (
	"testing"
)

func TestTcp(t *testing.T) {
	p := Tcp()

	for i, v := range p {
		t.Logf("%d -> %#v", i, v)
	}
}

func BenchmarkTcp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tcp()
	}
}
