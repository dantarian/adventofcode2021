package image

import "fmt"

type Algorithm []bool

func NewAlgorithm(data string) (Algorithm, error) {
	if len(data) != 512 {
		return nil, fmt.Errorf("incorrect algorithm length; expected 512, got %v", len(data))
	}

	algorithm := [512]bool{}
	for i, char := range data {
		algorithm[i] = char == '#'
	}
	return algorithm[:], nil
}
