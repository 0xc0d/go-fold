package fold

import (
	"bytes"
	"io"
	"strings"
)

type foldReader struct {
	r       io.Reader
	width   int
	lineLen int
	//hamper stores remaining bytes from last read which had been delayed du to
	//newline insertion.
	hamper []byte
}

func NewReader(r io.Reader, width int) io.Reader {
	if width == 0 {
		panic("width must be non-zero")
	}

	return &foldReader{
		r:     r,
		width: width,
	}
}

func NewReaderBytes(b []byte, width int) io.Reader {
	return NewReader(bytes.NewReader(b), width)
}

func NewReaderString(s string, width int) io.Reader {
	return NewReader(strings.NewReader(s), width)
}

func (f *foldReader) Read(p []byte) (n int, err error) {
	hn := copy(p, f.hamper)
	f.hamper = f.hamper[hn:]
	n, err = f.r.Read(p[hn:])
	n += hn

	switch err {
	case nil:
	case io.EOF:
		if n == 0 {
			return
		}
	default:
		return
	}

	var s int
	for bi := 0; bi < n; {
		bi = bytes.IndexByte(p[s:n], '\n')
		e := bi
		if e == -1 {
			e = n - s
		}

		i := f.width - f.lineLen
		for i < e {
			trim := make([]byte, len(p[s+i:n]))
			copy(trim, p[s+i:n])
			p[s+i] = '\n'
			nn := copy(p[s+i+1:], trim)
			n += 1 - len(trim[nn:])
			f.hamper = append(trim[nn:], f.hamper...)
			f.lineLen = 0

			s += i + 1
			e = e - i - len(trim[nn:])
			i = f.width
		}

		s += e + 1
		f.lineLen = e
		if bi == -1 {
			break
		}

		f.lineLen = 0
	}

	return n, err
}
