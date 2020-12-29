package command

import (
	"errors"
	"fmt"
)

func exchangeCommand(cmd string) (string, error) {
	if len(cmd) == 0 {
		return "", errors.New("cmd should provide")
	}

	return fmt.Sprintf("service %s", cmd), nil
}

//ErrCmdMissing 命令缺失
var ErrCmdMissing = errors.New("cmd should provide")

//ErrGameOver 不玩了
var ErrGameOver = errors.New("game over")

func exchangeCommand2(cmd string) (string, error) {
	if len(cmd) == 0 {
		return "", ErrCmdMissing
	}

	return fmt.Sprintf("service %s", cmd), nil
}

func exchangeCommand3(cmd string) string {
	if len(cmd) == 0 {
		panic(ErrCmdMissing)
	}

	return fmt.Sprintf("service %s", cmd)
}

func exchangeCommand4(cmd string) string {
	defer func() {
		if err := recover(); err != nil {
			if err == ErrCmdMissing {
				// 该错误不太重要，记录一下不管它
				fmt.Println(err)
			} else {
				// 其他的错误就致命了
				panic(err)
			}
		}
	}()

	if len(cmd) == 0 {
		panic(ErrCmdMissing)
	}

	if cmd == "gameover" {
		panic(ErrGameOver)
	}

	return fmt.Sprintf("service %s", cmd)
}
