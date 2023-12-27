package read

import (
	"fmt"
	"nb_client/config/logger"
	"nb_client/models"
	"nb_client/pkg/converter"
	"nb_client/pkg/path"
	"nb_client/pkg/util"
	"os"
	"path/filepath"
)

func Files(config models.AutoConfig) {
	dir, err := os.Open(config.PathFileSend)
	if err != nil {
		print(err)
	}
	files, err := dir.ReadDir(-1)
	if err != nil {
		print(err)
	}

	var filesToSend []os.DirEntry

	for _, file := range files {
		patterns := fmt.Sprintf(util.PatternFileNameENTRX, config.License)
		if util.IsValidPattern(patterns, file.Name()) || util.IsValidPattern(util.PatternFileRX, file.Name()) {
			filesToSend = append(filesToSend, file)
		}
	}

	if filesToSend == nil {
		logger.Info(fmt.Sprintf("Nâo tem dados para transmitir da liceca %s", config.License))
		return
	}

	for _, f := range filesToSend {

		logger.Info(fmt.Sprintf("Leando arquivo %s", f.Name()))
		saleCode, err := util.GetTextByPatterns("\\d{8}", f.Name())
		if err != nil {
			logger.Error("Erro ao buscar o codigo do vendedor", err)
		}
		base, err := convertFileToBase64(dir.Name(), f.Name())
		if err != nil {
			logger.Error("Erro ao converter o arquivo para base64", err)
		}
		fmt.Printf("tamanho %d do arquivo %s \n", len(base), f.Name())
		fmt.Println(saleCode)

		//uploaded, _ := server.UploadFile(base, f.Name(), saleCode, config)
		//if uploaded {
		path.MoveToEnviados(dir.Name(), f.Name())
		//}
	}
}

func convertFileToBase64(dirName, fileName string) (string, error) {
	file := filepath.FromSlash(dirName + "/" + fileName)
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return converter.ToBase64(bytes), nil
}
