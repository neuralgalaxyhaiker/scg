package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/forgoer/openssl"
)

const KEY = "NeuralGalaxy6666"

// @lancher encode server.js
// @lancher server.js
func main() {
	if len(os.Args) == 4 {
		inFile, _ := filepath.Abs(os.Args[2])
		outFile, _ := filepath.Abs(os.Args[3])
		encryptFile(inFile, outFile)
	} else if len(os.Args) == 2 {
		inFile, _ := filepath.Abs(os.Args[1])
		lancher(inFile)
	} else {
		fmt.Println(`Usage: 
	lancher encode server.js server.out.js    加密程序
	lancher server.out.js                   运行加密程序`)
		os.Exit(-1)
	}
}

func encryptFile(inFile, outFile string) {
	if !FileExists(inFile) {
		fmt.Println("file not found: ", inFile)
		os.Exit(-1)
	}
	fileBytes, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("read file ：", inFile, " ", err)
		os.Exit(-1)
	}
	dst, err := openssl.AesECBEncrypt(fileBytes, []byte(KEY), openssl.PKCS7_PADDING)
	if err != nil {
		fmt.Println("encrypt error：", err)
		os.Exit(-1)
	}
	if err = ioutil.WriteFile(outFile, dst, 0666); err != nil {
		fmt.Println("write out file：", outFile, " ,error:", err)
		os.Exit(-1)
	}
}

func decryptFile(inFile string) (body []byte, err error) {
	if !FileExists(inFile) {
		return nil, fmt.Errorf("file not found：%s", inFile)
	}
	if body, err = ioutil.ReadFile(inFile); err != nil {
		return nil, fmt.Errorf("reaf file：%s, error: %s ", inFile, err)
	}
	body, err = openssl.AesECBDecrypt(body, []byte(KEY), openssl.PKCS7_PADDING)
	return
}

func lancher(inFile string) {
	fmt.Println("lancher start ----")
	// todo 这里是直接读取文件，可以把文件打入到golang程序内部
	content, err := decryptFile(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	cmd := exec.Command("node")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = bytes.NewBuffer(content)

	//read env config
	cmd.Env = os.Environ()

	//Read the env config file named .env or .env.production for node.
	for _, envFile := range []string{".env", ".env.production"} {
		absEnvFile, _ := filepath.Abs(envFile)
		if FileExists(absEnvFile) {
			if envFileFi, err := os.OpenFile(absEnvFile, os.O_RDONLY, 0666); err != nil {
				fmt.Println("warning: read env file ", absEnvFile, " error: ", err)
			} else {
				fmt.Println("load env file: ", absEnvFile)
				cmd.ExtraFiles = append(cmd.ExtraFiles, envFileFi)
			}
		}
	}

	if err = cmd.Run(); err != nil {
		fmt.Println("running error：", err)
		os.Exit(2)
	}
	fmt.Println("---- lancher end")
}

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			return false
		}
	}
	return true
}
