package main

import "io/ioutil"

func getTracks() ([]string, error) {
	var files []string

	fileInfos, err := ioutil.ReadDir("tracks")
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}

	return files, nil
}
