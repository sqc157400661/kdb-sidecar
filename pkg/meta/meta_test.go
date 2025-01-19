package meta

import (
	"fmt"
	"log"
	"testing"
)

func TestMetaDB(t *testing.T) {
	fmt.Println(manager.Setup())
	//DB().

	var err error

	// 创建一个新的 Instance 实例
	//instance := Instance{
	//	ID:         "1",
	//	ClusterID:  "cluster1",
	//	MasterID:   "master1",
	//	ServerUUID: "uuid1",
	//	ServerID:   1,
	//	PodName:    "pod1",
	//	PodIP:      "192.168.1.1",
	//	Host:       "host1",
	//	Port:       8080,
	//	Role:       "worker",
	//	Version:    "v1.0",
	//	ReadOnly:   false,
	//	Status:     "active",
	//	Extra:      "extra info",
	//}
	//
	//// 插入数据
	//_, err = DB().Insert(&instance)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 查询数据
	var result Instance
	has, err := DB().Where("id =?", "1").Get(&result)
	if err != nil {
		log.Fatal(err)
	}
	if has {
		fmt.Printf("Found instance: %+v\n", result)
	} else {
		fmt.Println("Instance not found")
	}
}
