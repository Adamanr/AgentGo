package main

import (
	client "AgentGo/assets"
)

func main() {
	var serialno = client.FileFind()
	client.StartProgram()
	print(serialno + "\n")

	print("-=-=-=-=Агент запущен=-=-=-=-\n")
	var number = client.Register(serialno)
	if number != "" {
		client.CliRun(number)
		print("-=-Камера зарегистрирована-=-\n")
	} else {
		print("Ключ не найден\n")
	}
}
