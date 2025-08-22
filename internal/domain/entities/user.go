package entities

import ( 
	"time" 
	"golang.org/x/crypto/bcrypt" 
	"github.com/google/uuid" 
	"gorm.io/gorm" 
)

type User struct{
	ID 					string 				`gorm:"type:uuid;primaryKey"`
	Email 				string 				`gorm:"unique;not null"`
	Password 			string 				`gorm:"not null"`
	Role 				string 				`gorm:"not null;check:role IN ('user','organizer','admin')"`
	CreatedAt 			time.Time 
	UpdatedAt 			time.Time 
	DeletedAt 			gorm.DeletedAt 		`gorm:"index"`
}


func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	hashed, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}