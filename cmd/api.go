package main

type application struct {
	config config
}

// run
// mount

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string // user, passwd, dbname
}
