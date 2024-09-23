package bitwisecheck

import (
	lowlevelfunctions "github.com/NikoMalik/low-level-functions"
)

// The Flags type is defined as a byte that allows storing up to 8 different Boolean flags.
// Each bit in the byte can represent a different flag (set or not set).
type Flags byte

// Set sets the specified flag to true.
//
// For example, if flag = 0, it corresponds to 1 << 0, which equals 1 (00000001 in binary).
// If flag = 1, this corresponds to 1 << 1, which equals 2 (00000010 in binary).
// If flag = 2, this corresponds to 1 << 2, which equals 4 (00000100 in binary).
//
// The bitwise OR operator (|=) sets the specified flag to f.  This means that if this flag was not set, it will be set.
// For example, if f = 00000000 and we set flag 0 (1 << 0), f will become 00000001.
// If f = 00000001 and we set flag 0 again, the value of f will not change (it will remain 00000001).
func (f *Flags) Set(flag byte) {
	*f |= Flags(1 << flag)
}

// Clear sets the specified flag to false.  This does not change the value of any other flags.
// The bitwise AND operator (&=) is used to clear the specified flag.  The expression
//
//	f &= ^Flags(1 << flag)
//
// will clear the specified flag, but not change the value of any other flag.
// For example, if f = 00000001 and flag = 0, then f will become 00000000.
// If f = 00000101 and flag = 0, then f will become 00000100.
func (f *Flags) Clear(flag byte) {
	*f &^= Flags(1 << flag)
}

// IsSet returns true if the specified flag is set, false otherwise.
// The expression f&(1<<flag) is a binary operation that compares the value of the
// f byte with the mask (1<<flag).  If the flag is set, the expression will be
// non-zero; if the flag is not set, the expression will be zero.
func (f Flags) IsSet(flag byte) bool {
	return f&(1<<flag) != 0
}

// String returns a string representation of the flags in binary format.
//
// The function works by iterating over each bit in the byte from right to left.
// If the bit is set, '1' is appended to the string; otherwise, '0' is appended.
func (f Flags) String() string {
	var b = lowlevelfunctions.NewStringBuffer(8)
	for i := 7; i >= 0; i-- {
		// Check if the i-th bit is set
		if f&(1<<uint(i)) != 0 {
			b.WriteByte('1')
		} else {
			b.WriteByte('0')
		}
	}
	return b.String()
}
