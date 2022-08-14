package assets

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var address = "You address =)"

func Register(number string) string {
	aqua := map[string]string{
		"deviceId": number,
	}
	data, _ := json.Marshal(aqua)

	req, err := http.NewRequest("POST", address, bytes.NewBuffer(data))
	if err != nil {
		panic("Error POST request")
	}
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Error client get response")
	}
	defer resp.Body.Close()
	if body, err := io.ReadAll(resp.Body); err == nil {
		print(string(body[0:9]) + "\n")
		if string(body[0:9]) == "{\"result\"" {
			return string(body[11:21])
		}
		return ""
	}
	panic("Error read Body")
}

func GetIp(deviceId string, ip string) {
	aqua := map[string]string{
		"deviceId": deviceId,
		"ip":       ip,
	}
	data, _ := json.Marshal(aqua)
	req, err := http.NewRequest("POST", address, bytes.NewBuffer(data))
	if err != nil {
		panic("Error POST request")
	}
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Error client get response")
	}
	defer func(resp *http.Response) {
		err = resp.Body.Close()
		if err != nil {
			panic("Error closed resp.body")
		}
	}(resp)
	panic("Error read Body")
}

func FileFind() string {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic("Error read config")
		}
		if !info.IsDir() && info.Name() == "system.conf" {
			filepath.Ext(path)
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	var reads = ReadConfig(files[0])
	return reads
}

func ReadConfig(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic("Not open files!")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic("Error read configuration")
		}
		var text = scanner.Text()
		if len(text) > 8 {
			if text[0:8] == "serialno" {
				return text[9:]
			}
		}
	}
	return ""
}

func CliRun(number string) {
	cmd := exec.Command("./vpn", "-s", number)
	if _, err := cmd.CombinedOutput(); err != nil {
		panic("Не удалось запустить бинарный файл!")
	}
	print("Подключён к VPN\n")
}

func StartProgram() string {
	var pass string
	flag.StringVar(&pass, "s", "Null", "Вы не задали ключ")
	flag.Parse()
	return pass
}
