# Changelog

## [1.2.0] - 2025-10-21

### ✨ Added
- New `charm update` command for automatic self-updates
- Automatically checks for latest version from GitHub releases
- Downloads and installs the appropriate binary for the current OS/architecture

### 🐛 Fixed
- Removed 20-line truncation limit on response body display
- Response bodies are now displayed completely regardless of size

## [1.1.0] - 2025-10-21

### ✨ Added
- Shorthand flags: `-b` for bearer, `-d` for data-raw, `-H` for content-type
- URL validation with clear error messages
- Builder pattern for Display struct with fluent API
- New `RequestOptions` struct for cleaner API
- `HTTPMethod` struct moved to structs package for better organization

### 🔨 Refactored
- Eliminated code duplication in cmd/root.go (reduced ~38% file size)
- Replaced multiple pointer parameters with struct-based options
- Replaced 12-parameter `SetDisplay()` with builder pattern
- Split large `DoRequest()` function into smaller, focused functions
- Improved error handling with proper error wrapping
- Added constants for magic values (DefaultContentType, BearerPrefix, BasicPrefix)

### 📚 Improved
- Better separation of concerns across modules
- More idiomatic Go code patterns
- Cleaner and more maintainable codebase
- Removed all unnecessary comments for cleaner code

### 🐛 Fixed
- Proper error propagation instead of `os.Exit(1)` calls
- Better HTTP client organization with named imports

## [1.0.0] - Initial Release

### Added
- HTTP client with support for GET, POST, PUT, PATCH, DELETE
- Beautiful colored terminal output
- Request timing and metrics
- Bearer and Basic authentication
- Custom headers support
- Cross-platform support (Linux, macOS, Windows)

