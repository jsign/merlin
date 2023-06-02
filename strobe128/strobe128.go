package strobe128

import (
	"unsafe"
)

const (
	strobeR = 166

	flagI = 1 << 0
	flagA = 1 << 1
	flagC = 1 << 2
	flagT = 1 << 3
	flagM = 1 << 4
	flagK = 1 << 5
)

type AlignedKeccakState [200]byte

type Strobe128 struct {
	state    AlignedKeccakState
	pos      byte
	posBegin byte
	curFlags byte
}

func New(protocolLabel []byte) *Strobe128 {
	var st AlignedKeccakState
	copy(st[:6], []byte{1, strobeR + 2, 1, 0, 1, 96})
	copy(st[6:], string("STROBEv1.0.2"))
	unsafeKeccakF1600(&st)

	strobe := &Strobe128{
		state:    st,
		pos:      0,
		posBegin: 0,
		curFlags: 0,
	}

	strobe.MetaAD(protocolLabel)

	return strobe
}

func (s *Strobe128) MetaAD(data []byte) {
	s.beginOp(flagM | flagA)
	s.absorb(data)
}

func (s *Strobe128) beginOp(flags byte) {
	oldBegin := s.posBegin
	s.posBegin = s.pos + 1
	s.curFlags = flags

	s.absorb([]byte{oldBegin, flags})

	forceF := (flags & (flagC | flagK)) != 0
	if forceF && s.pos != 0 {
		s.runF()
	}
}

func (s *Strobe128) absorb(data []byte) {
	for i := range data {
		s.state[s.pos] ^= data[i]
		s.pos++
		if s.pos == strobeR {
			s.runF()
		}
	}
}

func (s *Strobe128) runF() {
	s.state[s.pos] ^= s.posBegin
	s.state[s.pos+1] ^= 0x04
	s.state[strobeR+1] ^= 0x80
	unsafeKeccakF1600(&s.state)
	s.pos = 0
	s.posBegin = 0
}

func unsafeKeccakF1600(state *AlignedKeccakState) {
	data := (*[25]uint64)(unsafe.Pointer(state))
	keccakF1600(data, 24)
}
