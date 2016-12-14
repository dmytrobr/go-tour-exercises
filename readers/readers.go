//https://tour.golang.org/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(p []byte) (int, error) {
	byteCOunter := 0
	for i := range p {
		p[i] = 'A'
		byteCOunter++
	}

	return byteCOunter, nil
}

func main() {
	reader.Validate(MyReader{})
}
