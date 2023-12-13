package main

import "fmt"

type Config struct {
	value string
	name  string
}

type ConfigGetter interface {
	getConfig(name string) Config
}

type StatsigAdapter struct {
	someConfig string
}

func (statsigAdapter StatsigAdapter) getConfig(name string) Config {
	return Config{statsigAdapter.someConfig, name}
}

type LocalAdapter struct {
	someConfig      string
	someOtherConfig string
}

func (localAdapter LocalAdapter) getConfig(name string) Config {
	return Config{localAdapter.someConfig + localAdapter.someOtherConfig, name}
}

func getConfig(g ConfigGetter, name string) Config {
	return g.getConfig(name)
}

func main() {
	a1 := StatsigAdapter{"hello"}
	fmt.Println(getConfig(a1, "name1"))

	a2 := LocalAdapter{"hello", "world"}
	fmt.Println(getConfig(a2, "name2"))
}
