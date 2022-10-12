package yaml

import (
	"os"
	"testing"
)

type MyData struct {
	M1 int
	M2 string
}

const BasePath = "../test"

func TestYaml(t *testing.T) {
	os.MkdirAll(BasePath, os.ModePerm)

	var d1 MyData
	configPath := BasePath + "/test.yaml"

	d1.M1 = 100
	d1.M2 = "3222"

	var err error
	if err = SaveYaml(configPath, d1); err != nil {
		t.Fatal(err)
	}

	var d2 MyData
	if err = LoadYaml(configPath, &d2); err != nil {
		t.Fatal(err)
	}

	if d1.M1 != d2.M1 {
		t.Fatal(err)
	}

	if d1.M2 != d2.M2 {
		t.Fatal(err)
	}
}

func BenchmarkYaml(b *testing.B) {
	configPath := BasePath + "/test.yaml"

	os.MkdirAll(BasePath, os.ModePerm)
	for i := 0; i < b.N; i++ {
		var d1 MyData
		d1.M1 = 100
		d1.M2 = "3222"

		var err error
		if err = SaveYaml(configPath, d1); err != nil {
			b.Fatal(err)
		}

		var d2 MyData
		if err = LoadYaml(configPath, &d2); err != nil {
			b.Fatal(err)
		}
	}
}
