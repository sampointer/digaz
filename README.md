# digaz [![GoDoc](https://godoc.org/github.com/sampointer/digaz?status.svg)](https://godoc.org/github.com/sampointer/digaz) [![Go Report Card](https://goreportcard.com/badge/github.com/sampointer/digaz)](https://goreportcard.com/report/github.com/sampointer/digaz) ![goreleaser](https://github.com/sampointer/digaz/workflows/goreleaser/badge.svg)

Look-up region and other information for any Azure-owned IP address:

```bash
$ digaz $(dig king.com +short)
prefix: 34.64.0.0/10 scope: none service: none
prefix: 34.120.0.0/16 scope: global service: Google Cloud
```

```bash
$ digaz 8.8.8.8 2a00:1450:4009:814::200e
prefix: 8.8.8.0/24 scope: none  service: none
prefix: 2a00:1450::/32 scope: none service: none
```

## Installation

### Homebrew

```bash
brew tap sampointer/digaz
brew install digaz
```

### Packages
Debian and RPM packages can be found on the [releases][1] page.

### Docker

```bash
git clone https://github.com/sampointer/digaz; cd digaz
docker build -t digaz .
docker run --rm -it digaz $(dig king.com +short)
```

## Similar tools

| Company  | Tool        |
|----------|-------------|
| AWS      | [digaws][a] |
| Azure    | [digaz][z]  |
| Google   | [digg][g]   |

[1]: https://github.com/sampointer/digaz/releases/

[a]: https://github.com/sampointer/digaws
[g]: https://github.com/sampointer/digg
[z]: https://github.com/sampointer/digaz
