package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Datetime time.Time

type HoraireOptions struct {
	WithDay bool
}

func (d Datetime) Sub(d2 Datetime) time.Duration {
	t := time.Time(d)
	t2 := time.Time(d2)
	return t.Sub(t2)

}
func (d Datetime) Horaire(options *HoraireOptions) string {
	t := time.Time(d)
	//fmt.Println(t.Unix())
	if t.Unix() == 0 || t.Unix() == -62135596800 {
		return ""
	}
	if options.WithDay {
		return time.Time(d).Format("[Monday 02] 15:04")
	} else {
		return time.Time(d).Format("15:04")
	}
}
func ParseDate(d string) (*Datetime, error) {
	if d != "" {
		array := strings.Split(d, ".")
		year, _ := strconv.Atoi(array[2])
		month, _ := strconv.Atoi(array[1])
		day, _ := strconv.Atoi(array[0])
		dd := Datetime(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
		return &dd, nil
	} else {
		return nil, errors.New("no value")
	}
}
func (d Datetime) MarshalJSON() ([]byte, error) {
	// Step 1. Format the time as a Go string.
	t := time.Time(d)
	formatted := t.Format("2006-01-02")

	// Step 2. Convert our formatted time to a JSON string.
	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}
func (d *Datetime) UnmarshalJSON(b []byte) error {
	if string(b) != "null" {
		// 0. Check that string is encapsulated between """
		if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
			return errors.New("not a json string")
		}

		// 1. Strip the double quotes from the JSON string.
		b = b[1 : len(b)-1]

		// 2. Parse the result using our desired format.
		t, err := time.Parse("2006-01-02T15:04:05-0700", string(b))
		if err != nil {
			return fmt.Errorf("failed to parse time: %w", err)
		}

		// finally, assign t to *d
		*d = Datetime(t)
	}
	return nil
}
