### File System Utilities

- **FormatFileSize**: Formats a file size given in bytes into a human-readable string with appropriate units (B, KB, MB, GB, TB).

- **FindFiles**: Searches for files with the specified extension in the given root directory and returns a slice of matching file paths.

- **GetDirectorySize**: Calculates the total size (in bytes) of all files within the specified directory.

- **FilesIdentical**: Compares two files byte by byte to determine if they are identical.

- **DirsIdentical**: Compares two directories to determine if they are identical.

- **GetFileMetadata**: Retrieves metadata for a specified file path. Returns a `FileMetadata` struct that can be marshaled to JSON.

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/fsutils/EXAMPLES.md)

---
