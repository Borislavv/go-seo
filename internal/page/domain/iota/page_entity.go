package iota

type PageEntity int

const (
	Title       PageEntity = iota
	Description PageEntity = iota
	Canonical   PageEntity = iota
	H1          PageEntity = iota
	Content     PageEntity = iota
	FAQ         PageEntity = iota
)
