package init

import (
	"fmt"
	"strings"

	"authwithtoken/lib/pkg/utils"

	"github.com/spf13/viper"
)

func setupMainConfig() {

	if utils.IsFileOrDirExist("config.yml") {
		viper.SetConfigFile("config.yml")
		err := viper.MergeInConfig()
		if err != nil {
			fmt.Println("error config not found")
		}
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
}
