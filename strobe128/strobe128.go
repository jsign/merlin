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

type alignedKeccakState [200]byte

type Strobe128 struct {
	state    alignedKeccakState
	pos      byte
	posBegin byte
	curFlags byte
}

func New(protocolLabel []byte) Strobe128 {
	var st alignedKeccakState
	copy(st[:6], []byte{1, strobeR + 2, 1, 0, 1, 96})
	copy(st[6:], string("STROBEv1.0.2"))
	unsafeKeccakF1600(&st)

	strobe := Strobe128{
		state:    st,
		pos:      0,
		posBegin: 0,
		curFlags: 0,
	}

	strobe.MetaAD(protocolLabel, false)

	return strobe
}

func (s *Strobe128) MetaAD(data []byte, more bool) {
	s.beginOp(flagM|flagA, more)
	s.absorb(data)
}

func (s *Strobe128) AD(data []byte, more bool) {
	s.beginOp(flagA, more)
	s.absorb(data)
}

func (s *Strobe128) PRF(data []byte, more bool) {
	s.beginOp(flagI|flagA|flagC, more)
	s.squeeze(data)
}

func (s *Strobe128) Key(data []byte, more bool) {
	s.beginOp(flagA|flagC, more)
	s.overwrite(data)
}

func (s *Strobe128) squeeze(data []byte) {
	for i := range data {
		data[i] = s.state[s.pos]
		s.state[s.pos] = 0
		s.pos += 1
		if s.pos == strobeR {
			s.runF()
		}
	}
}

func (s *Strobe128) overwrite(data []byte) {
	for i := range data {
		s.state[s.pos] = data[i]
		s.pos += 1
		if s.pos == strobeR {
			s.runF()
		}
	}
}

func (s *Strobe128) beginOp(flags byte, more bool) {
	if more {
		return
	}
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

func unsafeKeccakF1600(state *alignedKeccakState) {
	// This is a tension point between a clean port from Rust and leveraging
	// Go stdlib keccackF1600 implementation. The former express the state as []byte,
	// while the latter as [25]uint64.
	// I'm not a fan of unsafe nor copying 200 bytes to 25 uint64s in hotpaths.
	// You can see who won here :)
	data := (*[25]uint64)(unsafe.Pointer(state))
	keccakF1600(data, 24)
}
