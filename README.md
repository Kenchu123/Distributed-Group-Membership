# CS425 Distributed Systems MP1

## Description

See [MP Document](./docs/MP2.CS425.FA23.pdf)

## Installation

### Prerequisites

- Go 1.20

### Build

```bash
make build
```

## Usage

### Serve

`serve` command start the command server and wait for commands from clients.

```bash
./bin/msl serve [flags]

Flags:
  -c, --config string   path to config file (default ".msl/config.yml")
  -h, --help            help for serve
  -l, --log string      path to log file (default "logs/msl.log")
  -p, --port string     port to listen on (default "7132")
```

### Join

`member join` command tells the machine to join the group.

```bash
./bin/msl member join [flags]

Flags:
  -h, --help   help for join

Global Flags:
  -c, --config string          path to config file (default ".msl/config.yml")
  -m, --machine-regex string   regex for machines to join (e.g. "0[1-9]") (default ".*")
```

### Leave

`member leave` command tells the machine to leave the group.

```bash
./bin/msl member leave [flags]

Flags:
  -h, --help   help for leave

Global Flags:
  -c, --config string          path to config file (default ".msl/config.yml")
  -m, --machine-regex string   regex for machines to join (e.g. "0[1-9]") (default ".*")
```

### Suspicion

`config set-suspicion` command enables or disables suspicion.

```bash
./bin/msl config set-suspicion [flags]

Flags:
  -e, --enable   enable or disable suspicion
  -h, --help     help for suspicion

Global Flags:
  -c, --config string          path to config file (default ".msl/config.yml")
  -m, --machine-regex string   regex for machines to join (e.g. "0[1-9]") (default ".*")
```

### DropRate

`config set-droprate` command sets the drop rate.

```bash
./bin/msl config set-droprate [flags]

Flags:
  -d, --droprate float32   droprate
  -h, --help               help for droprate

Global Flags:
  -c, --config string          path to config file (default ".msl/config.yml")
  -m, --machine-regex string   regex for machines to join (e.g. "0[1-9]") (default ".*")
```

## Development

### Prerequisites

- Docker
- docker compose

### Set Environment

```bash
# on one session
docker compose -f docker-compose.dev.yml up [-d] [--build]

# on another session
docker exec -it dev /bin/ash
docker exec -it dev-m[1-10] /bin/ash

$ go run main.go [command] [flags]
```

## Contributor

- [Che-Kuang Chu (ckchu2)](https://gitlab.engr.illinois.edu/ckchu2)
- [Jhih-Wei Lin (jhihwei2)](https://gitlab.engr.illinois.edu/jhihwei2)
