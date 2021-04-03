package feed

import "github.com/mmcdole/gofeed"

// Article is an alias for the gofeed.Item. This is just a glue to
// not pollute the code with the gofeed specific impls
type Article = gofeed.Item
