package users

type Credential struct {
	UserID uint `gorm:"primaryKey;autoIncrement:false"`
	User User
	Password string
}
