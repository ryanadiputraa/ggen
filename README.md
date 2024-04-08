# ggen
A CLI for generating go project that use idiomatic go standards project layout

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout


## Requirements
- Go: 1.21.4

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
