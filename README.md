# oidc-cli

`oidc-cli` is a command-line interface (CLI) tool designed to facilitate testing of OpenID Connect (OIDC) authentication with an identity provider. It provides essential commands for logging in, logging out, and viewing ID token data.

## Features

- Login: Authenticate with an identity provider and store the `id_token`.
- Logout: Remove the stored `id_token`.
- Token Info: Display information from the `id_token` _[WIP]_.

## Installation

Clone this repository and navigate into the project directory:

```bash
git clone https://github.com/wallanaq/oidc-cli.git
cd oidc-cli
```

Install the dependencies and build the project:

```bash
go mod tidy
go build -o bin/oidc
```

## Usage

### Commands

The `oidc-cli` provides three primary commands:

- **login**: Authenticates with the identity provider and saves the id_token.
- **logout**: Clears the saved id_token.
- **info**: Displays decoded information from the id_token.

### Examples

- **Login**:

```bash
./oidc login
```

- **Logout**:

```bash
./oidc logout
```

- **Info**:

```bash
./oidc info
```

For verbose output, add the `--verbose` or `-v` flag to any command:

```bash
./oidc login --verbose
```
