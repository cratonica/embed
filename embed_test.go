package embed

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestPack(t *testing.T) {
	input := make(map[string][]byte)
	input["Foo"] = []byte{0, 1, 2, 3}
	packed, err := Pack(input)
	if err != nil {
		t.Fatal(err)
	}
	unpacked, err := Unpack(packed)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 4; i++ {
		if input["Foo"][i] != unpacked["Foo"][i] {
			t.Fatal("Data mismatch")
		}
	}
}

func TestPackLarge(t *testing.T) {
	input := make(map[string][]byte)
	data := make([]byte, 1024*1024)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}
	input["Foo"] = data
	packed, err := Pack(input)
	if err != nil {
		t.Fatal(err)
	}
	unpacked, err := Unpack(packed)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(data); i++ {
		if data[i] != unpacked["Foo"][i] {
			t.Fatal("Data mismatch")
		}
	}
}

func TestCreateFromFiles(t *testing.T) {
	workDir, err := ioutil.TempDir("", "datamap_test_")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(workDir)
	}()
	os.Mkdir(filepath.Join(workDir, "SubDir"), 0777)
	testMap := map[string][]byte{
		"File0": []byte{0, 1, 2, 3},
		"SubDir" + string(os.PathSeparator) + "File1": []byte{4, 5, 6, 7},
	}
	for k, v := range testMap {
		err = ioutil.WriteFile(filepath.Join(workDir, k), v, 0777)
		if err != nil {
			t.Fatal(err)
		}
	}
	resultMap, err := CreateFromFiles(workDir)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range testMap {
		for i := 0; i < len(v); i++ {
			if resultMap[k][i] != v[i] {
				t.Fatal("Data mismatch")
			}
		}
	}
}

func TestGenerateGoCode(t *testing.T) {
	code := GenerateGoCode("mypackage", "Resources", []byte{1, 2, 3})
	_, err := parser.ParseFile(token.NewFileSet(), "", code, 0)
	t.Log(code)
	if err != nil {
		t.Fatal(err)
	}
}
