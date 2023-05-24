# Simple Command Line Interface

**Build**:

``` bash
# Windows
go build -o simple-cli.exe main.go
# Linux
go build -o simple-cli main.go
```

**Run**:

``` bash
# help
./simple-cli -h

Simple-cli is a CLI tool based on Cobra.
This application is a demo to show how to quickly create a Cobra application.

Usage:
  simple-cli [flags]
  simple-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  echo        Display a line of text
  help        Help about any command
  version     Version prints the information for simple-cli

Flags:
  -h, --help      help for simple-cli
  -v, --version   print simple-cli version

Use "simple-cli [command] --help" for more information about a command.

# version
./simple-cli version
simple-cli v1.0.0

# use shorthand flag
./simple-cli -v
simple-cli v1.0.0

# echo
./simple-cli echo 123 456
123 456

# use shorthand flag
./simple-cli echo -f 123 456
456 123

# use flag
./simple-cli echo --flip 123 456
456 123
```
