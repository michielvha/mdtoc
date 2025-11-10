# Template Go Application

This repository servers as a starting point for each new Go application.

## Overview

Provide an overview of the project

## Dependencies

- The pipeline requires to add a `PGP_PRIVATE_KEY` in the repository secrets
- The `project_name` in the [.goreleaser.yml](.goreleaser.yml) file should be set to match your repo name.


## Features

- Automated versioning with GitVersion, using tags like `0.0.0`.
- Automated container release via custom action, defaults to `ghcr.io`.
- Automated binary release using goreleaser, with signature & multi arch support.