package main

import (
  "os"
  "log"
  "io/ioutil"
)

func main() {
  // Cria um arquivo ou abre se já existir
  file, err := os.Create("example.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  // Dados a serem escritos no arquivo
  content := []byte("Hello, file!")

  // Escreve dados no arquivo
  err = ioutil.WriteFile("example.txt", content, 0644)
  if err != nil {
    log.Fatal(err)
  }
}