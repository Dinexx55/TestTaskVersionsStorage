package config

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/spf13/viper"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type HTTPServerConfig struct {
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
	TimeOutSec        int
	Port              string
	Host              string
}

type Configurator struct {
}

func NewConfiguration() (*Configurator, error) {

	viper.SetConfigType("json")

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read conf file: %w", err)
	}

	c := &Configurator{}

	return c, nil
}

func (cfg *Configurator) GetHTTPSrvConfig() *HTTPServerConfig {

	return &HTTPServerConfig{
		ReadTimeout:       viper.GetDuration("srv.readTimeout"),
		WriteTimeout:      viper.GetDuration("srv.writeTimeout"),
		ReadHeaderTimeout: viper.GetDuration("srv.readHeaderTimeout"),
		TimeOutSec:        viper.GetInt("srv.timeOutSec"),
		Port:              viper.GetString("srv.port"),
		Host:              viper.GetString("srv.host"),
	}
}

func (cfg *Configurator) GetRabbitMQConfig() *RabbitMQConfig {
	return &RabbitMQConfig{
		Password: viper.GetString("rabbit.password"),
		Username: viper.GetString("rabbit.username"),
		Port:     viper.GetString("rabbit.port"),
		Host:     viper.GetString("rabbit.host"),
	}
}

func (cfg *Configurator) GetAMQPConnectionURL(rabbitCfg *RabbitMQConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitCfg.Username, rabbitCfg.Password, rabbitCfg.Host, rabbitCfg.Port)
}

type AppEnvironment string

const (
	Release             AppEnvironment = "release"
	Development         AppEnvironment = "development"
	DefaultEnv          AppEnvironment = Development
	EnvironmentVariable                = "APP_ENV"
)

func (cfg *Configurator) GetEnvironment(logger *zap.Logger) AppEnvironment {
	logger.With(
		zap.String("place", "GetEnvironment"),
	).Info("Reading GetEnvironment")

	env := os.Getenv(EnvironmentVariable)
	if env == "" {
		env = string(DefaultEnv)
	}

	logger.Info("Running in " + env)
	return AppEnvironment(env)
}

type AuthProviderConfig struct {
	Host         string
	Port         int
	Timeout      time.Duration
	Retry        int
	TimeoutRetry time.Duration
}

func (cfg *Configurator) GetAuthProviderConfig(logger *zap.Logger) *AuthProviderConfig {
	logger.With(
		zap.String("place", "GetAuthProviderConfig"),
	).Info("Reading AuthProviderConfig config from file")

	provider := &AuthProviderConfig{
		Host:         viper.GetString("auth.host"),
		Port:         viper.GetInt("auth.port"),
		Timeout:      viper.GetDuration("auth.timeout"),
		Retry:        viper.GetInt("auth.retry"),
		TimeoutRetry: viper.GetDuration("auth.timeoutRetry"),
	}
	return provider
}
