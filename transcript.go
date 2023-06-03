package transcript

import (
	"encoding/binary"
	"math/rand"

	"github.com/jsign/merlin/strobe128"
)

var (
	labelMerlinProtocol  = []byte("Merlin v1.0")
	labelRng             = []byte("rng")
	labelDomainSeparator = []byte("dom-sep")
)

type Transcript struct {
	str strobe128.Strobe128
}

// New returns a new Transcript object.
func New(label []byte) *Transcript {
	tr := &Transcript{
		str: strobe128.New(labelMerlinProtocol),
	}
	tr.AppendMessage(labelDomainSeparator, label)

	return tr
}

func (t *Transcript) AppendMessage(label []byte, message []byte) {
	var dataLen [4]byte
	binary.LittleEndian.PutUint32(dataLen[:], uint32(len(message)))
	t.str.MetaAD(label, false)
	t.str.MetaAD(dataLen[:], true)
	t.str.AD(message, false)
}

func (t *Transcript) AppendU64(label []byte, x uint64) {
	var xbytes [8]byte
	binary.LittleEndian.PutUint64(xbytes[:], x)
	t.AppendMessage(label, xbytes[:])
}

func (t *Transcript) ChallengeBytes(label []byte, dest []byte) {
	var dataLen [4]byte
	binary.LittleEndian.PutUint32(dataLen[:], uint32(len(dest)))
	t.str.MetaAD(label, false)
	t.str.MetaAD(dataLen[:], true)
	t.str.PRF(dest, false)
}

type TranscriptRngBuilder struct {
	str strobe128.Strobe128
}

func (trb *TranscriptRngBuilder) RekeyWithWitnessBytes(label []byte, witness []byte) {
	var dataLen [4]byte
	binary.LittleEndian.PutUint32(dataLen[:], uint32(len(witness)))
	trb.str.MetaAD(label, false)
	trb.str.MetaAD(dataLen[:], true)
	trb.str.Key(witness, false)
}

func (trb *TranscriptRngBuilder) Finalize(rand rand.Rand) Transcript {
	var randbytes [32]byte
	rand.Read(randbytes[:])
	trb.str.MetaAD(labelRng, false)
	trb.str.Key(randbytes[:], false)

	return Transcript{
		// A copy of trb.str is enough since the internal sturct
		// contains no pointers.
		str: trb.str,
	}
}
