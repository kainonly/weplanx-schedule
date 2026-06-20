package bootstrap

import (
	"os"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/kainonly/cronx/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/binding/go_playground"
	"github.com/kainonly/go/help"
	"gopkg.in/yaml.v3"
)

func LoadStaticValues(path string) (v *common.Values, err error) {
	v = new(common.Values)
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return
	}
	if err = yaml.Unmarshal(b, &v); err != nil {
		return
	}
	return
}

func UseGorm(v *common.Values) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(v.Database.Path), &gorm.Config{})
}

func UseCronx() *common.Cronx {
	return new(common.Cronx)
}

func UseHertz(v *common.Values) (h *server.Hertz, err error) {
	if v.Address == "" {
		return
	}
	vd := go_playground.NewValidator()
	vd.SetValidateTag("vd")
	opts := []config.Option{
		server.WithHostPorts(v.Address),
		server.WithCustomValidatorFunc(binding.MakeValidatorFunc(vd)),
	}

	opts = append(opts)
	h = server.Default(opts...)
	h.Use(
		help.ErrorHandler(),
	)

	return
}
