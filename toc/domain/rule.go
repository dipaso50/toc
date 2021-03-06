package domain

import "regexp"

type Rule struct {
	RegexpFilter      *regexp.Regexp
	DestinationFolder string
}

func NewRule(fileMatchRegularExpresion, destination string) Rule {
	rr := Rule{DestinationFolder: destination}
	rr.RegexpFilter = regexp.MustCompile(fileMatchRegularExpresion)

	return rr
}
