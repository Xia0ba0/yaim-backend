package config

const (
	DBDriver     = "mysql"
	DBUser       = "root"
	DBPassword   = "1005"
	DBDatabase   = "test"
	DBConnection = DBUser + ":" + DBPassword + "@/" + DBDatabase + "?charset=utf8"
)
