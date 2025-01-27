//go:build windows
// +build windows

package tea

import (
	"context"
	"io"

	localereader "github.com/mattn/go-localereader"
)

func readInputs(ctx context.Context, msgs chan<- Msg, input io.Reader) error {
	return readAnsiInputs(ctx, msgs, localereader.NewReader(input))
}
