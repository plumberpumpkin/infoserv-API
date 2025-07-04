package dms

type Document struct {
	Type       string
	Identifier string
	Properties map[string]Metadata
	Contents   [][]byte
}
