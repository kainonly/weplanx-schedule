package bootstrap

import (
	"os"
	"regexp"

	"github.com/kainonly/cronx/common"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/go-playground/validator/v10"
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

func UseCronx() *common.Cronx {
	return new(common.Cronx)
}

func UseHertz(v *common.Values) (h *server.Hertz, err error) {
	if v.Address == "" {
		return
	}
	vd := go_playground.NewValidator()
	vd.SetValidateTag("vd")
	vdx := vd.Engine().(*validator.Validate)
	vdx.RegisterValidation("snake", func(fl validator.FieldLevel) bool {
		matched, errX := regexp.MatchString("^[a-z_]+$", fl.Field().Interface().(string))
		if errX != nil {
			return false
		}
		return matched
	})
	vdx.RegisterValidation("sort", func(fl validator.FieldLevel) bool {
		matched, errX := regexp.MatchString("^[a-z_.]+:(-1|1)$", fl.Field().Interface().(string))
		if errX != nil {
			return false
		}
		return matched
	})

	opts := []config.Option{
		server.WithHostPorts(v.Address),
		server.WithCustomValidator(vd),
	}

	opts = append(opts)
	h = server.Default(opts...)
	h.Use(
		help.ErrorHandler(),
	)

	return
}
