package config

type Configuration struct {
	App         *AppConfiguration
	RestClient  *RestClientConfiguration
	Log         *LogConfiguration
	Permissions *PermissionsConfiguration
	Database    *DatabaseConfiguration
}

type PermissionsConfiguration struct {
	ModelFile  string
	PolicyFile string
}

type AppConfiguration struct {
	Name        string
	Group       string
	Env         string
	ProdMode    bool
	Port        int64
	BodyLimit   int
	ContextPath string
	Jwt         *JwtConfigInfo
}

type RestClientConfiguration struct {
	Debug            bool
	EnableTrace      bool
	TimeoutInSeconds int
}

type JwtConfigInfo struct {
	Issuer                   string
	SecretKey                string
	ExpirationMinutes        int
	RefreshExpirationMinutes int
}

type DatabaseConfiguration struct {
	Host           string
	Port           string
	DBName         string
	User           string
	Password       string
	ConnectTimeout int64
	EnableLog      bool
	AutoMigrate    bool
}
type LogConfiguration struct {
	FilePath string
	Level    string
	Format   string
}
