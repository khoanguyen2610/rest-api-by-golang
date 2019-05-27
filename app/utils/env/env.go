package env

import (
	"context"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"user-service/configs"
)

type Env struct {
	conf 	*configs.AppConfig
	db   	*gorm.DB
	ctx  	context.Context
	v		*validator.Validate
}

func NewEnv(conf *configs.AppConfig, db *gorm.DB, ctx context.Context, v *validator.Validate) *Env {

	if ctx == nil {
		ctx = context.Background()
	}

	return &Env{
		conf: conf,
		db:  db,
		ctx:  ctx,
		v: v,
	}
}

func (e *Env) GetAppConf() *configs.AppConfig {
	return e.conf
}

func (e *Env) GetDB() *gorm.DB {
	return e.db
}

func (e *Env) GetValidator() *validator.Validate {
	return e.v
}

// context.Context interface
func (e *Env) Deadline() (deadline time.Time, ok bool) {
	return e.ctx.Deadline()
}

func (e *Env) Done() <-chan struct{} {
	return e.ctx.Done()
}

func (e *Env) Err() error {
	return e.ctx.Err()
}

func (e *Env) Value(key interface{}) interface{} {
	return e.ctx.Value(key)
}

// Env changers
func WithValue(parent *Env, key, val interface{}) *Env {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	env := *parent
	env.ctx = context.WithValue(env.ctx, key, val)
	return &env
}

func WithContext(parent *Env, ctx context.Context) *Env {
	env := *parent
	env.ctx = ctx
	return &env
}

func MustWithTx(parent *Env) *Env {
	env := *parent
	env.db = mustGetTx(parent.db)
	return &env
}

func mustGetTx(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx
}
