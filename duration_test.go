package gtime

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Run("NewDuration", func(t *testing.T) {
		d := NewDuration(time.Second)
		if d.ToDuration() != time.Second {
			t.Fatalf("want 1s, but got %s", d.ToDuration().String())
		}

		d, err := NewDurationStr("1m")
		if err != nil {
			t.Fatalf("failed to NewDurationStr: %v", err)
		}
		if d.ToDuration() != time.Minute {
			t.Fatalf("want 1m, but got %s", d.ToDuration().String())
		}

		d, err = NewDurationStr("invalid")
		if err == nil {
			t.Fatalf("invalid duration string, but got %s", d.ToDuration().String())
		}

		func(){
			defer func() {
				if e := recover(); e != nil {
					t.Fatalf("unexpected panic err: %v", e)
				}
			}()
			MustNewDurationStr("1s")
		}()

		func(){
			defer func() {
				if e := recover(); e == nil {
					t.Fatal("expect panic")
				}
			}()
			MustNewDurationStr("1d")
		}()
	})

	t.Run("json", func(t *testing.T) {
		type A struct {
			Timeout Duration `json:"timeout"`
		}

		jsonText := `{"timeout":"3s"}`
		a := A{}
		err := json.Unmarshal([]byte(jsonText), &a)
		if err != nil {
			t.Fatalf("failed to parse json text: %v", err)
		}
		if a.Timeout.ToDuration() != 3*time.Second {
			t.Fatalf("want 3s, but got %s", a.Timeout.ToDuration().String())
		}

		jsonData, err := json.Marshal(a)
		if err != nil {
			t.Fatalf("failed to build json: %v", err)
		}
		if string(jsonData) != jsonText {
			t.Fatalf("different json text: %q", string(jsonData))
		}

		jsonText = `{"timeout":"ss"}`
		err = json.Unmarshal([]byte(jsonText), &a)
		if err == nil {
			t.Fatal("expect error")
		}
	})
}
