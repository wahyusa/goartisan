[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) [![Go Artisan CI](https://github.com/wahyusa/goartisan/actions/workflows/go.yml/badge.svg)](https://github.com/wahyusa/goartisan/actions/workflows/go.yml)

# Go Artisan

[Roadmap to v1.0.0](https://github.com/users/wahyusa/projects/12/views/1)

Go Artisan is a CLI tool designed to simplify and accelerate the development of RESTful APIs in Go with artisan like command.

Built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)

The code generated will using Gin, GORM, MySQL and godotenv with repository pattern, but you can modify it based on each template.

## FAQ

#### Why using goartisan if I can just scaffold any initial project with `git clone` ?

Well, by using go artisan you will have `config.toml` that you can customize and replace `templates` to fit your own need and initialize new RESTful API project faster by using that templates.

This is also a good fit for terminal enjoyers who love to do many things directly in terminal instead of creating each file manually one by one and `alt` + `tab`.

## Installation

Go Artisan are CLI tools, it can be installed as an executable file or you can build it yourself.

It is different with all rounder framework like [Goravel](https://goravel.dev) which is a web framework built with Go.

### Install latest release

This will add `goartisan` to your `~/go/bin` folder and should be ready to use anywhere.

```bash
go install github.com/wahyusa/goartisan@latest
```

### Install from executable file

You can download the latest executable file like `goartisan.exe` on releases page in this repository.

You will need to add it to your environment variable manually.

## General Usage

### Initialize new RESTful API project

```bash
goartisan init my-api-project
```

### Initialize new RESTful API project with git (recommended)

```bash
goartisan init my-api-project --git
```

### Generate a GORM Model / entity

```bash
goartisan make:model book
```

### Generate a repository

```bash
goartisan make:repo book
```

### Generate a service

```bash
goartisan make:service book
```

### Generate a HTTP handler

```bash
goartisan make:handler book
```

### Generate all model, repository, service and handler in one command

```bash
goartisan make:all book
```

### Other commands just find help

```bash
goartisan --help
```

## Default Configuration

Default configuration will automatically generated in your `~/` or `$HOME` directory.

You don't need to do anything with config file if you will just using `goartisan` by it's default configuration.

You still can adjust and modify it as you want but some directories are usually required and needed by `make:` command.

As for now it is using `.toml` format I believe it's already easy to read.

```toml
# your golang module or project name
[module]
name = "go-api"

# this is main app folder name, you may prefer to name it to "cmd" or something
[app]
folder = "app"

# you can change this to postgres but I have not tested it yet
[database]
default = "mysql"

[templates]
path = "internal/templates"

# all default directory and files generated by init command
[structure]
dirs = [
    "app/model",
    "app/service",
    "app/repository",
    "app/handler",
    "app/middleware",
    "app/config",
    "app/database",
    "app/routes",
]

files = [
    "app/config/config.go",
    "app/database/database.go",
    ".env",
    ".gitignore",
    "main.go",
    "app/routes/routes.go",
]
```

Of course because `goartisan` are built for repository pattern RESTful API if you want to replace many things or maybe using MVC pattern or you think my structures are not your best fit it is recommended to just build your own generator based on `goartisan` source code by creating your own fork version.

Consider other alternative if you want too,

[Go Blueprint](https://github.com/Melkeydev/go-blueprint)

[Go Scaffold](https://github.com/go-scaffold/go-scaffold)

## Credits

[Cobra](https://github.com/spf13/cobra)

[Viper](https://github.com/spf13/viper)

[Gin](https://github.com/gin-gonic/gin)

[GORM](https://github.com/go-gorm/gorm)

[godotenv](https://github.com/joho/godotenv)
