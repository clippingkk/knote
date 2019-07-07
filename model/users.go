package model

type User struct {
	ID        int64     `xorm:"INT(11) NOTNULL PK AUTOINCR 'id'" json:"id"`
	Name      string    `xorm:"VARCHAR(191) NOTNULL 'name'" json:"name" db:"name"`
	Email     string    `xorm:"VARCHAR(191) NOTNULL 'email'" json:"email" db:"email"`
	Pwd       string    `xorm:"VARCHAR(191) NOTNULL 'pwd'" json:"-" db:"pwd"`
	Avatar    string    `xorm:"VARCHAR(191) 'avatar'" json:"avatar" db:"avatar"`
	Checked   bool      `xorm:"BOOL 'checked'" json:"checked" db:"checked"`
	CreatedAt int64 	`xorm:"INT(11) 'created_at'" json:"created_at" db:"created_at"`
	UpdatedAt int64 	`xorm:"INT(11) 'updated_at'" json:"updated_at" db:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}