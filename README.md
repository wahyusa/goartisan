[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# Go Artisan

Go Artisan is a CLI tool designed to simplify and accelerate the development of RESTful APIs in Go with artisan like command.

Built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)

The code generated will using Gin, GORM, MySQL and godotenv with repository pattern, but you can modify it based on each template.

## FAQ

#### Why using goartisan if I can just scaffold any initial project with `git clone` ?

Well, you will have `config.toml` that you can customize and replace `templates` to fit your own need and initialize new RESTful API project faster by using that templates.

This is also a good fit for terminal enjoyers who love to do many things directly in terminal instead of creating each file manually one by one and `alt` + `tab`.

## Installation

Go Artisan are CLI tools, it can be installed as an executable file or you can build it yourself.

It is different with all rounder framework like [Goravel](https://goravel.dev) which is a web framework built with Go.

Go Artisan is just a CLI tools to generate boilerplate code, you own the code and you can customize it as you like.

### Simple installation

This will add `goartisan` to your environment variable and should be ready to use anywhere.

```bash
go install https://github.com/wahyusa/goartisan

# ready to use

goartisan init my-api-project

goartisan make:model book
goartisan make:repo book
goartisan make:service book
goartisan make:handler book
```

## Default Configuration

Default configuration will automatically generated in your `~/` or `$HOME` directory.

You can adjust and modify it as you want but some directories are usually required and needed by `make:` command.

Of course because `goartisan` are built for repository pattern RESTful API if you want to replace many things or maybe using MVC pattern or you think my structures are not your best fit it is recommended to just build your own generator based on `goartisan` source code by creating your own fork version.
