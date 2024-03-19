package connector

import "fmt"

func wrapError(err error, message string) error {
	return fmt.Errorf("panorama-connector: %s: %w", message, err)
}
