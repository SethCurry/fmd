package fmd

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Hash string   `yaml:"hash"`
	Tags []string `yaml:"tags"`
}

func FileMetadataPath(fileName string) string {
	return fileName + ".fmd"
}

func GetFileMetadata(fileName string) (*Metadata, error) {
	readPath := FileMetadataPath(fileName)

	readContents, err := ioutil.ReadFile(readPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file \"%s\"", readPath)
	}

	var md Metadata

	err = yaml.Unmarshal(readContents, &md)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid YAML in \"%s\"", readPath)
	}

	return &md, nil
}
