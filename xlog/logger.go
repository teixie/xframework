package xlog

import (
	"bytes"
	"fmt"
	"time"
)

type bufferLogger struct {
	buffer *bytes.Buffer
	timers map[string]time.Time
}

func (b *bufferLogger) StartTimer(name string) {
	b.timers[name] = time.Now()
}

func (b *bufferLogger) StopTimer(name string) {
	if timer, ok := b.timers[name]; ok {
		duration := float64(time.Now().Sub(timer).Nanoseconds()) / 1e6
		b.buffer.WriteString(fmt.Sprintf(" %s=%vms", name, duration))
	}
}

func (b *bufferLogger) Append(args ...interface{}) {
	b.buffer.WriteString(" " + appendString(args...))
}

func (b *bufferLogger) Appendf(format string, args ...interface{}) {
	b.buffer.WriteString(" " + appendString(args...))
}

func (b *bufferLogger) Flush() {
	Log.Infof("%s", b.buffer.String())
}

func appendString(args ...interface{}) string {
	if len(args) == 2 {
		return fmt.Sprintf("%+v=%+v", args[0], args[1])
	} else if len(args) == 1 {
		return fmt.Sprintf("%+v", args[0])
	}
	return ""
}

func NewBufferLogger(args ...interface{}) *bufferLogger {
	return &bufferLogger{
		buffer: bytes.NewBufferString(appendString(args...)),
		timers: make(map[string]time.Time),
	}
}
