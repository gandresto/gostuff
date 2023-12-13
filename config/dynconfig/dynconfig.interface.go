package dynconfig

type DynamicConfigServiceI interface {
	GetConfig() map[string]string
}
