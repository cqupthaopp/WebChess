package MySQLUtil

type User struct {
	Username string `gorm:"primary_key"`
	Password string
}
