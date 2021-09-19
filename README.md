# Overview

[![Go Reference][badge-svg]][badge-link]
[![Build and Tests][build-svg]][build-link]
[![Code Coverage][codecov-svg]][codecov-link]

`go-doze` is a wrapper around the [resty][1] golang library to more closely
focus its API for use in writing small and quick REST wrappers to eventually be
used in terraform provider plugins.

It also comes with a small wrapper test suite to make testing the REST APIs you
write easier.

## Installation

Simply call

```console
$ go get occult.work/doze
```

from your shell of choice, and go modules will take care of the rest.

## Development

`go-doze` uses [`task`][2] to run the most common operations. These tasks are
then duplicated with extra flags within the GitHub Actions workflow.

Task has [instructions on installation](https://taskfile.dev/#/installation).
Once installed, simply run `task` from the project directory.

[1]: https://github.com/go-rest/resty
[2]: https://github.com/go-task/task

[codecov-svg]: https://codecov.io/gh/slurps-mad-rips/go-doze/branch/main/graph/badge.svg
[build-svg]: https://github.com/slurps-mad-rips/go-doze/actions/workflows/build.yml/badge.svg
[badge-svg]: https://pkg.go.dev/badge/occult.work/doze.svg

[codecov-link]: https://codecov.io/gh/slurps-mad-rips/go-doze
[build-link]: https://github.com/slurps-mad-rips/go-doze/actions/workflows/build.yml
[badge-link]: https://pkg.go.dev/occult.work/doze
