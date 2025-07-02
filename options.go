package dingtalk

type Options struct {
	//
	AccessToken string
	//
	secret string
	//
	AtMobiles []string
	//
	IsAtAll bool
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	return newOptions(opts...)
}

func newOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithAccessToken(accessToken string) Option {
	return func(o *Options) {
		o.AccessToken = accessToken
	}
}

func WithSecret(secret string) Option {
	return func(o *Options) {
		o.secret = secret
	}
}
