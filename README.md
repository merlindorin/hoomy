<div style="text-align: center">
<img src="./assets/logo.png" style="padding: 15px; background: whitesmoke; border-radius: 25%; width: 200px; color-" />

# Hoomy

> Hoomy is a command-line interface (CLI) tool that streamlines the interaction and management of home automation
> systems. Built with Go, Hoomy offers an efficient, platform-independent utility for automating tasks and managing smart
> home devices.

</div>

---
## Table of Contents

<!-- TOC -->
* [Hoomy](#hoomy)
  * [Features](#features)
  * [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [Usage](#usage)
      * [Commands](#commands)
    * [Verifying the Binary](#verifying-the-binary)
  * [Docker](#docker)
  * [Building with GoReleaser](#building-with-goreleaser)
  * [Contributing](#contributing)
  * [License](#license)
  * [References](#references)
<!-- TOC -->
---

## Features

- Control Venetian blinds through an easy-to-use CLI (list, set, open, close)
- List devices within your home automation system
- Real-time event monitoring
- System discovery for straightforward device integration
- Secure API interactions with Kizbox enabled
- Configuration via environment variables or command-line flags

## Getting Started

### Prerequisites

- An operational Go environment
- A valid API key from Kizbox, obtained as described in the distributor guidelines or through
  the [Somfy TaHoma Developer Mode guide](https://github.com/Somfy-Developer/Somfy-TaHoma-Developer-Mode)

### Installation

To install Hoomy using Go:

```bash
go get github.com/merlindorin/hoomy
```

Alternatively, download the latest pre-built binary from Hoomy's GitHub Releases page.

### Usage

To run Hoomy:

```bash
hoomy <command>
```

For command usage and options:

```bash
hoomy --help
```

#### Commands

- `version` - Displays version information.
- `licence` - Print licensing information.
- `venitian list` - List Venetian blinds in the network.
- `venitian set` - Adjust Venetian blinds to a specified position.
- `venitian open` - Open Venetian blinds.
- `venitian close` - Close Venetian blinds.
- `devices list` - Display a list of all connected devices.
- `listen` - Listen for and print events in real-time.
- `discover` - Discover communicable systems on the network.

### Verifying the Binary

It's important to validate the integrity and authenticity of Hoomy's binary. Follow these steps:

```bash
# Download checksums and signature 
wget https://github.com/merlindorin/hoomy/releases/download/vX.Y.Z/checksums.txt
wget https://github.com/merlindorin/hoomy/releases/download/vX.Y.Z/checksums.txt.sig

# Download the signing certificate
wget https://github.com/merlindorin/hoomy/releases/download/vX.Y.Z/checksums.txt.pem

# Verify the signature against the checksum file
cosign verify-blob --signature checksums.txt.sig --cert checksums.txt.pem checksums.txt

# Upon successful verification, download the appropriate binary
wget https://github.com/merlindorin/hoomy/releases/download/vX.Y.Z/hoomy_vX.Y.Z_linux_amd64.tar.gz

# Now, confirm the SHA256 checksum
sha256sum --ignore-missing -c checksums.txt
```

Replace `vX.Y.Z` with the actual version of Hoomy that you're downloading.

## Docker

To use Hoomy from a pre-built Docker image hosted on GitHub Container Registry:

```bash
docker pull ghcr.io/merlindorin/hoomy:latest
docker run --rm -it ghcr.io/merlindorin/hoomy:latest <command>
```

## Development

All information related to development can be found in [`DEVELOPMENT.md`](./DEVELOPMENT.md).

## Contributing

Your contributions to improve Hoomy are welcome. Open an issue or submit pull requests to collaborate.

## License

Hoomy is released under a specific open-source license. Refer to the `licence` command within the application or view
the `LICENSE` file in this repository for full details.

## References

- Icons: https://mageicons.com/
---

*Hoomy is an independent project and not officially associated with Kizbox or Somfy.*
