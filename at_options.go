package dingtalk

type AtOption interface {
	Apply(model *AtModel)
}

type funcAtOption struct {
	f func(model *AtModel)
}

func (o *funcAtOption) Apply(m *AtModel) {
	o.f(m)
}

func newFuncAtOption(f func(model *AtModel)) *funcAtOption {
	return &funcAtOption{f: f}
}

func WithAtAll(isAtAll bool) AtOption {
	return newFuncAtOption(func(o *AtModel) {
		o.IsAtAll = isAtAll
	})
}

func WithAtMobiles(atMobiles []string) AtOption {
	return newFuncAtOption(func(o *AtModel) {
		o.AtMobiles = atMobiles
	})
}

func WithAtUserIds(atUserIds []string) AtOption {
	return newFuncAtOption(func(o *AtModel) {
		o.AtUserIds = atUserIds
	})
}
