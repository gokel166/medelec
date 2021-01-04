package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
  dbUser string
  dbPswd string
  dbHost string
  dbPort string
  dbName string
  testDbHost string
  testDbName string
}

func Get() *Config {
  conf := &Config{}

  flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
  flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("POSTGRES_PASS"), "DB user pass")
  flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "DB port")
  flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB host")
  flag.StringVar(&conf.dbName, "dbhost", os.Getenv("POSTGRES_HOST"), "DB name")
  flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")

  flag.Parse()

  return conf
}

func (c *Config) GetDbConnectionString() string {
  return c.GetDbConnectionString(c.dbHost, c.dbName)
}

func (c *Config) GetTestDbConnectionString() string {
  return c.GetDbConnectionString(c.testDbHost, c.testDbName)
}

func (c *Config) GetDbConnectionString(dbhost string, dbname string) string {
  return fmt.Sprint(
    "postgres://%s:%s@%s:%s/%s?sslmode=disable",
    c.dbUser,
    c.dbPswd,
    dbhost,
    c.dbPort,
    dbname
  )
}
