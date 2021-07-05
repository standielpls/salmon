package ctx_values

import (
	"context"
	"fmt"
)

func Get(ctx context.Context, key string) string {
	return fmt.Sprintf("%v", ctx.Value(key))
}
