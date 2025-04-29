package crc16

type digest struct {
	sum uint16
	t   *Table
}

func (h *digest) Write(data []byte) (int, error) {
	h.sum = Update(h.sum, data, h.t)
	return len(data), nil
}

func (h digest) Sum(b []byte) []byte {
	s := h.Sum16()
	return append(b, byte(s>>8), byte(s))
}

func (h *digest) Reset() {
	h.sum = h.t.params.Init
}

func (h digest) Size() int {
	return 2
}

func (h digest) BlockSize() int {
	return 1
}

func (h digest) Sum16() uint16 {
	return Complete(h.sum, h.t)
}