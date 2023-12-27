package util

import (
	"fmt"
	logger "nb_client/config/logger"
	"regexp"
)

var (
	PatternFileNameENTRX = "((?i)ENTRX|RX|SFRAI)_%s_\\d{8}_\\d{1}((?i)\\.TXT\\.GZ)"
	PatternFileRX        = "(?i)RX[0-9]{8}.TXT\\.GZ"
)

func GetIntervalePeriodo(value string) (uint64, string, error) {
	interval := getValue(value, "\\d")
	period := getValue(value, "\\D")

	var intervalNumber uint64
	_, err := fmt.Sscan(interval, &intervalNumber)
	if err != nil {
		logger.Error("Erro ao converter o intervalo", err)
		return 0, "", err
	}

	return intervalNumber, period, nil

}

func getValue(value, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(value)
}

func IsValidPattern(pattern, fileName string) bool {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return regex.MatchString(fileName)
}

func GetTextByPatterns(pattern, value string) (string, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return regex.FindString(value), nil
}
