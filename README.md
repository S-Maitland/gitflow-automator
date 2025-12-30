# Git Workflow Automator (GWA)

**GWA** is a command-line tool to automate common git workflow operations, especially for teams using structured branching models (like git-flow). It streamlines routine tasks such as creating feature branches, making conventional commits, and displaying enhanced repository status via intuitive commands.

## Features

- **Automatic feature branch creation** with naming conventions
- **Interactive, conventional commit interface** for clear commit messages
- **Colorized and enhanced git status** display
- **Dry run and verbose modes** for transparency and safety

## Installation

### Using Pre-compiled Binaries (Recommended)

Download the appropriate binary for your platform from the `bin/` directory.

#### Linux

```sh
# Download and make executable
chmod +x bin/gwa-linux

# Move to a directory in your PATH (e.g., /usr/local/bin)
sudo mv bin/gwa-linux /usr/local/bin/gwa
```

#### macOS (Intel)

```sh
chmod +x bin/gwa-mac-amd64
sudo mv bin/gwa-mac-amd64 /usr/local/bin/gwa
```

#### macOS (Apple Silicon)

```sh
chmod +x bin/gwa-mac-arm64
sudo mv bin/gwa-mac-arm64 /usr/local/bin/gwa
```

#### Windows

```bat
:: Copy to a directory in your PATH (e.g., C:\Windows\System32 or add to your user PATH)

copy bin\gwa-windows.exe C:\Windows\System32\gwa.exe
```

Alternatively, you can run the binary directly from the `bin/` directory without global installation.

### Building from Source (Optional)

If you prefer to build from source, ensure you have Go installed.

```sh
make install
```

This will place the `gwa` binary in a Go bin location (such as `~/go/bin` or `/usr/local/bin`), depending on your OS.

Or, build manually:

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
| `help`           | Display help and usage information         |

### Options (for `feature`)

- `-v`, `--verbose` Show detailed output
- `-d`, `--dry-run` Preview actions without making changes
- `-h`, `--help` Show help for the command

### Examples

Create a feature branch named `add-user-login`:

```sh
gwa feature add-user-login
```

Preview the branch creation (dry-run):

```sh
gwa feature --dry-run add-user-login
```

Make a conventional commit interactively:

```sh
gwa commit
```

Show enhanced colorized status:

```sh
gwa status
```

Display CLI help:

```sh
gwa help
```

## Requirements

- Git must be installed and initialized in your working directory.
- Go is only required if building from source (not needed when using pre-compiled binaries).

## Contributing

Pull requests and suggestions are welcome! Open an issue for feedback, bugs, or feature requests.

## License

MIT License

---
