package dbfactory

import (
	"os"
	"testing"

	"github.com/sanservices/kit/database"
	"{{ cookiecutter.module_name }}/goutils/settings"
)

const dbName = "dbutils"

func TestMain(m *testing.M) {
	code := m.Run()

	os.Remove(dbName + ".db")
	os.Exit(code)
}

func TestNew(t *testing.T) {

	cases := []struct {
		Name          string
		DBSettings    database.DatabaseConfig
		ExpectedError error
	}{
		{
			Name: "Should create a sqlite connection",
			DBSettings: database.DatabaseConfig{
				Engine: "sqlite",
				Name:   dbName,
			},
			ExpectedError: nil,
		},

		{
			Name: "Should fail with unsuported database engine",
			DBSettings: database.DatabaseConfig{
				Engine: "random-engine",
				Name:   dbName,
			},
			ExpectedError: ErrInvalidDBEngine,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			config := &settings.Settings{DB: c.DBSettings}

			db, err := New(config)
			if err != c.ExpectedError {
				t.Errorf("Expected error: %v instead got: %v", c.ExpectedError, err)
			}

			if db != nil {
				db.Close()
			}
		})
	}

}
