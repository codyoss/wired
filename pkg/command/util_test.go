package command

import (
	"reflect"
	"testing"
)

func TestParseID(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		wantID int
		wantOK bool
	}{
		{"works", []string{"1"}, 1, true},
		{"unable to parse", []string{"abc"}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, gotOK := parseID(tt.input)
			if gotID != tt.wantID {
				t.Errorf("got %v, want %v", gotID, tt.wantID)
			}
			if gotOK != tt.wantOK {
				t.Errorf("got %v, want %v", gotOK, tt.wantOK)
			}
		})
	}

}

func TestExtractIDs(t *testing.T) {
	input := []string{"http://someapi/resource/1/"}
	want := []int{1}
	got := extractIDs(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCollectNames(t *testing.T) {
	ids := []int{1}
	fn := func(id int) string { return "Cody" }
	want := []string{"Cody"}
	got := collectNames(ids, fn)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
