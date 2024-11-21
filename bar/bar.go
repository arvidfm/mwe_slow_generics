package bar

import "play.ground/foo"

type Bar struct{}

func (Bar) BarFunc()                     {}
func (Bar) BarFooFunc(foo.Foo[Bar, int]) {}
