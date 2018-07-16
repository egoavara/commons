package colors

type Predefine interface {
	Find(name string) *U32
	Keys() []string
	ToString(c U32) string
	Type() Expression
}

