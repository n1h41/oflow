package config

import "testing"

func TestSetup(t *testing.T) {
	Setup()
	t.Log(configInstance.Db.DbSslmode)
}
