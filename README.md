# Hoomy

> Hoomy is a command-line interface (CLI) tool that streamlines the interaction and management of home automation
> systems. Built with Go, Hoomy offers an efficient, platform-independent utility for automating tasks and managing smart
> home devices.

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
<!-- TOC -->

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

## Building with GoReleaser

Hoomy uses GoReleaser for reproducible builds and to manage releases:

```bash
goreleaser release --snapshot --skip-publish --rm-dist
```

Before creating a release, tag your code and push the tag:

```bash
git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push origin vX.Y.Z
```

Use GoReleaser to create the release:

```bash
goreleaser release
```

## Contributing

Your contributions to improve Hoomy are welcome. Open an issue or submit pull requests to collaborate.

## License

Hoomy is released under a specific open-source license. Refer to the `licence` command within the application or view
the `LICENSE` file in this repository for full details.

---

*Hoomy is an independent project and not officially associated with Kizbox or Somfy.*

```