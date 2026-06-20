package bootstrap

import (
	"database/sql"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/binding/go_playground"
	"github.com/hertz-contrib/cors"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
	"github.com/kainonly/go/passport"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"resty.dev/v3"
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

func UseGorm(v *common.Values) (orm *gorm.DB, err error) {
	if orm, err = gorm.Open(sqlite.Open(``), &gorm.Config{}); err != nil {
		return
	}
	var db *sql.DB
	if db, err = orm.DB(); err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return
}

func UseVictorialogs(v *common.Values) *common.Victorialogs {
	return &common.Victorialogs{
		Client: resty.New().
			SetBaseURL(v.Database.Victorialogs).
			SetHeader("Content-Type", "application/stream+json"),
	}
}

func UsePassport(v *common.Values) *passport.Passport {
	return passport.New(
		passport.SetKey(v.Key),
		passport.SetIssuer(v.Node),
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
		cors.New(cors.Config{
			AllowOrigins: v.Origins,
			AllowMethods: []string{"GET", "POST"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
				"Authorization", "X-Requested-With",
			},
			MaxAge: 12 * time.Hour,
		}),
	)
	return
}
