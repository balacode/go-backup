// -----------------------------------------------------------------------------
// go-backup/storage/bof_mark.go

package storage

// BOFMark is the Beginning-Of-File Marker. A magic sequence of bytes
// written before the metadata and content of every file in the archive.
//
// So, if a part of the archive is partially damaged,
// the program will be able to recover all intact files.
//
var BOFMark = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 8 zeros
	0x23, 0x7E, 0x91, 0x4D, 0x2E, 0xE0, 0x43, 0x1A, // 8 "magic" bytes
}

// end
