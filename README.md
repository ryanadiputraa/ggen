# ggen

A CLI for generating go project, it helps automate the process of creating a new Go project with a predefined directory [structure](https://github.com/ryanadiputraa/ggen-template), configuration files, and third party library/package.

## Installation

Install using go

```bash
go install github.com/ryanadiputraa/ggen/v2@latest
```

## Usage

run ggen `generate` command with `name` and `mod` flags to generate go project

```bash
ggen generate -n go-project -m github.com/username/go-project
```
