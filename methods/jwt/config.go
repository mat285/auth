package jwt

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"os"

	"github.com/mat285/auth/pkg/crypto"
)

type Config struct {
	Key []byte

	PrivateKeyFile string             `json:"privateKeyFile" yaml:"privateKeyFile"`
	PublicKeyFile  string             `json:"publicKeyFile" yaml:"publicKeyFile"`
	PrivateKey     ed25519.PrivateKey `json:"-" yaml:"-"`
	PublicKey      ed25519.PublicKey  `json:"-" yaml:"-"`

	CAFile string            `json:"caFile" yaml:"caFile"`
	CACert *x509.Certificate `json:"-" yaml:"-"`
}

type ServiceConfig struct {
	ServiceIdentity
	Config
}

func (c *Config) Resolve(ctx context.Context) error {
	_, err := c.GetPrivateKey()
	if err != nil {
		return err
	}
	_, err = c.GetPublicKey()
	if err != nil {
		return err
	}
	// TODO parse CAcert
	return nil
}

func (c *Config) GetCACert() (*x509.Certificate, error) {
	if c.CACert != nil {
		return c.CACert, nil
	}
	bytes, err := os.ReadFile(c.CAFile)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, err
	}
	c.CACert = cert
	return c.CACert, nil
}

func (c *Config) GetPrivateKey() (ed25519.PrivateKey, error) {
	if len(c.PrivateKey) != 0 {
		return c.PrivateKey, nil
	}
	if len(c.PrivateKeyFile) == 0 {
		return nil, nil
	}
	bytes, err := os.ReadFile(c.PrivateKeyFile)
	if err != nil {
		return nil, err
	}
	key, err := crypto.ParsePemEd25519PrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	c.PrivateKey = key
	return c.PrivateKey, nil
}

func (c *Config) GetPublicKey() (ed25519.PublicKey, error) {
	if len(c.PublicKey) != 0 {
		return c.PublicKey, nil
	}
	if len(c.PublicKeyFile) == 0 {
		return nil, nil
	}
	bytes, err := os.ReadFile(c.PublicKeyFile)
	if err != nil {
		return nil, err
	}
	key, err := crypto.ParsePemEd25519PublicKey(bytes)
	if err != nil {
		return nil, err
	}
	c.PublicKey = key
	return c.PublicKey, nil
}

func NewServiceConfig(service string, cfg Config) ServiceConfig {
	return ServiceConfig{
		ServiceIdentity: ServiceIdentity{Name: service},
		Config:          cfg,
	}
}

func DevAuthConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{
		PrivateKeyFile: "_dev_certs/jwt_ed25519",
		PublicKeyFile:  "_dev_certs/jwt_ed25519.pub",
	}
	return cfg, cfg.Resolve(ctx)
}

func DevAuthVerifyConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{
		PublicKeyFile: "_dev_certs/jwt_ed25519.pub",
	}
	return cfg, cfg.Resolve(ctx)
}
