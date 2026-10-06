package enums

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type normalizationCase struct {
	ID      string      `json:"id"`
	Enum    string      `json:"enum"`
	Value   string      `json:"value"`
	Pairs   []EnumValue `json:"pairs"`
	Dropped bool        `json:"dropped"`
	Error   string      `json:"error"`
}

func TestSharedNormalizationCases(t *testing.T) {
	payload, err := os.ReadFile("enum-normalization.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture struct {
		Cases []normalizationCase `json:"cases"`
	}
	if err := json.Unmarshal(payload, &fixture); err != nil {
		t.Fatal(err)
	}
	for _, test := range fixture.Cases {
		t.Run(test.ID, func(t *testing.T) {
			result, err := Normalize(test.Enum, test.Value)
			if test.Error != "" {
				if err == nil {
					t.Fatalf("unknown enum %q returned %+v", test.Enum, result)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			pairs := make([]EnumValue, result.Len())
			for i := range pairs {
				pairs[i] = result.At(i)
			}
			if !reflect.DeepEqual(pairs, test.Pairs) || result.Dropped != test.Dropped {
				t.Fatalf("got %+v dropped=%v; want %+v dropped=%v", pairs, result.Dropped, test.Pairs, test.Dropped)
			}
		})
	}
}

func TestCallerCannotPoisonReadViews(t *testing.T) {
	result, err := Normalize("colour", "navy")
	if err != nil {
		t.Fatal(err)
	}
	pair := result.At(0)
	pair.Value = "poison"
	result.Dropped = true
	again, err := Normalize("colour", "navy")
	if err != nil {
		t.Fatal(err)
	}
	if again.At(0) != (EnumValue{Enum: "colour", Value: "blue"}) || again.Dropped {
		t.Fatalf("caller mutation changed registry result: %+v", again)
	}
	drop, err := Normalize("colour", "teal")
	if err != nil {
		t.Fatal(err)
	}
	drop.Dropped = false
	again, err = Normalize("colour", "teal")
	if err != nil {
		t.Fatal(err)
	}
	if !again.Dropped || again.At(0) != (EnumValue{Enum: "colour", Value: "teal"}) {
		t.Fatalf("caller mutation changed drop: %+v", again)
	}
}

func TestReadViewsAllocateNothing(t *testing.T) {
	for _, value := range []string{"blue", "navy", "maroon", "teal", "unknown"} {
		if allocations := testing.AllocsPerRun(100, func() {
			result, err := Normalize("colour", value)
			if err != nil || result.At(0).Enum == "" {
				panic("invalid read view")
			}
		}); allocations != 0 {
			t.Fatalf("Normalize(%q) allocated %v times", value, allocations)
		}
	}
}

func TestReadViewIndexBoundaries(t *testing.T) {
	for _, value := range []string{"blue", "navy", "teal"} {
		result, err := Normalize("colour", value)
		if err != nil {
			t.Fatal(err)
		}
		for _, index := range []int{-1, result.Len()} {
			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("At(%d) accepted for %q with length %d", index, value, result.Len())
					}
				}()
				result.At(index)
			}()
		}
	}
}
