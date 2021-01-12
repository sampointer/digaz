# digaz [![GoDoc](https://godoc.org/github.com/sampointer/digaz?status.svg)](https://godoc.org/github.com/sampointer/digaz) [![Go Report Card](https://goreportcard.com/badge/github.com/sampointer/digaz)](https://goreportcard.com/report/github.com/sampointer/digaz) ![goreleaser](https://github.com/sampointer/digaz/workflows/goreleaser/badge.svg)

Look-up region and other information for any Azure-owned IP address:

```bash
$ digaz $(dig dev.azure.com +short)
changeNumber: 2 networkFeatures: ["API" "NSG" "UDR" "FW"] platform: "Azure" region: "" regionId: 0 systemService: "AzureFrontDoor"
```

An online version of this tool can be found at [runson.cloud][r].

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
docker run --rm -it digaz $(dig dev.azure.com +short)
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
[r]: https://runson.cloud
