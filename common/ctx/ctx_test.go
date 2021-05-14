package ctx

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := NewContextComponentName(context.Background(), "Shivang")
	fmt.Println(ComponentNameFromContext(ctx))
}
