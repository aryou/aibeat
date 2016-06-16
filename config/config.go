// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Aibeat AibeatConfig
}

type AibeatConfig struct {
	Period string `config:"period"`
	Test string `config:"test"`
}
