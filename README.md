
# Kai CLI
Kai is a command-line interface (CLI) tool for kickstarting a Go application project. It utilizes template files to generate a new Go project structure with a specified name.

## Features
Create a new Go project with a single command.
Template-based project structure.
Simple and intuitive CLI interface.
## Requirements
Go 1.16 or later.
Installation
To install Kai, clone the repository and build the application:

```bash
git clone https://github.com/yourusername/kai-cli.git
cd kai-cli
make build
```

You can also deploy the application to your local ~/bin directory:

```bash
Copy code
make deploy
```

## Usage
### Creating a New Project
To create a new Go project, use the kai command followed by the project name:

```bash
Copy code
./dist/kai my-new-project
```

This will create a new directory called my-new-project with the necessary files and structure.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License
This project is licensed under the MIT License.
