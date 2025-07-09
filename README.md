# NetCLI

NetCLI is a command-line tool written in Go that provides various network utilities.

## Installation

### Requirements

- Go 1.18 or higher
- For some features like, Administrator/Root privileges may be required

### Build Instructions

```bash
git clone https://github.com/yourname/netcli.git
cd netcli
go build -o netcli.exe
```

### Usage

After building the binary, you can run commands like:

```bash
netcli ping google.com
netcli dns facebook.com
netcli portscan 192.168.1.1
```

To view all available commands:

```bash
netcli --help
```
