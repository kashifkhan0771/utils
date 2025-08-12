### Slugger

The `slugger` package provides a simple and efficient way to generate URL-friendly slugs from strings. Below are the main features and methods of the package:

#### **Slugger Constructor**

- **`New(substitutions map[string]string, withEmoji bool) *Slugger`**:  
  Creates a new `Slugger` instance.
  - **`substitutions`**: A map of string replacements to apply before generating the slug.
  - **`withEmoji`**: If true, emojis will be included in a slug-friendly format.

#### **Slugger Methods**

- **`Slug(s, separator string) string`**:  
    Generates a slugified version of the input string `s`. If `separator` is provided, it will be used to separate words in the slug; otherwise, a default separator(`-`) is applied.


- **`AddSubstitution(key, value string)`**:  
    Adds a new substitution pair to the `Slugger` instance. This allows you to define custom replacements that will be applied when generating slugs.


- **`RemoveSubstitution(key string)`**:  
    Removes a substitution pair from the `Slugger` instance. This is useful for cleaning up or modifying existing substitutions.


- **`ReplaceSubstitution(key, newValue string)`**:  
    Replaces an existing substitution value with a new value. This allows you to update the slug generation rules without having to recreate the `Slugger` instance.


- **`SetSubstitutions(substitutions map[string]string)`**:  
    Sets the entire map of substitutions for the `Slugger` instance. This can be used to initialize or update multiple substitutions at once.

  
#### **Notes**

- If a `substitutions` map is provided, it will replace all occurrences of the specified keys with their corresponding values. For example, given a substitution pair `{"the": ""}` and the input string `over there`, the resulting slug will be `over-re`.

## Examples:

For examples of each function, please check out [EXAMPLES.md](/slugger/EXAMPLES.md)

---
