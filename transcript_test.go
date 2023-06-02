package transcript

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEquivalenceSimple(t *testing.T) {
	t.Parallel()

	tr := New([]byte("test protocol"))
	tr.AppendMessage([]byte("some label"), []byte("some data"))

	challenge := make([]byte, 32)
	tr.ChallengeBytes([]byte("challenge"), challenge)

	require.Equal(t, "d5a21972d0d5fe320c0d263fac7fffb8145aa640af6e9bca177c03c7efcf0615", hex.EncodeToString(challenge))
}

func TestEquivalenceComplex(t *testing.T) {
	t.Parallel()

	tr := New([]byte("test protocol"))

	tr.AppendMessage([]byte("step1"), []byte("some data"))

	challenge := make([]byte, 32)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = 99
	}

	expChallenges := []string{
		"b61a5540c3dc16e5d746c22e0bfd3509621e37e088a1445581813a746e9c12ea",
		"c66c13d15ff45b5318ee2cef99d3eb8346eb037468090ee32569714617c66c78",
		"cab17754446aabaee9b962006c85457fd61ee1c70ee75a6f92db430bb4c44a7d",
		"15a0d953adf1b55ad2defe664955866200ba2cfb3c1d02209efd63691b07e23e",
		"b7598ad695abf1b8682301d13d3a1208547c79cd00d8fd5c7817bd9beea306d0",
		"3b5518effca8e27942425a8d027223da4c2b86226f247f4cea2df01f7153bc9a",
		"a4964319661dd522f2f9bd676e1c975ca0004a57df92fc5dcafe90d1cd19ab4b",
		"9907ebb80c2c865c82c21642ed7a391a2bfd57510647b073aaa6bd0e50b48941",
		"9de843864ba4eb4bdb9cb17c8081418c4fa6d250576e6be19ca4613c052ce691",
		"663771dec95f0f4df735ad79c20adb5598601a7ec3d96d39ba6be8d6abc19bcc",
		"4b3b2b8646791a5badad81b277b261f5c92403fd70fde4f309bc0b16f51e6bab",
		"a201a470f41038154a7d1d30c5218de1eec441841a5d1f7c562c15fc31ea2746",
		"9a5f590e2494c60db3d5e956fe0ac1e07150bdd0c220f6261aa2d02484911436",
		"d829b1ffeeb4204ab1d2ef7e7337c0382b3c3edf10e4f94073cd1d782c39be04",
		"89149073193054670dfdae73eb3669bdde497597e1e1916a618a7325e8175218",
		"3235b2e321974e9be15b2d6f5a73d4c605a3cc681fec13ab2327d2ebe0bafe32",
		"0bbe994b9f0e9a95a646a67d572d52a11431692c6244db8ed0197cd911a7379b",
		"f6a2b8d174ff192a22bd6f5ed4973a8c86398a6f5f0c1f07a12c70fc8eba7e35",
		"93ec8d8d33c0d5a67c57b1663245be0b065178460180b390e7c3f3c657060aab",
		"fd5ec47965f27ca5decb2275ea3a279bc9b9148b1d55b66d65813808bc32ea83",
		"a3dbffb2c49c9a02ff59baff013f94605511181679c8a6a7a1996f08b1ec2c83",
		"6297acf2d6054b235e90f44223addaab8b990b720c93dc19b60f54006cd81521",
		"d6e5c0db1039ff7d28dfef0135e6387e3da057539a485ea29cf9ada629d50901",
		"16901da1e38bf7de8e27e06d1bd7fb684496d36211ef16982afd498db8ef19c8",
		"7c86ec1f6a97881ada2af2b14bf303d102bc715fc8fa68974782246f5d636211",
		"e7ab040204bffb73564de12914b4d1836852d37698ee5beff1089fa1c10fd2f2",
		"f1d3a50ecf8a8514ed04f5011042979efe44084b8dc33245659c67400dce8686",
		"172c8e5b77c665c2cd9fa8f4fd942dc50137a6e46718d713a0f01cd0f3026ef3",
		"6393df668f1c16f08b1eb4f4d6532681dabe205c2400f9dd271dd7d8ae43e751",
		"84834f888c4d87e4ef962dbf5e11609ccc3126a7c829c0b360023b11b566237c",
		"a3f6a8056fe1e090a0d5a07f2361f8b804a3e45df5c74a81e8c1c430ee7e24f8",
		"a8c933f54fae76e3f9bea93648c1308e7dfa2152dd51674ff3ca438351cf003c",
	}
	for _, expChallenge := range expChallenges {
		tr.ChallengeBytes([]byte("challenge"), challenge)
		require.Equal(t, expChallenge, hex.EncodeToString(challenge))
		tr.AppendMessage([]byte("bigdata"), data)
		tr.AppendMessage([]byte("challengedata"), challenge)
	}
}
