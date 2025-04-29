package crc16

import (
	"fmt"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestHash(t *testing.T) {
	tbl := MakeTable(XMODEM)
	h := New(tbl)

	fmt.Fprint(h, "standard")
	fmt.Fprint(h, " library hash interface")
	sum1 := h.Sum16()
	h.Reset()
	fmt.Fprint(h, "standard library hash interface")
	sum2 := h.Sum16()
	assert.Equal(t, sum1, sum2)

	var crc uint16 = 0xe698
	assert.Equal(t, sum1, crc)

	assert.Equal(t, h.Size(), 2)
	buf := make([]byte, 0, 10)
	buf = h.Sum(buf)
	expected := []byte{0xe6, 0x98}
	assert.False(t, len(buf) != 2 || buf[0] != expected[0] || buf[1] != expected[1])
	assert.Equal(t, h.BlockSize(), 1)
}
