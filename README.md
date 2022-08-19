# `spongo`: a `sponge(1)` replacement written in Go
The `moreutils` package for Linux contains a nifty utility called [`sponge`](https://linux.die.net/man/1/sponge):

> sponge reads standard input and writes it out to the specified file. Unlike a shell redirect, sponge soaks up all its input before opening the output file. This allows constricting pipelines that read from and write to the same file.
>
> If no output file is specified, sponge outputs to stdout.

Sadly I couldn't easily find it for macOS, so I tracked down the [source code for `moreutils`](https://joeyh.name/code/moreutils/) and then did something similar in Go.

# License
[MIT](LICENSE).
