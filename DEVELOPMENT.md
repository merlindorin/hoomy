# Hoomy Development Guide

Welcome to the development guide for Hoomy! This document outlines the steps, tools, and procedures we follow to work
effectively on the Hoomy project. Whether you're a seasoned contributor or new to the project, this guide will help
ensure a smooth development experience.

---

## Table of Contents

<!-- TOC -->
* [Getting Started](#getting-started)
* [Prerequisites](#prerequisites)
* [Setting Up the Development Environment](#setting-up-the-development-environment)
* [Workflow](#workflow)
    * [Development](#development)
    * [Dependencies](#dependencies)
    * [Updating Tools](#updating-tools)
* [Testing](#testing)
* [Code Formatting and Linting](#code-formatting-and-linting)
* [Building and Running](#building-and-running)
* [Documentation](#documentation)
* [Releasing](#releasing)
* [Contributing](#contributing)
<!-- TOC -->

---

### Getting Started

To contribute to Hoomy, you'll first need a local copy of the code. You can obtain it by forking the repository and then
cloning it to your development machine.

```
git clone https://github.com/merlindorin/hoomy.git
cd hoomy
```

### Prerequisites

Ensure you have the following installed on your development machine:

- [Go](https://golang.org/doc/install) (version specified in `go.mod`)
- [Task](https://taskfile.dev) (for running predefined tasks)
- Git (for version control)
- Basic Unix utilities: `curl`, `sh`, `rm`, `cp`

### Setting Up the Development Environment

To bootstrap your environment with the tools needed for development, run:

```
task tools
```

This command will install all required dependencies, including but not limited to linters, formatters, and any other
binaries needed for development and continuous integration (CI).

### Workflow

#### Development

Setting up pre-commit git hooks:

```
task dev
```

This will prepare your git hooks to ensure your code passes tests before committing.

#### Dependencies

To install project dependencies:

```
task setup
```

#### Updating Tools

To update pinned versions of tools:

```
task update-tools
```

### Testing

To run the tests for Hoomy:

```
task test
```

This runs the tests with race detection and coverage reporting.

### Code Formatting and Linting

To format your code:

```
task fmt
```

To lint the code:

```
task lint
```

### Building and Running

To build the project binary:

```
task build
```

To build and run Hoomy:

```
task run
```

Note: Both `build` and `run` tasks accept extra CLI arguments through `CLI_ARGS` environment variable. You can set it
before the command like this:

```
CLI_ARGS="--my-flag" task run

or

task run --my-flag
```

### Documentation

To generate the documentation:

```
task docs:generate
```

To serve the documentation locally:

```
task docs:serve
```

To build the documentation:

```
task docs:build
```

### Releasing

Creating a new release with a new tag:

```
task release
```

Snapshot releases can be done via GoReleaser:

```
task goreleaser:snapshot
```

For actual releases without the snapshot flag:

```
task goreleaser
```

### Contributing

Contributions are highly appreciated! Before contributing, make sure to read the `CONTRIBUTING.md` guide to understand
the contribution process.

For any changes you make, ensure that you:

1. Write meaningful commit messages.
2. Squash your commits into a single commit per feature/bugfix.
3. Before creating a pull request, rebase your changes on the latest `main` branch.

---

Thank you for contributing to Hoomy. Happy developing! 🚀