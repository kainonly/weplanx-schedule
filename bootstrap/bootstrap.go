package bootstrap

import (
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/binding/go_playground"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
	"github.com/kainonly/go/passport"
	"gopkg.in/yaml.v3"

	badger "github.com/dgraph-io/badger/v4"
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

func UseBadger(v *common.Values) (db *badger.DB, err error) {
	if db, err = badger.Open(badger.DefaultOptions(v.Database.Path)); err != nil {
		panic(err)
	}
	return
}

func UsePassport(v *common.Values) *passport.Passport {
	return passport.New(
		passport.SetKey(v.Key),
		passport.SetIssuer(v.Domain),
	)
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
