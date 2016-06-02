package lib

type Type interface {
	Hello() string
	World()
}

type T struct {
	data Type
}

func New(data Type) *T {
	return &T{data: data}
}

func Exec(t *T) {
	t.data.Hello()
	t.data.World()
}
