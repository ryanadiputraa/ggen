# ggen
A CLI for generating go project that use idiomatic go project standard layout

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout


## Requirements
- Go: 1.21.4

## Installation

clone ggen repository
```bash
git clone git@github.com:ryanadiputraa/ggen.git
```

build and install ggen go binary
```bash
cd ggen && go build && go install
```

## Usage
run ggen `generate` command with `name` and `mod` flags to generate go project
```bash
ggen generate -n go-project -m github.com/ryanadiputraa/ggen
```
