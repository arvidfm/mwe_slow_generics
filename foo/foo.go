package foo

type Foo[A, B any] struct {
	EmbeddedFoo[A, B]
}

func (Foo[A, B]) FooFunc1()                       {}
func (Foo[A, B]) FooFunc2(FooPair[A, A, B])       {}
func (f Foo[A, B]) FooFunc3() AnyFoo[A]           { return f }
func (f Foo[A, B]) FooFunc4() FooOrEmbedded[A, B] { return f }

type EmbeddedFoo[A, B any] struct{}

func (EmbeddedFoo[A, B]) EmbeddedFunc1(A, B) {}
func (EmbeddedFoo[A, B]) EmbeddedFunc2()     {}
func (EmbeddedFoo[A, B]) EmbeddedFunc3()     {}
func (EmbeddedFoo[A, B]) EmbeddedFunc4()     {}

type FooPair[A, B, C any] struct {
	foo1 Foo[A, C]
	foo2 Foo[B, C]
}

func (FooPair[A, B, C]) PairFunc1() {}
func (FooPair[A, B, C]) PairFunc2() {}
func (FooPair[A, B, C]) PairFunc3() {}

type AnyEmbedded[A any] interface {
	EmbeddedFunc2()
	EmbeddedFunc3()
}

type AnyFoo[A any] interface {
	AnyEmbedded[A]
	FooFunc1()
	FooFunc3() AnyFoo[A]
}

type FooOrEmbedded[A, B any] interface {
	AnyEmbedded[A]
	EmbeddedFunc1(A, B)
}
