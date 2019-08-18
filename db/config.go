package db

const (
	// MaxOpenConn 最大连接数
	MaxOpenConn = 512
	// MaxIdleConn 最大空闲连接
	MaxIdleConn = 10
	// MaxConnLifeTime 链接最大有效时间
	MaxConnLifeTime = 300 // 单位秒
)

// Config 数据库配置
type Config struct {
	LogMode         bool
	Driver          string
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifeTime int
}

// InitWithDefaults 初始化数据库配置
func (c *Config) InitWithDefaults() {
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = MaxOpenConn
	}

	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = MaxIdleConn
	}

	if c.MaxConnLifeTime <= 0 {
		c.MaxConnLifeTime = MaxConnLifeTime
	}
}
