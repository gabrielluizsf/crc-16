package crc16

// Params represents parameters of CRC-16 algorithms.
// More information about algorithms parametrization and parameter descriptions
// can be found here - http://www.zlib.net/crc_v3.txt
type Params struct {
	Poly   uint16
	Init   uint16
	RefIn  bool
	RefOut bool
	XorOut uint16
	Check  uint16
	Name   string
}

// Predefined CRC-16 algorithms.
// List of algorithms with their parameters borrowed from http://reveng.sourceforge.net/crc-catalogue/16.htm
var (
	// ARC (aka CRC-16, CRC-16/LHA, CRC-IBM)
	// Used in: AUTOSAR, LHA, ARC, ZOO archive formats
	ARC = Params{0x8005, 0x0000, true, true, 0x0000, 0xBB3D, "CRC-16/ARC"}

	// AUG-CCITT (aka CRC-16/SPI-FUJITSU)
	// Used in: Fujitsu FlexRay ASSP MB88121B
	AUG_CCITT = Params{0x1021, 0x1D0F, false, false, 0x0000, 0xE5CC, "CRC-16/AUG-CCITT"}
	SPI_FUJITSU = AUG_CCITT

	// BUYPASS (aka CRC-16/UMTS, CRC-16/VERIFONE)
	// Used in: Verifone payment systems, 3GPP UMTS
	BUYPASS = Params{0x8005, 0x0000, false, false, 0x0000, 0xFEE8, "CRC-16/BUYPASS"}
	UMTS = BUYPASS
	VERIFONE = BUYPASS

	// CCITT-FALSE (aka CRC-16/AUTOSAR, CRC-16/IBM-3740)
	// Used in: IBM 3740 floppy disks, AUTOSAR
	CCITT_FALSE = Params{0x1021, 0xFFFF, false, false, 0x0000, 0x29B1, "CRC-16/CCITT-FALSE"}
	IBM_3740 = CCITT_FALSE

	// CDMA2000
	// Used in: 3GPP2 CDMA2000 systems
	CDMA2000 = Params{0xC867, 0xFFFF, false, false, 0x0000, 0x4C06, "CRC-16/CDMA2000"}

	// CMS
	// Used in: CERN CMS Online Software, Samsung mobiles
	CMS = Params{0x8005, 0xFFFF, false, false, 0x0000, 0xAEE7, "CRC-16/CMS"}

	// CRC-A (aka CRC-16/ISO-IEC-14443-3-A)
	// Used in: Contactless IC cards (ISO/IEC 14443-3)
	CRC_A = Params{0x1021, 0xC6C6, true, true, 0x0000, 0xBF05, "CRC-16/CRC-A"}

	// DDS-110
	// Used in: ELV DDS-110 function generator
	DDS_110 = Params{0x8005, 0x800D, false, false, 0x0000, 0x9ECF, "CRC-16/DDS-110"}

	// DECT-R (aka R-CRC-16)
	// Used in: DECT A-fields
	DECT_R = Params{0x0589, 0x0000, false, false, 0x0001, 0x007E, "CRC-16/DECT-R"}

	// DECT-X (aka X-CRC-16)
	// Used in: DECT B-fields
	DECT_X = Params{0x0589, 0x0000, false, false, 0x0000, 0x007F, "CRC-16/DECT-X"}

	// DNP
	// Used in: Distributed Network Protocol
	DNP = Params{0x3D65, 0x0000, true, true, 0xFFFF, 0xEA82, "CRC-16/DNP"}

	// EN-13757
	// Used in: Wireless M-Bus protocol
	EN_13757 = Params{0x3D65, 0x0000, false, false, 0xFFFF, 0xC2B7, "CRC-16/EN-13757"}

	// GENIBUS (aka CRC-16/DARC, CRC-16/EPC, CRC-16/I-CODE)
	// Used in: RFID tags, DARC, EPC, I-CODE
	GENIBUS = Params{0x1021, 0xFFFF, false, false, 0xFFFF, 0xD64E, "CRC-16/GENIBUS"}

	// GSM
	// Used in: GSM networks, ECMA-130 (CD-ROM)
	GSM = Params{0x1021, 0x0000, false, false, 0xFFFF, 0xCE3C, "CRC-16/GSM"}

	// IBM-SDLC (aka CRC-16/ISO-HDLC, CRC-16/X-25, CRC-B, X-25)
	// Used in: HDLC, X.25, V.42, PPP
	IBM_SDLC = Params{0x1021, 0xFFFF, true, true, 0xFFFF, 0x906E, "CRC-16/IBM-SDLC"}
	ISO_HDLC = IBM_SDLC
	X_25 = IBM_SDLC

	// KERMIT (aka CRC-16/CCITT, CRC-16/CCITT-TRUE, CRC-16/V-41-LSB)
	// Used in: Kermit protocol, Bluetooth
	KERMIT = Params{0x1021, 0x0000, true, true, 0x0000, 0x2189, "CRC-16/KERMIT"}

	// LJ1200
	// Used in: LoJack vehicle tracking system
	LJ1200 = Params{0x6F63, 0x0000, false, false, 0x0000, 0xBDF4, "CRC-16/LJ1200"}

	// M17
	// Used in: M17 digital voice protocol
	M17 = Params{0x5935, 0xFFFF, false, false, 0x0000, 0x772B, "CRC-16/M17"}

	// MAXIM (aka CRC-16/MAXIM-DOW)
	// Used in: Maxim/Dallas 1-Wire buses
	MAXIM = Params{0x8005, 0x0000, true, true, 0xFFFF, 0x44C2, "CRC-16/MAXIM"}

	// MCRF4XX
	// Used in: Microchip RFID devices
	MCRF4XX = Params{0x1021, 0xFFFF, true, true, 0x0000, 0x6F91, "CRC-16/MCRF4XX"}

	// MODBUS
	// Used in: MODBUS protocol
	MODBUS = Params{0x8005, 0xFFFF, true, true, 0x0000, 0x4B37, "CRC-16/MODBUS"}

	// NRSC-5
	// Used in: HD Radio (NRSC-5 standard)
	NRSC_5 = Params{0x080B, 0xFFFF, true, true, 0x0000, 0xA066, "CRC-16/NRSC-5"}

	// OPENSAFETY-A
	// Used in: OpenSAFETY protocol
	OPENSAFETY_A = Params{0x5935, 0x0000, false, false, 0x0000, 0x5D38, "CRC-16/OPENSAFETY-A"}

	// OPENSAFETY-B
	// Used in: OpenSAFETY protocol
	OPENSAFETY_B = Params{0x755B, 0x0000, false, false, 0x0000, 0x20FE, "CRC-16/OPENSAFETY-B"}

	// PROFIBUS (aka CRC-16/IEC-61158-2)
	// Used in: PROFIBUS networks
	PROFIBUS = Params{0x1DCF, 0xFFFF, false, false, 0xFFFF, 0xA819, "CRC-16/PROFIBUS"}

	// RIELLO
	// Used in: Riello Dialog UPS
	RIELLO = Params{0x1021, 0xB2AA, true, true, 0x0000, 0x63D0, "CRC-16/RIELLO"}

	// T10-DIF
	// Used in: SCSI Data Integrity Field
	T10_DIF = Params{0x8BB7, 0x0000, false, false, 0x0000, 0xD0DB, "CRC-16/T10-DIF"}

	// TELEDISK
	// Used in: Teledisk disc archive format
	TELEDISK = Params{0xA097, 0x0000, false, false, 0x0000, 0x0FB3, "CRC-16/TELEDISK"}

	// TMS37157
	// Used in: Texas Instruments TMS37157
	TMS37157 = Params{0x1021, 0x89EC, true, true, 0x0000, 0x26B1, "CRC-16/TMS37157"}

	// USB
	// Used in: USB protocol
	USB = Params{0x8005, 0xFFFF, true, true, 0xFFFF, 0xB4C8, "CRC-16/USB"}

	// XMODEM (aka CRC-16/ACORN, CRC-16/LTE, CRC-16/V-41-MSB, ZMODEM)
	// Used in: XMODEM, ZMODEM, MMC, LTE
	XMODEM = Params{0x1021, 0x0000, false, false, 0x0000, 0x31C3, "CRC-16/XMODEM"}
)