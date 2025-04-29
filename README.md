# oidc-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/wallanaq/oidc-cli)](https://goreportcard.com/report/github.com/wallanaq/oidc-cli)
[![Go Version](https://img.shields.io/github/go-mod/go-version/wallanaq/oidc-cli)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/wallanaq/oidc-cli)](https://github.com/wallanaq/oidc-cli/releases)

**oidc-cli** is a versatile command-line tool (CLI) built with Go for interacting with **OpenID Connect (OIDC) Providers**.

This tool simplifies common OIDC tasks such as fetching provider configurations, performing authentication, generating and validating tokens, making it useful for development, testing, and automation.

## ✨ Features

*   Fetch OIDC Configuration (`.well-known/openid-configuration`)
*   Perform Authorization Code Flow, including PKCE
*   Token Introspection (Check activity and metadata)
*   Verify ID Token Signatures using JWKS
*   Revoke Tokens
