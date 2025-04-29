package crc16


import "hash"

type Hash16 interface {
	hash.Hash
	Sum16() uint16
}

func New(t *Table) Hash16 {
	h := digest{t: t}
	h.Reset()
	return &h
}
