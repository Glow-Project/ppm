# ppm

**Pour Package Manager** <br>
Easily install and manage Godot-plugins from GitHub using the command line

## Commands

Initialize a `ppm.json` file

```
$ ppm init
```

<hr>

Either declare a plugin to install or install all plugins that are declared in the `ppm.json` file

```
$ ppm install <plugin>
```

<hr>
Update all your plugins

```
$ ppm update
```

## Installation

```
$ git clone https://github.com/Glow-Project/ppm

$ go mod download

$ go install
```

## Manual installation

Windows:

-   Download the binary
-   Move binary to `\Users\<User>\go\bin`

## Requirements

-   [Git](https://git-scm.com/)
-   [Golang](https://golang.org/)
