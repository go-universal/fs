# Flexible File System

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/fs?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/fs.svg)](https://pkg.go.dev/github.com/go-universal/fs)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/fs/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/fs)](https://goreportcard.com/report/github.com/go-universal/fs)
![Contributors](https://img.shields.io/github/contributors/go-universal/fs)
![Issues](https://img.shields.io/github/issues/go-universal/fs)

The `fs` package provides a flexible file system abstraction for working with both local and embedded file systems. It offers various utilities for file operations such as checking existence, reading files, searching, and more.

## Installation

To use the `fs` package, add it to your Go project:

```bash
go get github.com/go-universal/fs
```

## Features

- Support for both local and embedded file systems.
- File existence checks.
- File reading and opening.
- Searching for files by patterns or content.
- Lookup for multiple files matching a pattern.
- Integration with `fs.FS` and `http.FileSystem`.

## Usage

Creates a `FlexibleFS` instance backed by the local file system at the specified directory path.

```go
fs := fs.NewDir(".")
```

Creates a `FlexibleFS` instance backed by the provided embedded file system.

```go
//go:embed *.go
var embeds embed.FS

fs := fs.NewEmbed(embeds)
```

## Exists

Checks if a file with the given name exists in the file system.

```go
exists, err := fs.Exists("file.go")
if err != nil {
    log.Fatal(err)
}
fmt.Println("File exists:", exists)
```

## Open

Opens a file with the given name and returns an `fs.File` interface for reading the file.

```go
file, err := fs.Open("file.go")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

## ReadFile

Reads the entire content of the file with the given name and returns the content as a byte slice.

```go
content, err := fs.ReadFile("file.go")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(content))
```

## Search

Searches for a phrase in files within the specified directory, optionally ignoring certain files and filtering by extension.

**Example:**

```go
result, err := fs.Search(".", "main", "", "go")
if err != nil {
    log.Fatal(err)
}
if result != nil {
    fmt.Println("Found:", *result)
} else {
    fmt.Println("No match found")
}
```

## Find

Searches for a file matching the given regex pattern in the specified directory.

```go
result, err := fs.Find(".", ".*\\.go")
if err != nil {
    log.Fatal(err)
}
if result != nil {
    fmt.Println("Found:", *result)
} else {
    fmt.Println("No match found")
}
```

## Lookup

Searches for files matching the given regex pattern in the specified directory and returns a slice of matching file paths.

```go
results, err := fs.Lookup(".", ".*\\.go")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Found files:", results)
```

## FS

Returns the underlying `fs.FS` interface of the file system.

```go
fsInterface := fs.FS()
```

## Http

Returns the `http.FileSystem` instance of the file system.

```go
httpFS := fs.Http()
http.Handle("/", http.FileServer(httpFS))
log.Fatal(http.ListenAndServe(":8080", nil))
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
