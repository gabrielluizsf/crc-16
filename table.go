package crc16

import "math/bits"

// Table is a 256-word table representing polinomial and algorithm settings for efficient processing.
type Table struct {
	params Params
	data   [256]uint16
}

// MakeTable returns the Table constructed from the specified params.
func MakeTable(params Params) *Table {
	table := new(Table)
	table.params = params
	for n := range 256 {
		crc := uint16(n) << 8
		for range 8 {
			bit := (crc & 0x8000) != 0
			crc <<= 1
			if bit {
				crc ^= params.Poly
			}
		}
		table.data[n] = crc
	}
	return table
}

// Init returns the initial value for CRC register corresponding to the specified algorithm.
func Init(table *Table) uint16 {
	return table.params.Init
}

// Update returns the result of adding the bytes in data to the crc.
func Update(crc uint16, data []byte, table *Table) uint16 {
	for _, d := range data {
		if table.params.RefIn {
			d = bits.Reverse8(d)
		}
		crc = crc<<8 ^ table.data[byte(crc>>8)^d]
	}
	return crc
}

// Complete returns the result of CRC calculation and post-calculation processing of the crc.
func Complete(crc uint16, table *Table) uint16 {
	if table.params.RefOut {
		return bits.Reverse16(crc) ^ table.params.XorOut
	}
	return crc ^ table.params.XorOut
}

// Checksum returns CRC checksum of data usign scpecified algorithm represented by the Table.
func Checksum(data []byte, table *Table) uint16 {
	crc := Init(table)
	crc = Update(crc, data, table)
	return Complete(crc, table)
}
