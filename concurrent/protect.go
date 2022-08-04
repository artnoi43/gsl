package concurrent

import (
	"context"
	"fmt"
	"log"
)

// protect recovers if f panics, allowing its caller to continue execution
func Protect(f func() error) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Protect: recovered panic, reason: %s\n", fmt.Sprintf("%v", r))
		}
	}()

	return f()
}

// ProtectWithContext recovers if f panics, allowing its caller to continue execution
func ProtectWithContext(
	f func(context.Context) error,
	ctx context.Context,
) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("ProtectWithContext: recovered panic, reason: %s\n", fmt.Sprintf("%v", r))
		}
	}()

	return f(ctx)
}
