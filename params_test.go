package crc16

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestARC(t *testing.T) {
	testSelectedCRC(t, ARC)
}

func TestAUG_CCIT(t *testing.T) {
	testSelectedCRC(t, AUG_CCITT)
}

func TestBUYPASS(t *testing.T) {
	testSelectedCRC(t, BUYPASS)
}

func TestCCITT_FALSE(t *testing.T) {
	testSelectedCRC(t, CCITT_FALSE)
}

func TestCDMA2000(t *testing.T) {
	testSelectedCRC(t, CDMA2000)
}

func TestDDS_110(t *testing.T) {
	testSelectedCRC(t, DDS_110)
}

func TestDECT_R(t *testing.T) {
	testSelectedCRC(t, DECT_R)
}

func TestDECT_X(t *testing.T) {
	testSelectedCRC(t, DECT_X)
}

func TestDNP(t *testing.T) {
	testSelectedCRC(t, DNP)
}

func TestEN_13757(t *testing.T) {
	testSelectedCRC(t, EN_13757)
}

func TestGENIBUS(t *testing.T) {
	testSelectedCRC(t, GENIBUS)
}

func TestMAXIM(t *testing.T) {
	testSelectedCRC(t, MAXIM)
}

func TestMCRF4XX(t *testing.T) {
	testSelectedCRC(t, MCRF4XX)
}

func TestRIELLO(t *testing.T) {
	testSelectedCRC(t, RIELLO)
}

func TestT10_DIF(t *testing.T) {
	testSelectedCRC(t, T10_DIF)
}

func TestTELEDISK(t *testing.T) {
	testSelectedCRC(t, TELEDISK)
}

func TestTMS37157(t *testing.T) {
	testSelectedCRC(t, TMS37157)
}

func TestUSB(t *testing.T) {
	testSelectedCRC(t, USB)
}

func TestCRC_A(t *testing.T) {
	testSelectedCRC(t, CRC_A)
}

func TestKERMIT(t *testing.T) {
	testSelectedCRC(t, KERMIT)
}

func TestMODBUS(t *testing.T) {
	testSelectedCRC(t, MODBUS)
}

func TestX_25(t *testing.T) {
	testSelectedCRC(t, X_25)
}

func TestXMODEM(t *testing.T) {
	testSelectedCRC(t, XMODEM)
}

func TestCMS(t *testing.T) {
	testSelectedCRC(t, CMS)
}

func TestGSM(t *testing.T) {
	testSelectedCRC(t, GSM)
}

func TestIBM_3740(t *testing.T) {
	testSelectedCRC(t, IBM_3740)
}

func TestIBM_SDLC(t *testing.T) {
	testSelectedCRC(t, IBM_SDLC)
}

func TestISO_HDLC(t *testing.T) {
	testSelectedCRC(t, ISO_HDLC)
}

func TestLJ1200(t *testing.T) {
	testSelectedCRC(t, LJ1200)
}

func TestM17(t *testing.T) {
	testSelectedCRC(t, M17)
}

func TestNRSC_5(t *testing.T) {
	testSelectedCRC(t, NRSC_5)
}

func TestOPENSAFETY_A(t *testing.T) {
	testSelectedCRC(t, OPENSAFETY_A)
}

func TestOPENSAFETY_B(t *testing.T) {
	testSelectedCRC(t, OPENSAFETY_B)
}

func TestProfibius(t *testing.T) {
	testSelectedCRC(t, PROFIBUS)
}

func TestSPI_FUJITSU(t *testing.T) {
	testSelectedCRC(t, SPI_FUJITSU)
}

func TestUMTS(t *testing.T) {
	testSelectedCRC(t, UMTS)
}

func TestVerifone(t *testing.T) {
	testSelectedCRC(t, VERIFONE)
}

var testData = []byte("123456789")

func testSelectedCRC(t *testing.T, params Params) {
	table := MakeTable(params)
	assert.NotNil(t, table)

	crc := Checksum(testData, table)

	assert.Equal(t, crc, table.params.Check)
}
