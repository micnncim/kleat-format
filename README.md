# kleat-format

[![actions-workflow-test][actions-workflow-test-badge]][actions-workflow-test]
[![release][release-badge]][release]
[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![license][license-badge]][license]

`kleat-format` is a CLI to format `halconfig` for [Kleat](https://github.com/spinnaker/kleat).

## Installation

Precompiled binaries are available in [GitHub Releases](https://github.com/micnncim/kleat-format/releases).

### Via Go

```
$ go get github.com/micnncim/kleat-format/cmd/kleat-format
```

## Usage

```console
$ kleat-format --help
Usage:
  kleat-format /path/to/halconfig [flags]

Flags:
      --check     If true, only check whether there is diff between source halconfig and formatted one
  -h, --help      help for kleat-format
  -q, --quiet     If true, suppress printing logs
      --version   If true, print version information
  -w, --write     If true, write result to source halconfig instead of stdout

```

<!-- badge links -->

[actions-workflow-test]: https://github.com/micnncim/kleat-format/actions?query=workflow%3ATest
[actions-workflow-test-badge]: https://img.shields.io/github/workflow/status/micnncim/kleat-format/Test?label=Test&style=for-the-badge&logo=github

[release]: https://github.com/micnncim/kleat-format/releases
[release-badge]: https://img.shields.io/github/v/release/micnncim/kleat-format?style=for-the-badge&logo=github

[pkg.go.dev]: https://pkg.go.dev/github.com/micnncim/kleat-format?tab=overview
[pkg.go.dev-badge]: http://bit.ly/pkg-go-dev-badge

[license]: LICENSE
[license-badge]: https://img.shields.io/github/license/micnncim/kleat-format?style=for-the-badge
