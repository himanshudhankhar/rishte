package envvar

import "os"

func DBUser() string {
	user, exists := os.LookupEnv("MYSQL_USER")
	if !exists {
		user = "admin_rishte"
	}
	return user
}

func DBPassword() string {
	password, exists := os.LookupEnv("MYSQL_PASSWORD")
	if !exists {
		password = "Dhankhar7924@"
	}
	return password
}

func DBName() string {
	dbname, exists := os.LookupEnv("MYSQL_DB_NAME")
	if !exists {
		dbname = "db_rishte"
	}
	return dbname
}
func DBHost() string {
	host, exists := os.LookupEnv("MYSQL_HOST")
	if !exists {
		host = "localhost"
	}
	return host
}

func DBPort() string {
	port, exists := os.LookupEnv("MYSQL_PORT")
	if !exists {
		port = "3306"
	}
	return port
}
