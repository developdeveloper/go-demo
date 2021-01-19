package slicearg

import "strings"

type sliceArg []string

func New(p *[]string, def []string) *sliceArg {
	*p = def
	return (*sliceArg)(p)
}

func (s *sliceArg) Set(val string) error {
	*s = sliceArg(strings.Split(val, ","))
	return nil
}

//func (s *sliceArg) Get() interface{} { return []string(*s) }

func (s *sliceArg) String() string { return strings.Join(*s, ",") }
