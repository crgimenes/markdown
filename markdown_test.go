package markdown

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	h1 := "# title"

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name         string
		args         args
		wantElements []Element
		wantErr      bool
	}{
		{
			name: "h1",
			args: args{
				r: strings.NewReader(h1),
			},
			wantErr: false,
			wantElements: []Element{
				Element{
					Raw:        h1,
					Tag:        "h1",
					SubElement: []Element{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotElements, err := Parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotElements, tt.wantElements) {
				t.Errorf("Parse() = %v, want %v", gotElements, tt.wantElements)
			}
		})
	}
}
