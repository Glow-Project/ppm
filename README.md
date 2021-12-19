# ppm

![badge](https://img.shields.io/github/license/Glow-Project/ppm)
![badge](https://img.shields.io/github/issues/Glow-Project/ppm)

**Pour Package Manager** <br>
Easily install and manage Godot-plugins from GitHub using the command line

## Commands

### Initialize

Initialize a `ppm.json` file

```bash
$ ppm init
```

### Install

Either declare a plugin to install or install all plugins that are declared in the `ppm.json` file

```bash
$ ppm install <plugin>
```

### Update

Update all your plugins

```bash
$ ppm update
```

### Uninstall

Uninstall a dependency

```bash
$ ppm uninstall <plugin>
```

### More

For further information and commands use

```bash
$ ppm -h
```

## Installation

```bash
$ git clone https://github.com/Glow-Project/ppm

$ go mod download

$ go install
```

## Manual installation

### Windows:

- Download the binary
- Move binary to a directory that is part of the `$PATH` variable

### Mac/Linux:

- Download the binary
- Move binary to `/usr/local/bin`

## Requirements

- [**Git**](https://git-scm.com/) _Only needed with [v1.0.1](https://github.com/Glow-Project/ppm/releases/tag/1.0.1) or lower_
- [**Golang**](https://golang.org/)
