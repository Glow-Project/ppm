<div align="center">
    <img src="./logo.png" alt=">ppm">
    <br>
    <br>
    <img src="https://img.shields.io/github/license/Glow-Project/ppm" alt="badge">
    <img src="https://img.shields.io/github/issues/Glow-Project/ppm" alt="badge">
    <img src="https://img.shields.io/github/actions/workflow/status/Glow-Project/ppm/ci.yml?branch=main&label=ci" alt="badge">
    <br>
    <br>
    <a href="https://asciinema.org/a/D0cRCFOtekhWOmeC8TxUcpJXo" target="_blank"><img src="https://asciinema.org/a/D0cRCFOtekhWOmeC8TxUcpJXo.svg" /></a>
    <br>
</div>

**Pour Package Manager**

Easily install and manage Godot-plugins from GitHub and the Godot Asset Library using the command line

## Commands

### Initialize

Initialize a `ppm.json` file

```bash
$ ppm init
```

### Install

Either declare a plugin to install or install all plugins that are declared in the `ppm.json` file

#### Install plugin from GitHub

```bash
$ ppm install <user>/<repository>
```

#### Install plugin from the Godot asset library

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

### Main way

Download the binary from the [**release page**](https://github.com/Glow-Project/ppm/releases)

### Installation from Source

```bash
$ git clone https://github.com/Glow-Project/ppm

$ go install
```

### Migration to newer versions

To migrate your ppm project from an older to a newer version you can simply run:

```sh
$ ppm tidy
```

## Create ppm-compatible plugins

1. (optional) Create a new project in Godot and create a directory with the name of your plugin inside the `addons` directory
1. Create the plugin (`plugin.cfg`, `plugin.gd`, ...)
1. (optional) run `ppm init` inside of the directory of your plugin
1. publish your plugin (e.g. `/home/GodotProject/addons/my-plugin`) to GitHub as a public repository
1. Now you can install your plugin via `ppm install <your-username>/<your-repository-name>`

#### References

- [Glow-Project/ppm-ui](https://github.com/Glow-Project/ppm-ui)
- [Glow-Project/SettingsManager](https://github.com/Glow-Project/SettingsManager)
- [Glow-Project/DialogueSystem](https://github.com/Glow-Project/DialogueSystem)
