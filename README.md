# purekeychain

[![GoDoc](https://godoc.org/github.com/Fabianexe/purekeychain?status.svg)](https://godoc.org/github.com/Fabianexe/purekeychain)
[![Go Version](https://img.shields.io/badge/go-1.24+-00ADD8.svg)](https://golang.org/doc/go1.24)

A pure Go implementation for accessing the macOS Keychain without CGO dependencies. 
This library provides a simple API to interact with the macOS Keychain Services, 
allowing secure credential management while maintaining Go's cross-compilation benefits.

## Features

- ✨ Pure Go implementation (no CGO required)
- 🔐 Secure access to macOS Keychain
- 🛠 Cross-compilation friendly

## Why purekeychain?

Most macOS Keychain libraries in Go rely on CGO, which can complicate cross-compilation and increase build complexity.
This library uses `purego` to provide direct access to the macOS Keychain APIs without CGO,
making it easier to maintain and integrate into your projects.

## Installation
```bash
go get github.com/Fabianexe/purekeychain
```

## Requirements
- macOS 10.13 or later
- Go 1.24 or later

## Usage
The purekeychain library provides a simple interface to interact with the macOS Keychain. The library uses a service-based approach where each service can store credentials like username/password pairs.
### Creating a Service
First, create a Service object that will interact with the keychain:
```go
import "github.com/Fabianexe/purekeychain"

// Create a new service with a name
service := purekeychain.New("MyAppService")
```
### Saving Credentials
Save username/password credentials to the keychain:
```go
// Save login credentials
err := service.Save("username", "password123")
if err != nil {
    log.Fatalf("Failed to save to keychain: %v", err)
}
```
### Retrieving Credentials
Retrieve the credentials for a service:
```go
// Retrieve credentials
username, password, err := service.Load()
if err != nil {
    log.Fatalf("Failed to load from keychain: %v", err)
}
fmt.Printf("Username: %s, Password: %s\n", username, password)
```

### Updating Credentials
Update existing credentials in the keychain:
```go
// Update credentials
err := service.Update("username", "newPassword456")
if err != nil {
    log.Fatalf("Failed to update keychain: %v", err)
}
```
### Deleting Credentials
Delete credentials associated with a service:
```go
// Delete credentials
err := service.Delete()
if err != nil {
    log.Fatalf("Failed to delete from keychain: %v", err)
}
```
### Error Handling
The library provides detailed error messages directly from the macOS Security framework.
Always check for errors after keychain operations, especially in production environments.
```go
err := service.Save("username", "password123")
if err != nil {
    // Handle errors appropriately
    // The error messages come directly from the Security framework
    fmt.Fprintf(os.Stderr, "Keychain error: %v\n", err)
}
```

### Complete Example
```go
package main

import (
    "fmt"
    "log"

    "github.com/Fabianexe/purekeychain"
)

func main() {
    // Create a new service
    service := purekeychain.New("MyApp")

    // Save credentials
    err := service.Save("johndoe", "securePassword123")
    if err != nil {
        log.Fatalf("Failed to save: %v", err)
    }
    
    // Retrieve credentials
    username, password, err := service.Load()
    if err != nil {
        log.Fatalf("Failed to load: %v", err)
    }
    fmt.Printf("Retrieved credentials: %s / %s\n", username, password)
    
    // Update password
    err = service.Update(username, "newPassword456")
    if err != nil {
        log.Fatalf("Failed to update: %v", err)
    }
    
    // Verify update
	username, newPassword, err := service.Load()
	if err != nil {
		log.Fatalf("Failed to load: %v", err)
	}
    fmt.Printf("Updated password: %s / %s\n", username, newPassword)
    
    // Delete the entry when no longer needed
    err = service.Delete()
    if err != nil {
        log.Fatalf("Failed to delete: %v", err)
    }
    
    fmt.Println("Keychain operations completed successfully")
}
```

## Security Considerations

- This library handles sensitive information. Always follow security best practices.
- Keychain items are protected by macOS's security mechanisms.
- The library uses the user's default login keychain.

## Limitations

- macOS only
- Some advanced Keychain features are not implemented

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
For major changes, please open an issue first to discuss what you would like to change.

## License

MIT License

Copyright (c) 2025 Fabianexe

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
