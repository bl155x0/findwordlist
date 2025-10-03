# About

findwordlist is a tiny helper that makes working with wordlists on the command line seamless.
It integrates with fzf to let you search, filter, and pick a wordlist interactively, then automatically stores your choice in the environment variable $W.

Once set, $W can be reused across your favorite tools — no more hunting for long file paths when switching wordlists.

# Prerequisite

* A recent version of `fzf` (https://github.com/junegunn/fzf)
  
# Installation

```
$ go install github.com/bl155x0/findwordlist@v0.3.0
$ export WORDLISTS=/opt/mywordlists
$ alias findwl='eval $(findwordlist)'
```
# Usage

Use `findwl` to search for wordlist

```bash

# select the wordlist
$ findwl

# reference it with $W
$ ffuf -u http://localhost:FUZZ -w $W
```
