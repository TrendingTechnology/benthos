// Copyright (c) 2018 Ashley Jeffs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package message

import (
	"time"

	"github.com/Jeffail/benthos/lib/types"
)

//------------------------------------------------------------------------------

// Lock wraps a message into a read only type that restricts access to only a
// single message part, accessible either by index 0 or -1.
func Lock(msg types.Message, part int) types.Message {
	return &lockedMessage{
		m:    msg,
		part: part,
	}
}

//------------------------------------------------------------------------------

// lockedMessage wraps a message in a read only restricted type.
type lockedMessage struct {
	m    types.Message
	part int
}

// Copy simply returns the same type, it's read only and therefore doesn't need
// copying.
func (m *lockedMessage) Copy() types.Message {
	return m
}

// DeepCopy simply returns the same type, it's read only and therefore doesn't
// need copying.
func (m *lockedMessage) DeepCopy() types.Message {
	return m
}

func (m *lockedMessage) Get(index int) types.Part {
	if index != 0 && index != -1 {
		return nil
	}
	return m.m.Get(m.part)
}

func (m *lockedMessage) SetAll(p []types.Part) {
}

func (m *lockedMessage) Append(b ...types.Part) int {
	return 0
}

func (m *lockedMessage) Len() int {
	if m.m.Len() == 0 {
		return 0
	}
	return 1
}

func (m *lockedMessage) Iter(f func(i int, b types.Part) error) error {
	return f(0, m.m.Get(m.part))
}

func (m *lockedMessage) LazyCondition(label string, cond types.Condition) bool {
	return cond.Check(m)
}

func (m *lockedMessage) CreatedAt() time.Time {
	return m.m.CreatedAt()
}

//------------------------------------------------------------------------------
