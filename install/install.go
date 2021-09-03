package install

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sikalabs/install-slu/utils/tar_gz_utils"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Install(sluVersion, sluOs, sluArch, outputFile string) {
	tar_gz_utils.WebTarGzToBin(
		"https://github.com/sikalabs/slu/releases/download/"+sluVersion+
			"/slu_"+sluVersion+"_"+sluOs+"_"+sluArch+".tar.gz",
		"slu",
		outputFile,
	)
}

type LatestReleaseResponse struct {
	TagName string `json:"tag_name"`
}

func GetLatestRelease() string {
	var err error

	resp, err := http.Get("https://api.github.com/repos/sikalabs/slu/releases/latest")
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	handleError(err)
	data := LatestReleaseResponse{}

	err = json.Unmarshal(body, &data)
	handleError(err)

	return data.TagName
}
