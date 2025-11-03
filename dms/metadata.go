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
	Date
)

type Metadata struct {
	Value        interface{}
	PropertyType PropertyInfo
}

type PropertyInfo struct {
	Name         string
	propertyType MetadataType
}

func (m Metadata) Type() MetadataType {
	return m.PropertyType.propertyType
}
