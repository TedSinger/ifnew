## Ifnew

### Don't redo that command

Like `make` and every other sane build tool, `ifnew` does a command only if its sources are newer than its targets. Unlike a build tool, `ifnew` is not a build tool. Just a prefix for a shell command

It currently handles:

- `cp`
- `curl`
- `wget`
- `tar`

### Usage

`ifnew cp src dst` will run `cp src dst`. Unless `dst` is newer than `src`. Then it won't

### Roadmap

- `gzip`
