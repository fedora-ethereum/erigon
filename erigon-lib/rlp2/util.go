// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package rlp

type Token int32

func (T Token) String() string {
	switch T {
	case TokenDecimal:
		return "decimal"
	case TokenShortBlob:
		return "short_blob"
	case TokenLongBlob:
		return "long_blob"
	case TokenShortList:
		return "short_list"
	case TokenLongList:
		return "long_list"
	case TokenEOF:
		return "eof"
	case TokenUnknown:
		return "unknown"
	default:
		return "nan"
	}
}

func (T Token) Plus(n byte) byte {
	return byte(T) + n
}

func (T Token) Diff(n byte) byte {
	return n - byte(T)
}

func (T Token) IsListType() bool {
	return T == TokenLongList || T == TokenShortList
}

func (T Token) IsBlobType() bool {
	return T == TokenLongBlob || T == TokenShortBlob
}

const (
	TokenDecimal   Token = 0x00
	TokenShortBlob Token = 0x80
	TokenLongBlob  Token = 0xb7
	TokenShortList Token = 0xc0
	TokenLongList  Token = 0xf7

	TokenUnknown Token = 0xff01
	TokenEOF     Token = 0xdead
)

func identifyToken(b byte) Token {
	switch {
	case b <= 127:
		return TokenDecimal
	case b >= 128 && b <= 183:
		return TokenShortBlob
	case b >= 184 && b <= 191:
		return TokenLongBlob
	case b >= 192 && b <= 247:
		return TokenShortList
	case b >= 248:
		return TokenLongList
	}
	return TokenUnknown
}

// BeInt parses Big Endian representation of an integer from given payload at given position
func nextBeInt(w *buf, length int) (int, error) {
	dat, err := nextFull(w, length)
	if err != nil {
		return 0, ErrUnexpectedEOF
	}
	return BeInt(dat, 0, length)
}

func nextFull(dat *buf, size int) ([]byte, error) {
	d := dat.Next(size)
	if len(d) != size {
		return nil, ErrUnexpectedEOF
	}
	return d, nil
}
