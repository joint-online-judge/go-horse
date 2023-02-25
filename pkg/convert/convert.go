package convert

import "github.com/jinzhu/copier"

func To[DstType any](src any) (dst DstType, err error) {
	err = copier.Copy(&src, &dst)
	return
}
