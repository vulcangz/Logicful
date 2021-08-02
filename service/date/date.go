package date

import (
	"time"
)

func ISO8601(date time.Time) string {
	return date.Format(time.RFC3339)
}
