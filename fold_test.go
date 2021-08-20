package fold_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	
	"github.com/0xc0d/go-fold"
)

func TestFoldReader(t *testing.T) {
	tests := []struct {
		name      string
		input     io.Reader
		width     int
		want      []byte
		wantErr   bool
		wantPanic bool
	}{
		{
			name:    "big line",
			input:   strings.NewReader(strings.Repeat("12345", 3)),
			width:   5,
			want:    []byte("12345\n12345\n12345"),
			wantErr: false,
		},
		{
			name: "big lines",
			input: strings.NewReader(
				strings.Repeat("12345", 2) +
					"1\n" + strings.Repeat("12345", 2)),
			width:   4,
			want:    []byte("1234\n5123\n451\n1234\n5123\n45"),
			wantErr: false,
		},
		{
			name:    "small fold",
			input:   strings.NewReader(strings.Repeat("12\n", 6)),
			width:   1,
			want:    []byte("1\n2\n1\n2\n1\n2\n1\n2\n1\n2\n1\n2\n"),
			wantErr: false,
		},
		{
			name:      "EOF while having hamper",
			input:     strings.NewReader(strings.Repeat("0", 16)),
			width:     5,
			want:      []byte("00000\n00000\n00000\n0"),
			wantErr:   false,
			wantPanic: false,
		},
		{
			name:      "many new lines",
			input:     strings.NewReader(strings.Repeat("\n", 10)),
			width:     5,
			want:      []byte("\n\n\n\n\n\n\n\n\n\n"),
			wantErr:   false,
			wantPanic: false,
		},
		{
			name:      "panic width 0",
			input:     strings.NewReader("whatever"),
			width:     0,
			want:      nil,
			wantErr:   false,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//defer func() {
			//	if r := recover(); (r != nil) != tt.wantPanic {
			//		t.Errorf("Read() paic = %v, wantErr %v", r, tt.wantPanic)
			//	}
			//}()

			fr := fold.NewReader(tt.input, tt.width)
			got, err := readAll(fr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if bytes.Compare(got, tt.want) != 0 {
				t.Errorf("Read() got = \n%s, want \n%s", got, tt.want)
			}
		})
	}
}

// readAll is an identical function with io.ReadAll but only difference is
// buffer size is smaller to easily be able to check the reader output.
func readAll(r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 16)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}
