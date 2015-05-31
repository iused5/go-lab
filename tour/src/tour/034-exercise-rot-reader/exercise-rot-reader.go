package main 

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader)Read(p []byte) (n int, err error) {
	n, err = reader.r.Read(p)
	if err != nil {
		return n, err 
	}

	/*	
	for _, b := range p {
		if (b >= 'A' && b < 'N') || (b >= 'a' && b < 'n') {
			b += 13
		} else
		if (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z') {
			b -= 13
		}
	}
	*/
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13
		} else
		if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

