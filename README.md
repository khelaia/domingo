# Domingo - EPP Client in Go

Domingo is a Golang-based Extensible Provisioning Protocol (EPP) client built to communicate with domain registry systems. This client is initially designed to work with VeriSign's EPP API but is structured to support additional registries in the future.

## Overview

The Domingo EPP client provides a modular and extensible way to perform domain-related operations using EPP. It currently includes core EPP functionalities and will gradually incorporate additional commands and registry compatibility.

## Features

- **VeriSign EPP Compatibility**: Currently configured for VeriSign; other registries will be added in future releases.
- **Modular Design**: Organized to allow easy expansion and customization.
- **Environment-Driven Configuration**: Sensitive information such as credentials, host, and port is configurable via environment variables.
- **Extensible Command Support**: Commands are added incrementally, with support for the most common EPP commands initially.

## Getting Started

### Prerequisites

- Go 1.23.2 or later
- A valid EPP account and credentials for the target registry (e.g., VeriSign)
- SSL certificate files for secure connections with the registry server

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/khelaia/domingo.git
cd domingo
```

### Environment Configuration
Set the following environment variables to configure the client:


* UserID: Your EPP User ID
* Password: Your EPP password
* hostname: EPP server hostname (e.g., epp.verisign.com)
* port: EPP server port
* clientCertFile: Path to your client certificate file
* clientKeyFile: Path to your client private key file


## Usage
The Domingo client currently provides basic EPP commands, such as logging in and checking domain availability. As development progresses, additional commands will be available.

### Example: Running the Client
To run the EPP client, use the main entry point:

```bash 
go run main.go
```

## Project Structure
The project follows a modular structure with dedicated folders for configuration, internal methods, and tests:
```plaintext
domingo/
├── cmd/                 # Entry point for command-line usage
├── pkg/domingo/config/  # Configuration files and helpers
├── pkg/domingo/internal # Internal helpers and logic
├── pkg/domingo/methods/ # EPP command methods (e.g., login.go, checkDomain.go)
├── tests/               # Test cases
└── main.go              # Main application entry
```

## Roadmap

1. **Core Commands**: Implement basic EPP commands (e.g., login, domain check, registration).
2. **Additional Registries**: Extend compatibility with other registries beyond VeriSign.
3. **Expanded Command Set**: Incrementally add support for all EPP commands.

## Contributing

Contributions are welcome! \
To get started:

1. **Fork** the repository.
2. Create a **feature branch**:
   ```bash
   git checkout -b feature/new-feature
   ```
3. **Commit** your changes
    ```bash 
    git commit -m 'Add new feature'
    ```
4. **Push** to the branch:
    ```bash 
    git push origin feature/new-feature
    ```
5. Open a **Pull Request**.


## License
This project is licensed under the MIT License.
