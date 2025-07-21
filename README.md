# aurgo

A Go-based command-line tool for interacting with the Arch User Repository (AUR).

## Overview

aurgo is a lightweight and efficient AUR helper written in Go. It provides essential functionality for searching, installing, updating, and removing packages from the AUR with a simple command-line interface.

## Features

- **Search** - Search for packages in the AUR
- **Install** - Install packages from the AUR
- **Update** - Update installed AUR packages
- **Remove** - Remove AUR packages

## Installation

### From Source

```bash
git clone https://github.com/joaogiacometti/aurgo.git
cd aurgo
make install
```

## Usage

### Search for packages

```bash
aurgo -Ss <package-name>
```

### Install a package

```bash
aurgo -S <package-name>
```

### Update packages

```bash
aurgo -U
```

### Remove a package

```bash
aurgo -R <package-name>
```

### Get help

```bash
aurgo --help
```

## Requirements

- Go
- Linux system (Arch Linux or Arch-based distribution)
- Git (for cloning AUR package repositories)
- Base development tools:
  - `make`
  - `gcc` or compatible C compiler
- `pacman` package manager
- Sudo privileges (for package installation)
