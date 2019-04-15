package service

import (
	"github.com/ProtocolONE/auth1.protocol.one/pkg/config"
	"github.com/ProtocolONE/auth1.protocol.one/pkg/persist"
	"github.com/ProtocolONE/auth1.protocol.one/pkg/persist/redis"
	"github.com/ProtocolONE/mfa-service/pkg/proto"
	"github.com/globalsign/mgo"
	"github.com/go-redis/redis"
	h "github.com/ory/hydra-legacy-sdk"
	"go.uber.org/zap"
)

type RegistryBase struct {
	redis   *redis.Client
	session *mgo.Session
	as      *ApplicationService
	ott     *OneTimeTokenService
	watcher persist.Watcher
	hydra   *h.CodeGenSDK
	mfa     proto.MfaService
	mailer  Mailer
}

type RegistryConfig struct {
	MgoSession  *mgo.Session
	RedisClient *redis.Client
	MfaService  proto.MfaService
	HydraConfig *config.Hydra
	Mailer      Mailer
}

func NewRegistryBase(config *RegistryConfig) InternalRegistry {
	h, err := h.NewSDK(&h.Configuration{AdminURL: config.HydraConfig.AdminURL})
	if err != nil {
		zap.L().Fatal("Hydra SDK creation failed", zap.Error(err))
	}

	return &RegistryBase{
		session: config.MgoSession,
		redis:   config.RedisClient,
		hydra:   h,
		mfa:     config.MfaService,
		mailer:  config.Mailer,
	}
}

func (r *RegistryBase) Watcher() persist.Watcher {
	if r.watcher == nil {
		r.watcher = rediswatcher.NewWatcher(r.redis)
	}

	return r.watcher
}

func (r *RegistryBase) MgoSession() *mgo.Session {
	return r.session
}

func (r *RegistryBase) HydraSDK() *h.CodeGenSDK {
	return r.hydra
}

func (r *RegistryBase) MfaService() proto.MfaService {
	return r.mfa
}

func (r *RegistryBase) Mailer() Mailer {
	return r.mailer
}

func (r *RegistryBase) ApplicationService() *ApplicationService {
	if r.as == nil {
		r.as = NewApplicationService(r)
	}

	return r.as
}

func (r *RegistryBase) OneTimeTokenService() *OneTimeTokenService {
	if r.as == nil {
		r.ott = NewOneTimeTokenService(r.redis)
	}

	return r.ott
}
