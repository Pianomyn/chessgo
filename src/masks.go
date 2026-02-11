package main

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Prevent Wrap arounds
const (
	FileAMask Bitboard = 0x0101010101010101
	FileBMask Bitboard = FileAMask << 1
	FileCMask Bitboard = FileAMask << 2
	FileDMask Bitboard = FileAMask << 3
	FileEMask Bitboard = FileAMask << 4
	FileFMask Bitboard = FileAMask << 5
	FileGMask Bitboard = FileAMask << 6
	FileHMask Bitboard = FileAMask << 7

	Rank1Mask Bitboard = 0x00000000000000FF
	Rank2Mask Bitboard = Rank1Mask << 8
	Rank3Mask Bitboard = Rank1Mask << 16
	Rank4Mask Bitboard = Rank1Mask << 24
	Rank5Mask Bitboard = Rank1Mask << 32
	Rank6Mask Bitboard = Rank1Mask << 40
	Rank7Mask Bitboard = Rank1Mask << 48
	Rank8Mask Bitboard = Rank1Mask << 56
)

const (
	NotFileABMask = ^(FileAMask | FileBMask)
	NotFileGHMask = ^(FileGMask | FileHMask)
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
