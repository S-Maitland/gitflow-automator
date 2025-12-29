# Git Workflow Automator (GWA)

**GWA** is a command-line tool that helps automate common git workflow operations, especially for teams following structured branching models (like git-flow). It assists with routine tasks such as creating feature branches, making conventional commits, and displaying enhanced repository status with easy-to-use commands.

## Features

- **Create feature branches** with automatic naming and prefixing
- **Interactive, conventional commit interface** for clear commit messages
- **Colorized, enhanced git status** for visibility
- **Dry run and verbose modes** for transparency and safety

## Installation

Build and install with Go (requires Go installed):

```sh
make install
```

This will place the `gwa` binary into your `~/go/bin`, `/usr/local/bin`, or corresponding location based on your OS.

Alternatively, build manually:

```sh
go build -o gwa cmd/gwa/main.go
```

## Usage

```sh
gwa <command> [options] [arguments]
```

### Commands

| Command          | Description                                |
| ---------------- | ------------------------------------------ |
| `feature <name>` | Create a new feature branch                |
| `commit`         | Create a conventional (interactive) commit |
| `status`         | Show enhanced git status                   |
| `help`           | Show help and usage information            |

### Options (for `feature`)

- `-v`, `--verbose` Show detailed output
- `-d`, `--dry-run` Preview actions without making changes
- `-h`, `--help` Show help for the command

### Examples

Create a feature branch named `add-user-login`:

```sh
gwa feature add-user-login
```

Preview the creation without actually doing it:

```sh
gwa feature --dry-run add-user-login
```

Make a conventional commit interactively:

```sh
gwa commit
```

Show status with enhanced colors:

```sh
gwa status
```

Display CLI help:

```sh
gwa help
```

## Requirements

- Git must be installed and initialized in your working directory.
- Go is required for building from source.

## Contributing

Pull requests and suggestions are welcome! Open an issue for feedback, bugs, or feature requests.

## License

MIT License

---
