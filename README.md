# About

A helper utility for dealing with wordlists on the command line.

# Usage

Use `findwl` to search for wordlist

```bash

# fiwndwl <SEARCH_TERM>
$ findwl password
```

The `fzf` support allows to seamlessly select a wordlist from the search results.
<img width="1115" height="319" alt="grafik" src="https://github.com/user-attachments/assets/51f68f0a-4205-4499-986a-33e9285cb92e" />


After selecting a wordlist the path to the list is stored in the $W environment variable and can be used with any command.

```bash
# The wordlist can be referenced with $W
$ john --wordlist=$W hash.txt
```
