<img
 width="33%"
 align="right"
 src="https://merlin.cool/merlin.png"/>

## Merlin: composable proof transcripts for public-coin arguments of knowledge
[Merlin][merlin_cool] is a [STROBE][strobe]-based transcript
construction for zero-knowledge proofs. It automates the Fiat-Shamir
transform, so that by using Merlin, non-interactive protocols can be
implemented as if they were interactive.

## What is the goal of this repository?
This is a Go port of the [Rust implementation](https://github.com/zkcrypto/merlin) of the Merlin protocol. 

The goal is to provide a Go implementation of the Merlin protocol that is compatible with the Rust implementation. This will allow for the use of Merlin in Go projects that need to interact with Rust projects that use Merlin.

The implementation tries to be close to the original Rust one, with some minor changes to make it more idiomatic and performant in Go.


## About
Merlin is authored by Henry de Valence, with design input from Isis
Lovecruft and Oleg Andreev.  The construction grew out of work with Oleg
Andreev and Cathie Yun on a [Bulletproofs implementation][bp].
Thanks also to Trevor Perrin and Mike
Hamburg for helpful discussions.  Merlin is named in reference to
[Arthur-Merlin protocols][am_wiki] which introduced the notion of
public coin arguments.

The header image was created by Oleg Andreev as a composite of Arthur Pyle's
[The Enchanter Merlin][merlin_pyle] and the Keccak Team's [θ-step
diagram][keccak_theta].

## License 
MIT

[merlin_cool]: https://merlin.cool
[bp]: https://doc.dalek.rs/bulletproofs/
[strobe]: https://strobe.sourceforge.io/
[am_wiki]: https://en.wikipedia.org/wiki/Arthur%E2%80%93Merlin_protocol
[merlin_pyle]: https://commons.wikimedia.org/wiki/File:Arthur-Pyle_The_Enchanter_Merlin.JPG
[keccak_theta]: https://keccak.team/figures.html