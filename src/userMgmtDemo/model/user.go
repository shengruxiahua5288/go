package model

type Users struct {
	Id           int    `gorm:"column:Id;primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"column:Name"`
	Password     string `gorm:"column:password"`
    Status        int   `gorm:"column:status"`

}
