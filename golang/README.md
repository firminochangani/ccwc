# ccwc/golang

A reimplementation of [wc](https://en.wikipedia.org/wiki/Wc_(Unix)) as based on [Coding Challenges / Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc).

## Instructions

- Build a binary by running `make build` which will save the binary in `bin/ccwc`
- Running commands against a file: `bin/ccwc -c path/to/file`
- Running commands against the standard input: `cat path/to/file | bin/ccwc -c`

## Commands implemented

- `-c` - The number of bytes in each input file is written to the standard output.
- `-l` - The number of lines in each input file is written to the standard output.
- `-m` - The number of characters in each input file is written to the standard output.
-  `w` - The number of words in each input file is written to the standard output.
