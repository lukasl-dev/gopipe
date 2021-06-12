/*
 * MIT License
 *
 * Copyright (c) 2021 lukas.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package gopipe

import (
	"errors"
	"io"
	"os"
)

var (
	// ErrNotAvailable occurs whenever no shell pipe
	// was used.
	ErrNotAvailable = errors.New("pipe is not available")
)

// Available returns true whenever a shell pipe was
// used.
func Available() (bool, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return info.Mode()&os.ModeNamedPipe != 0, nil
}

// Read returns the content of the shell pipe if one
// was used.
// If this is not the case, ErrNotAvailable is returned.
func Read() ([]byte, error) {
	ok, err := Available()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNotAvailable
	}
	return io.ReadAll(os.Stdin)
}
