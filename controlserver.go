package main

import (
	"encoding/json"
	log "github.com/EntropyPool/entropy-logger"
	authapi "github.com/NpoolDevOps/fbc-auth-service/authapi"
	authtypes "github.com/NpoolDevOps/fbc-auth-service/types"
	fccredis "github.com/NpoolDevOps/fbc-control-service/redis"
	"github.com/NpoolDevOps/fbc-control-service/types"
	"github.com/NpoolRD/http-daemon"
	"io/ioutil"
	"net/http"
)

type ControlServerConfig struct {
	RedisCfg fccredis.RedisConfig `json:"redis"`
	Port     int                  `json:"port"`
}

type ControlServer struct {
	config      ControlServerConfig
	redisClient *fccredis.RedisCli
}

func NewControlServer(configFile string) *ControlServer {
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf(log.Fields{}, "cannot read file %v: %v", configFile, err)
		return nil
	}

	config := ControlServerConfig{}
	err = json.Unmarshal(buf, &config)
	if err != nil {
		log.Errorf(log.Fields{}, "cannot parse file %v: %v", configFile, err)
		return nil
	}

	log.Infof(log.Fields{}, "create redis cli: %v", config.RedisCfg)
	redisCli := fccredis.NewRedisCli(config.RedisCfg)
	if redisCli == nil {
		log.Errorf(log.Fields{}, "fail to create redis client")
		return nil
	}

	return &ControlServer{
		redisClient: redisCli,
	}
}

func (s *ControlServer) Run() error {
	httpdaemon.RegisterRouter(httpdaemon.HttpRouter{
		Location: types.CreateDeployAPI,
		Method:   "POST",
		Handler: func(w http.ResponseWriter, req *http.Request) (interface{}, string, int) {
			return s.CreateDeployRequest(w, req)
		},
	})

	log.Infof(log.Fields{}, "start http daemon at %v", s.config.Port)
	httpdaemon.Run(s.config.Port)
	return nil
}

func (s *ControlServer) CreateDeployRequest(w http.ResponseWriter, req *http.Request) (interface{}, string, int) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err.Error(), -1
	}

	input := types.CreateDeployInput{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, err.Error(), -2
	}

	if input.AuthCode == "" {
		return nil, "auth code is must", -3
	}

	_, err = authapi.UserInfo(authtypes.UserInfoInput{
		AuthCode: input.AuthCode,
	})
	if err != nil {
		return nil, err.Error(), -4
	}

	return nil, "", 0
}
