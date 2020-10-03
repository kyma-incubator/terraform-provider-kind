package kind

import "github.com/pelletier/go-toml"

func normalizeToml(tomlString interface{}) (string, error) {
	if tomlString == nil || tomlString.(string) == "" {
		return "", nil
	}

	s := tomlString.(string)
	tree, err := toml.Load(s)
	if err != nil {
		return s, err
	}
	return tree.ToTomlString()
}
