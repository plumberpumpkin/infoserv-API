package dms

type MetadataType int

// Extend if necessary
const (
	String MetadataType = iota
	Integer
	Long
	Float
	Blob
	Boolean
)

type Metadata struct {
	Value        interface{}
	PropertyType PropertyType
}

type PropertyType struct {
	Name string
	Type MetadataType
}
