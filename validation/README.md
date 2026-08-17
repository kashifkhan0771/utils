### Validation (validation)

- **StringNotEmpty**: Checks that a string value is not empty.
- **StringIn**: Checks that a string value is one of the allowed values.
- **IntBetween**: Checks that an integer value is within a min and max range.
- **IsEmail**: Checks that a string value is a valid email address.
- **IsURL**: Checks that a string value is a valid URL, optionally restricted to allowed schemes.
- **IsSlug**: Checks that a string value is a valid slug.
- **Validator**: Interface for types that want to validate multiple values.
- **ValidateStruct**: Validates a struct that implements the Validator interface.
- **Errors**: Aggregates multiple validation errors into a single error.
- **Add**: Appends a validation error to the Errors aggregator.
- **AsError**: Returns the aggregated errors as a single error (nil if none).

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/validation/EXAMPLES.md)

---