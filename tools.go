package toolkit

import (
	"crypto/rand"
	"fmt"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

type Tools struct {
}

func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))

		fmt.Printf("p = %d\n", p)

		x, y := p.Uint64(), uint64(len(r))

		fmt.Printf("x = %d\n", x)
		fmt.Printf("y = %d\n", y)
		fmt.Printf("y/x = %d\n", x%y)

		s[i] = r[x%y]
	}

	return string(s)
}
