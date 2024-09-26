## Minimake

### Don't redo that command

Like `make` and every other sane build tool, `minimake` skips a command if its results are newer than its sources. Unlike a build tool, `minimake` is not a build tool. Just a prefix for a shell command

It currently handles:

- `cp`
- `curl`
- `wget`

### Usage

`mm cp src dst` will run `cp src dst`. Unless `dst` is newer than `src`. Then it won't