package convert

import "github.com/jinzhu/copier"

func To[DstType any, SrcType any](src SrcType) (dst DstType, err error) {
	err = copier.Copy(&dst, &src)
	return
}
