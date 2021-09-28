# emommit

[![codecov](https://codecov.io/gh/chyroc/emommit/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/emommit)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/emommit "go report card")](https://goreportcard.com/report/github.com/chyroc/emommit)
[![test status](https://github.com/chyroc/emommit/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/emommit/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/emommit)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Femommit.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Femommit)

![](demo.gif)

## Install

### By Brew

```shell
brew install chyroc/tap/emommit
```

### By Go

```shell
go install github.com/chyroc/emommit@latest
```

## Usage

### Simple to use

Just run `emommit`:

```shell
emommit
```

This is equivalent to running `git commit -m "<generate message>"`.

`emommit` will start an interactive program to let you choose emoji(such as 🐛) and git commit type(such as fix).

### Customized git commit command

If you want to customize the git commit command, you can add `--` after the `emommit` command, and then add any commands supported by git commit.

Such as auto stage all modified files and commit with a message:

```shell
emommit -- -a
```