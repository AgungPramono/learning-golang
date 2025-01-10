package golang_viper

import (
	viper2 "github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	var config *viper2.Viper = viper2.New()
	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
	config := viper2.New()
	config.SetConfigName("config")
	config.SetConfigFile("config.json")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "agung", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))

}
