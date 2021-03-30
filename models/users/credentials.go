package users

type Credential struct {
	UserID          uint `gorm:"primaryKey;autoIncrement:false"`
	User            User
	Password        string `form:"password" validate:"required,strongPassword,eqfield=ConfirmPassword"`
	ConfirmPassword string `form:"confirmPassword" gorm:"-" validate:"required"`
}
