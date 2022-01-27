package cmd

import (
	"bytes"
	"os"
	"template-renderer/test"
	"testing"
)

const yamlData1 = `
a: 1
b:
  c: 2
`

const yamlData2 = `
c: 3
`

const jsonData = `{"a": 1, "b": {"c": 2}}`

const TempPath = "../tmp/"

func TestSimple(t *testing.T) {
	defer os.RemoveAll(TempPath)

	var buffer bytes.Buffer
	rootCmd.SetOut(&buffer)
	rootCmd.SetArgs([]string{"-i", "../test/data/templates", "-o", testDir(t), "-d", yamlData1})

	err := rootCmd.Execute()

	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "", buffer.String())

	file, err := os.ReadFile(testDir(t) + "/test.txt")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "A=1\nB=2", string(file))
}

func TestExtension(t *testing.T) {
	defer os.RemoveAll(TempPath)

	var buffer bytes.Buffer
	rootCmd.SetOut(&buffer)
	rootCmd.SetArgs([]string{"-i", "../test/data/templates", "-o", testDir(t), "-d", yamlData1, "-t", ".template2"})

	err := rootCmd.Execute()

	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "", buffer.String())

	file, err := os.ReadFile(testDir(t) + "/test.txt")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "A=1", string(file))
}

func TestJSON(t *testing.T) {
	defer os.RemoveAll(TempPath)

	var buffer bytes.Buffer
	rootCmd.SetOut(&buffer)
	rootCmd.SetArgs([]string{"-i", "../test/data/templates", "-o", testDir(t), "-d", jsonData, "-t", ".template"})

	err := rootCmd.Execute()

	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "", buffer.String())

	file, err := os.ReadFile(testDir(t) + "/test.txt")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "A=1\nB=2", string(file))
}

func TestMultipleData(t *testing.T) {
	defer os.RemoveAll(TempPath)

	var buffer bytes.Buffer
	rootCmd.SetOut(&buffer)
	rootCmd.SetArgs([]string{"-i", "../test/data/templates", "-o", testDir(t), "-d", yamlData1, "-d", yamlData2, "-t", ".template3"})

	err := rootCmd.Execute()

	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "", buffer.String())

	file, err := os.ReadFile(testDir(t) + "/test.txt")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, "A=1\nC=3", string(file))
}

func testDir(t *testing.T) string {
	return TempPath + t.Name()
}
