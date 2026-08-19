### HTTP Utils (httputils)

The httputils package provides tools and utilities for working with HTTP. It prevents boilerplate code and encourages consistent patterns while utilizing tools already in this library (logging, ctxutils, errutils).

- **Logger**: An exported package-level logger. Replace it with your app details and it handles all httputils logging via the utils/logging package.
- **JSON**: Takes in-memory data in your server/program, marshals it into JSON bytes, and sends it with the given status code to the client via the ResponseWriter.
- **Error**: Uses the logging package to log errors and sends a consistent JSON error response to the client.
- **ErrorResponse**: The standard JSON error body (`error`, optional `code`) written by Error.
- **WithRequestID**: Checks if a request carried an X-Request-ID and assigns one if not. Stores the ID in the request context (via ctxutils) and echoes it on the response.
- **Recoverer**: Recovers the program from a panic, logs the error, and returns a safe 500 response.
- **Chain**: Takes any number of middlewares and composes them into a single http.Handler.

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/httputils/EXAMPLES.md)

---