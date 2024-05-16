# ggen
A CLI for generating go project, it helps automate the process of creating a new Go project with a predefined directory structure, configuration files, and third party library/package.

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout


## Installation

Install using go
```bash
go install github.com/ryanadiputraa/ggen
```

## Usage
run ggen `generate` command with `name` and `mod` flags to generate go project
```bash
ggen generate -n go-project -m github.com/username/go-project
```
