# POM to Markdown Converter

A Go-based command-line tool that converts Maven `pom.xml` files into well-formatted Markdown documentation.

## Features

- Converts Maven POM XML files to readable Markdown format
- Extracts project information, dependencies, plugins, developers, licenses, and more
- Supports output to file or stdout
- Handles missing fields gracefully
- Clean, structured Markdown output with tables for dependencies and plugins

## Installation

1. Clone or download the project
2. Build the binary:
   ```bash
   go build -o pom-to-markdown
   ```

## Usage

### Basic Usage

```bash
# Convert pom.xml to stdout
./pom-to-markdown

# Convert specific file to stdout  
./pom-to-markdown -input my-pom.xml

# Convert to specific output file
./pom-to-markdown -input pom.xml -output README.md

# Convert using default pom.xml and save to file
./pom-to-markdown -output project-docs.md
```

### Command Line Options

- `-input <file>`: Input POM XML file path (default: "pom.xml")
- `-output <file>`: Output Markdown file path (default: stdout)
- `-help`: Show help information

## Output Format

The tool generates Markdown with the following sections (when data is available):

1. **Project Title** - Uses `<name>` or falls back to `<artifactId>`
2. **Project Information** - Group ID, Artifact ID, Version, Packaging, URL
3. **Description** - Project description from POM
4. **Parent Project** - Parent POM information
5. **Properties** - Java version, Maven compiler settings, framework versions
6. **Dependencies** - Table of all dependencies with scope and type
7. **Build Plugins** - Table of Maven plugins
8. **Developers** - List of project developers with contact info
9. **Licenses** - Project licenses with links
10. **Source Control** - SCM information and repository links

## Example

For a Spring Boot project POM, the tool will generate documentation showing:
- Project metadata and description
- Spring Boot parent configuration  
- Java version and encoding settings
- All dependencies organized in a table
- Build plugins like Spring Boot Maven plugin
- Developer contact information
- License information
- Git repository links

## Requirements

- Go 1.16 or later
- Valid Maven POM XML file as input

## Error Handling

The tool provides clear error messages for:
- Missing or unreadable input files
- Invalid XML format
- File write permissions issues