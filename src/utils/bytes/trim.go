package bytes

import "bytes"

// TrimAfter returns a byte slice with the first appearance of `c` and all its trailing bytes removed.
func TrimAfter(s []byte, sep byte) []byte {
	pos := bytes.IndexByte(s, sep)
	if pos == -1 {
		return s
	}
	return s[:pos]
}

// TrimC returns a byte slice with the first appearance of `\x00` (null character) and all its trailing bytes removed.
func TrimC(s []byte) []byte {
	return TrimAfter(s, '\x00')
}