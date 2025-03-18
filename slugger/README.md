### Slugger

The `slugger` package provides a simple and efficient way to generate URL-friendly slugs from strings. Below are the main features and methods of the package:

#### **Slugger Constructor**

- **`New(substitutions map[string]string, withEmoji, unique bool) *Slugger`**:  
  Creates a new `Slugger` instance.
  - **`substitutions`**: A map of string replacements to apply before generating the slug.
  - **`withEmoji`**: If true, emojis will be included in a slug-friendly format.
  - **`unique`**: If true, slugger will append a UUID to the end of the slug.

#### **Slugger Methods**

- **`Slug(s, separator string) string`**:  
    Generates a slugified version of the input string `s`. If `separator` is provided, it will be used to separate words in the slug; otherwise, a default separator(`-`) is applied.

#### **Notes**

- If a `substitutions` map is provided, it will replace all occurrences of the specified keys with their corresponding values. For example, given a substitution pair `{"the": ""}` and the input string `over there`, the resulting slug will be `over-re`.

## Examples:

For examples of each function, please check out [EXAMPLES.md](/slugger/EXAMPLES.md)

---
