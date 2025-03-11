### Boolean Functions

- **IsTrue**: Checks if the provided string represents a true value (e.g., "T", "1", "TRUE").
- **Toggle**: Toggles the given boolean value, returning its negation. (e.g., `true` becomes `false` and vice versa).
- **AllTrue**: Checks if all the values in a slice of booleans are `true`. Returns `false` if the slice is empty.
- **AnyTrue**: Checks if at least one value in a slice of booleans is `true`. Returns `false` if the slice is empty.
- **NoneTrue**: Checks if none of the values in a slice of booleans are `true`. Returns `true` if the slice is empty.
- **CountTrue**: Counts the number of `true` values in a slice of booleans. Returns `0` for an empty slice.
- **CountFalse**: Counts the number of `false` values in a slice of booleans. Returns `0` for an empty slice.
- **Equal**: Checks if all the values in a variadic boolean argument are equal. Returns `true` if the slice contains only one or no elements.
- **And**: Performs a logical AND operation on all the values in a slice of booleans. Returns `true` only if all values are `true`. Returns `false` for an empty slice.
- **Or**: Performs a logical OR operation on all the values in a slice of booleans. Returns `true` if at least one value is `true`. Returns `false` for an empty slice.

## Examples:

For examples of each function, please checkout [EXAMPLES.md](/boolean/EXAMPLES.md)

---
