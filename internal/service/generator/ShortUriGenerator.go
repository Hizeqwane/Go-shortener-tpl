package generator

import "strconv"

type ShortUriGenerator struct {
	increment int
}

func (generator *ShortUriGenerator) GenerateShortUri(longUri string) string {
	generator.increment++

	return strconv.Itoa(generator.increment)
}
