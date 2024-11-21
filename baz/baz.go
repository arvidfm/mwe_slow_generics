package baz

import "play.ground/bar"

type Baz struct{}

func (Baz) BazFunc()           {}
func (Baz) BazBarFunc(bar.Bar) {}
