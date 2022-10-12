package yaml

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 加载yaml配置
func LoadYaml(configPath string, v any) (err error) {
	bs, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(bs, v); err != nil {
		return
	}

	return
}

// 保存yaml配置
func SaveYaml(configPath string, v any) (err error) {
	fp, err := os.Create(configPath)
	if err != nil {
		return
	}
	defer fp.Close()

	encoder := yaml.NewEncoder(fp)
	err = encoder.Encode(v)

	return
}
