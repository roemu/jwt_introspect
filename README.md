# Jwt Introspect

`jwt-introspect` is a command-line utility written in Go that simplifies the process of introspecting JSON Web Tokens (JWTs). 
Designed with usability in mind, this tool is perfect for web developers that always have to look up some site to decode their JWTs


## Features

- [x] **Stdin Access**: Accepts JWTs from standard input, allowing for quick and efficient processing.
- [x] **Clipboard Access**: Simplifies developer experience by enabling JWT introspection directly from the clipboard.
- [x] **File Access**: Accepts file path were JWT is stored
- [x] **Human readable timestamps**: Will print timestamps in human readable format unless `--unparsed` is set.

## Installation

### Homebrew

```bash
brew tap roemu/jwt_introspect https://github.com/roemu/jwt_introspect
brew install jwt-introspect
```

### From Source

To build `jwt-introspect` from source, ensure you have Go installed, then run:

```bash
git clone https://github.com/roemu/jwt-introspect.git
cd jwt-introspect
make

make install
```

## Usage

```bash
jwt-introspect [options] <JWT>
```

### Options

- `-h`, `--help`: Show help information.
- `--clipboard`: Introspects a JWT from the clipboard.
- `--file=<path>`: Introspects a JWT from a file.
- `--stdin`: Introspects a JWT from standard input.
- `--unparsed`: Introspects a JWT and prints raw JWT without human readable values.

### Examples

1. Introspect a JWT from standard input:

   ```bash
   echo "your.jwt" | jwt-introspect --stdin
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

For more information, check out source code :)
