package config

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/opensourceways/community-robot-lib/utils"
	kfklib "github.com/opensourceways/kafka-lib/agent"
	redislib "github.com/opensourceways/redis-lib"

	"project/xihe-statistics/common/infrastructure/redis"
)

var reIpPort = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}:[1-9][0-9]*$`)

func LoadConfig(path string, cfg interface{}) error {
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Name     string `json:"name"`
	HttpPort int    `json:"http_port"`
	Duration int    `json:"duration"`
	PGSQL    PGSQL  `json:"pgsql"`
	Redis    Redis  `json:"redis"        required:"true"`
	MQ       MQ     `json:"mq"`
	GitLab   GitLab `json:"gitlab"`
}

type PGSQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
}

type Redis struct {
	DB redis.Config `json:"db" required:"true"`
}

type MQ struct {
	Address  string `json:"address"`
	Version  string `json:"version" required:"true"`
	MaxRetry int    `json:"max_retry"`
	Topics   `json:"topics"`
}

func (cfg *Config) GetKfkConfig() kfklib.Config {
	return kfklib.Config{
		Address: cfg.MQ.Address,
		Version: cfg.MQ.Version,
	}
}

func (cfg *Config) GetRedisConfig() redislib.Config {
	return redislib.Config{
		Address:  cfg.Redis.DB.Address,
		Password: cfg.Redis.DB.Password,
		DB:       cfg.Redis.DB.DB,
		Timeout:  cfg.Redis.DB.Timeout,
	}
}

type GitLab struct {
	RootToken    string        `json:"root_token"`
	Endponit     string        `json:"endpoint"`
	CountPerPage int           `json:"count_per_page"`
	RefreshTime  time.Duration `json:"refresh_time"`
}

type Topics struct {
	Statistics      string `json:"statistics" required:"true"`
	GitLab          string `json:"gitlab" required:"true"`
	Like            string `json:"like"             required:"true"`
	Fork            string `json:"fork"             required:"true"`
	Download        string `json:"download"         required:"true"`
	Training        string `json:"training"         required:"true"`
	Finetune        string `json:"finetune"         required:"true"`
	Following       string `json:"following"        required:"true"`
	Inference       string `json:"inference"        required:"true"`
	Evaluate        string `json:"evaluate"         required:"true"`
	Submission      string `json:"submission"       required:"true"`
	OperateLog      string `json:"operate_log"      required:"true"`
	RelatedResource string `json:"related_resource" required:"true"`
	Cloud           string `json:"cloud"            required:"true"`
	Async           string `json:"async"            required:"true"`
	BigModel        string `json:"bigmodel"         required:"true"`
}

func (cfg *MQ) Validate() error {
	if r := cfg.ParseAddress(); len(r) == 0 {
		return errors.New("invalid mq address")
	}

	return nil
}

func (cfg *MQ) ParseAddress() []string {
	v := strings.Split(cfg.Address, ",")
	r := make([]string, 0, len(v))
	for i := range v {
		if reIpPort.MatchString(v[i]) {
			r = append(r, v[i])
		}
	}

	return r
}
