package main

import (
	"fmt"

	"github.com/mateusoliveira43/golang-studies/git"
)

// func buscaGitTag() {
// 	fmt.Println("Buscando commits pertencentes a git tag...")
// }

// var donoDoRepositorio string

func main() {
	fmt.Println("#VAAAAAAAAI")

	tag := "v1.26.2"
	// fmt.Println(tag)

	var donoDoRepositorio string
	donoDoRepositorio = "kubernetes"
	// fmt.Println(donoDoRepositorio)

	repo := "kubernetes"

	// fmt.Printf("%s %s\n", tag, donoDoRepositorio)

	// git.BuscaGitTag(tag, donoDoRepositorio, repo)

	// Orientação a objetos
	// b := new(git.BuscadorGit)
	b := git.BuscadorGit{DonoDoRepositorio: donoDoRepositorio, Repo: repo}
	b.BuscaGitTag(tag)

	// Concorrência
	// go git.BuscaGitTag()

	fmt.Println("Sai do programa")
}
