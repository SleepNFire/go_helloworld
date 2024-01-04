package units_test

import (
	"os"
	"testing"

	"github.com/SleepNFire/go_bootstrap/project/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name          string
		environnement [][]string
		want          config.Config
	}{
		{
			name: "default",
			want: config.Default(),
		},
		{
			name: "custom_conf",
			environnement: [][]string{
				{"PROJECT_MYSQL_ADRESS", "custom_localhost"},
				{"PROJECT_MYSQL_ROOT_PASSWORD", "custom_root"},
				{"PROJECT_MYSQL_DATABASE", "custom_project"},
				{"PROJECT_MYSQL_USER", "custom_user"},
				{"PROJECT_MYSQL_PASSWORD", "custom_password"},
			},
			want: config.Config{
				MySql: config.MySql{
					Adress:   "custom_localhost",
					User:     "custom_user",
					Password: "custom_password",
					Database: "custom_project",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if len(tt.environnement) > 0 {
				for _, env := range tt.environnement {
					os.Setenv(env[0], env[1])
				}
			}

			conf, err := config.Init()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, *conf)

			os.Clearenv()
		})
	}
}
