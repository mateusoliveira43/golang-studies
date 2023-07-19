package git

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mateusoliveira43/golang-studies/handler"
)

// Orientação a objetos
type Buscador interface {
	BuscaGitTag(tag, donoDoRepositorio, repo string)
}

// type BuscadorGit struct{}
type BuscadorGit struct {
	DonoDoRepositorio string
	Repo              string
}

type JSONDoGithub struct {
	FullName    string   `json:"full_name"`
	HTMLUrl     string   `json:"html_url"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}

type JSONDeErroGithub struct {
	Message string `json:"message"`
	// DocumentationUrl string `json:"documentation_url"`
}

func (b *BuscadorGit) BuscaGitTag(tag string) {
	// func (b *BuscadorGit) BuscaGitTag(tag, donoDoRepositorio, repo string) {
	// func BuscaGitTag(tag, donoDoRepositorio, repo string) {
	// fmt.Println("Buscando commits pertencentes a git tag...")
	fmt.Printf("Buscando commits pertencentes a git tag %s em %s/%s\n", tag, b.DonoDoRepositorio, b.Repo)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", b.DonoDoRepositorio, b.Repo)
	resp, err := http.Get(url)
	if handler.HandleError(err) {
		bytes, errorBytes := io.ReadAll(resp.Body)
		if handler.HandleError(errorBytes) {
			if resp.StatusCode == 200 {
				// fmt.Println(resp.Body)
				// fmt.Println(string(bytes))
				j := new(JSONDoGithub)
				erro := json.Unmarshal(bytes, j)
				if handler.HandleError(erro) {
					fmt.Println(j.FullName)
					fmt.Println(j.Description)
					fmt.Println(j.HTMLUrl)
					fmt.Println(j.Topics)
				}
			} else {
				j := new(JSONDeErroGithub)
				erro := json.Unmarshal(bytes, j)
				if handler.HandleError(erro) {
					fmt.Printf("ERROR: HTTP Status Code %s, %s\n", fmt.Sprint(resp.StatusCode), j.Message)
				}
			}
		}
	}
}
