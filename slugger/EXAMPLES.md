## Slugger Examples

### Basic slug generation

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{}, false)
	fmt.Println(s.Slug("Wôrķšpáçè ~~sèťtïñğš~~", ""))
}

```

#### Output:

```
workspace-settings
```

### Generate slugs with different separators

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{}, false)
	// Will use the default separator
	fmt.Println(s.Slug("Wôrķšpáçè ~~sèťtïñğš~~", ""))
	// Will use the custom separator
	fmt.Println(s.Slug("Wôrķšpáçè ~~sèťtïñğš~~", "/"))
}

```

#### Output:

```
workspace-settings
workspace/settings
```

### Generate slugs with substitutions

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{"%": "percent", "€": "euro"}, false)
	fmt.Println(s.Slug("10% or 5€", ""))
}

```

#### Output:

```
10-percent-or-5-euro
```

### Generate slugs with emoji replacement

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{}, true)
	fmt.Println(s.Slug("a 😺, 🐈‍⬛, and a 🦁 go to 🏞️", ""))
}

```

#### Output:

```
a-grinning-cat-black-cat-and-a-lion-go-to-national-park
```

### Remove stop words
```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/slugger"
)

func main() {
	s := slugger.New(map[string]string{"and": "", "the": "", "of": ""}, false)
	fmt.Println(s.Slug("The Beauty and the Power of Nature", ""))
}

```

#### Output:

```
beauty-power-nature
```