package gtime

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
)

type Duration time.Duration

func NewDuration(d time.Duration) Duration {return Duration(d)}

func NewDurationStr(s string) (Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return Duration(0), err
	}
	return Duration(d), nil
}

func MustNewDurationStr(s string) Duration {
	d, err := NewDurationStr(s)
	if err != nil {
		panic(err)
	}
	return d
}

var _ json.Marshaler = Duration(0)

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

var _ json.Unmarshaler = (*Duration)(nil)

func (d *Duration) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `"`)
	dur, err := time.ParseDuration(string(data))
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(strings.Trim(s, `"`))
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}

func (d Duration) ToDuration() time.Duration {
	return time.Duration(d)
}

func (d Duration) D() time.Duration {
	return time.Duration(d)
}

