package convert

import "github.com/jinzhu/copier"

// Return a new variable copied from source
func To[DstType any, SrcType any](src SrcType) (dst DstType, err error) {
	err = copier.Copy(&dst, &src)
	return
}

func Update[Model any, Schema any](model *Model, schema Schema) error {
	return copier.Copy(model, schema)
}
