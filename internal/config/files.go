package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile               = configCile("ca.pem")
	ServerCertFile       = configCile("server.pem")
	ServerKeyFile        = configCile("server-key.pem")
	RootClientCertFile   = configCile("root-client.pem")
	RootClientKeyFile    = configCile("root-client-key.pem")
	NobodyClientCertFile = configCile("nobody-client.pem")
	NobodyClientKeyFile  = configCile("nobody-client-key.pem")
	ACLModelFile         = configCile("model.conf")
	ACLPolicyFile        = configCile("policy.csv")
)

func configCile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".proglog", filename)
}
