# mdtoc - Markdown Table of Contents Generator

A simple Go CLI tool that automatically generates a table of contents (INDEX.md) from markdown files organized in subdirectories.

## Overview

`mdtoc` scans a directory structure for markdown files, extracts their titles (from the first `# Heading`), and generates a hierarchical table of contents with links to each file. Perfect for organizing documentation, blog posts, or any collection of markdown files

## Usage

```bash
# Generate INDEX.md in the current directory
mdtoc

# Generate INDEX.md in a specific directory
mdtoc /path/to/docs

# Specify custom output filename
mdtoc /path/to/docs CONTENTS.md
```

### Example

Given a directory structure like:
```
docs/
├── network/
│   ├── nmap.md
│   └── ssh.md
└── kubernetes/
    ├── helm.md
    └── kubectl.md
```

Running `mdtoc docs/` will generate `INDEX.md`:

```markdown
## Table of Content
### [kubernetes](kubernetes)
- [Helm Basics](kubernetes/helm.md)
- [kubectl Commands](kubernetes/kubectl.md)

### [network](network)
- [Nmap Network Scanning](network/nmap.md)
- [SSH Configuration](network/ssh.md)
```

## Installation

```bash
# Clone the repository
git clone https://github.com/michielvha/mdtoc.git
cd mdtoc

# Build the binary
go build -o mdtoc main.go

# Optionally, install to your PATH
go install
```

## How It Works

1. Scans the specified directory for subdirectories (ignores hidden folders starting with `.`)
2. Finds all `.md` files in each subdirectory
3. Extracts the title from the first `# Heading` in each file (falls back to filename if no heading found)
4. Generates a sorted table of contents grouped by folder
5. Writes the output to `INDEX.md` (or specified filename)

## Dependencies

- The pipeline requires to add a `PGP_PRIVATE_KEY` in the repository secrets
- The `project_name` in the [.goreleaser.yml](.goreleaser.yml) file should be set to match your repo name.


## Features

- Automated versioning with GitVersion, using tags like `0.0.0`.
- Automated container release via custom action, defaults to `ghcr.io`.
- Automated binary release using goreleaser, with signature & multi arch support.