# Jwt Introspect

`jwt-introspect` is a command-line utility written in Go that simplifies the process of introspecting JSON Web Tokens (JWTs). Designed with usability in mind, this tool is perfect for developers looking to integrate JWT introspection into their pipelines seamlessly.

## Features

- [ ] **Pipeline Friendly**: Easily integrate `jwt-introspect` into your existing workflows and automation scripts.
- [x] **Stdin Access**: Accepts JWTs from standard input, allowing for quick and efficient processing.
- [x] **Clipboard Access**: Simplifies developer experience by enabling JWT introspection directly from the clipboard.
- [x] **File Access**: Accepts file path were jwt token is stored
- [ ] **Nix Package/Flake**: Provides a Nix package and flake for easy installation and management, ensuring a smooth setup process.

## Installation

### From Source

To build `jwt-introspect` from source, ensure you have Go installed, then run:

```bash
git clone https://github.com/roemu/jwt-introspect.git
cd jwt-introspect
make
```

### From nixpkgs

To install `jwt-introspect` using the nix package manager, run the following.

```bash
nix shell nixpkgs#jwt-introspect
```

Note that experimental features 'flakes' and 'commands' must be enabled.

## Usage

```bash
jwt-introspect [options] <jwt>
```

### Options

- `-h`, `--help`: Show help information.
- `-c`, `--clipboard`: Introspects a JWT from the clipboard.
- `-s`, `--stdin`: Introspects a JWT from standard input.

### Examples

1. Introspect a JWT from standard input:

   ```bash
   echo "your.jwt.token" | jwt-introspect --stdin
   ```

2. Introspect a JWT from the clipboard:

   ```bash
   jwt-introspect --clipboard
   ```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests on the [GitHub repository](https://github.com/roemu/jwt-introspect).

## License

This project is licensed under the GNU General Public License. See the [LICENSE](./LICENSE) file for details.

---

For more information, check out the [documentation](https://github.com/roemu/jwt-introspect/docs) or reach out to the community for support. Happy introspecting!
