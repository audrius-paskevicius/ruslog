package ruslog

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/audrius-paskevicius/logrus"
)

const defaultTimestampFormat = time.RFC3339

const (
	// formatter types
	FORMATTER_SIMPLE = "Simple"
	FORMATTER_JSON   = "JSON"
	FORMATTER_TEXT   = "Text"
)

type (
	Formatter struct {
		Name      string
		Formatter logrus.Formatter
	}

	SimpleFormatter struct {
		// TimestampFormat sets the format used for marshaling timestamps.
		TimestampFormat string
	}
)

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	fmt.Fprintf(b, "[%s]", entry.Time.Format(timestampFormat))
	fmt.Fprintf(b, " [%s]", strings.ToUpper(entry.Level.String()))
	fmt.Fprintf(b, " %s", entry.Message)

	for key, value := range entry.Data {
		if key != "time" && key != "level" && key != "msg" {
			fmt.Fprintf(b, " %s=%s", key, value)
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
