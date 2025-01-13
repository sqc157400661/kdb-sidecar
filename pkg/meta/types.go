package meta

type Instance struct {
	PodName string `gorm:"type:varchar(100)"`
	PodIP   string `gorm:"type:varchar(100)"`
	Host    string `gorm:"type:varchar(200)"`
	Port    int
	Role    string `gorm:"type:varchar(20)"`
	Status  string `gorm:"type:varchar(50)"`
	Extra   string `gorm:"type:varchar(255)"`
}
