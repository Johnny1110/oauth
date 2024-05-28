package enum

type RedisKey struct {
	KeyPrefix  string
	KeySuffix  string
	DefaultTTL int64
}

var (
	ACCESS_TOKEN_RDS_KEY  = RedisKey{KeyPrefix: "frizo:oauth:access_token", DefaultTTL: 3600}
	REFRESH_TOKEN_RDS_KEY = RedisKey{KeyPrefix: "frizo:oauth:refresh_token", DefaultTTL: 3600}
)
