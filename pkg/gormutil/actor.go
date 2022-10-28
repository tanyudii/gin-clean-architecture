package gormutil

type HaveCreator interface {
	SetCreator(s string)
}

type HaveEditor interface {
	SetEditor(s string)
}
