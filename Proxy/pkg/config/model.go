package config

type appConfig struct {
	LdapConfig     ldapConfig     `mapstructure:"ldap" json:"ldap"`
	ServerConfig   serverConfig   `mapstructure:"server" json:"server"`
	DatabaseConfig databaseConfig `mapstructure:"database" json:"database"`
}

type ldapConfig struct {
	Url           string `mapstructure:"url" json:"url"`
	Host          string `mapstructure:"host" json:"host"`
	IsSecure      bool   `mapstructure:"isSecure" json:"isSecure"`
	AdminUsername string `mapstructure:"adminUsername" json:"adminUsername"`
	AdminPassword string `mapstructure:"adminPassword" json:"adminPassword"`
}
type serverConfig struct {
	AppName     string `mapstructure:"appName" json:"appName"`
	Header      string `mapstructure:"header" json:"header"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	SecretKey   string `mapstructure:"secretKey" json:"secretKey"`
	PrivateKey  string `mapstructure:"privateKey" json:"privateKey"`
	ProxySecret string `mapstructure:"proxySecret" json:"proxySecret"`
}
type databaseConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
	Schema   string `mapstructure:"schema" json:"schema"`
	SSLMode  string `mapstructure:"sslMode" json:"sslMode"`
}
