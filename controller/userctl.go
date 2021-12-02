package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getAccounts(authPath string) gin.Accounts {
	_f, e := os.Open(authPath)
	if e != nil {
		fmt.Errorf("Fail to find authFile %s\n", authPath)
		return nil
	}
	defer _f.Close()

	_br := bufio.NewReader(_f)
	_accounts := make(gin.Accounts, 3)
	for {
		_accPair, _, e := _br.ReadLine()
		if e == io.EOF {
			break
		}
		_accPairDict := strings.Split(string(_accPair), " ")
		if len(_accPairDict) != 2 {
			fmt.Errorf("Format error meets when loading auths.\n")
			return nil
		}
		_accounts[_accPairDict[0]] = _accPairDict[1]
	}
	return _accounts
}
