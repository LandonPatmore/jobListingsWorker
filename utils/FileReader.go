package utils

import (
	"bufio"
	"os"
)

func ReadStationsList() ([]string, error) {
	fullPath, _ := os.Getwd()
	file, err := os.Open(fullPath + "/dataPullerWorker/files/stationsList.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [] string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
