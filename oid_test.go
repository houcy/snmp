package snmp

import (
	"bytes"
	"testing"
)

func TestEncodeAndParseOID(t *testing.T) {
	oid := ObjectIdentifier{1, 3, 6, 1, 4, 1, 2636, 3, 2, 3, 1, 20}

	b, err := oid.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if expected := []byte{
		0x6, 0x0c,

		0x2b, 0x06, 0x01, 0x04,
		0x01, 0x94, 0x4c, 0x03,
		0x02, 0x03, 0x01, 0x14,
	}; !bytes.Equal(expected, b) {
		t.Errorf("encoded ObjectIdentifer incorrect. Expected %v, got %v", expected, b)
	}

	parsed := MustParseOID(oid.String())
	if oid.String() != parsed.String() {
		t.Errorf("expected parsed ObjectIdentifer %v, got %v", oid, parsed)
	}
}

func TestOIDLargeNumbers(t *testing.T) {
	oid := MustParseOID(".1.3.6.1.2.1.7.7.1.8.1.4.0.0.0.0.68.1.4.0.0.0.0.0.2464081")

	if oid.String() != ".1.3.6.1.2.1.7.7.1.8.1.4.0.0.0.0.68.1.4.0.0.0.0.0.2464081" {
		t.Errorf("expected ObjectIdentifer %s, got %s",
			".1.3.6.1.2.1.7.7.1.8.1.4.0.0.0.0.68.1.4.0.0.0.0.0.2464081", oid.String())
	}
}
