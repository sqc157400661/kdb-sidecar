package biz

import (
	"github.com/sqc157400661/kdb-sidecar/internal/types"
	"github.com/sqc157400661/kdb-sidecar/pkg/meta"
)

func ListInstance(req types.InstancesReq) (instances []*meta.Instance, err error) {
	err = meta.DB().Find(&instances)
	return
}
