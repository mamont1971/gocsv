package gocsv

//Wraps around SafeCSVWriter and makes it thread safe.
import (
	"sync"
)

type SafeCSVWriter struct {
	Writer CSVWriter
	m sync.Mutex
}

type CSVWriter interface {
	Error() error
	Flush()
	Write(record []string) error
}

func NewSafeCSVWriter(original CSVWriter) *SafeCSVWriter {
	return &SafeCSVWriter{
		Writer: original,
	}
}

//Override write
func (w *SafeCSVWriter) Write(row []string) error {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(row)
}

//Override flush
func (w *SafeCSVWriter) Flush() {
	w.m.Lock()
	w.Writer.Flush()
	w.m.Unlock()
}
