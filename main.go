package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pessoa struct {
	ID     string `json:"id"`
	Nome   string `json:"nome"`
	Funcao string `json:"funcao"`
}

type equipe struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Id_Membro1 string `json:"id_membro1"`
	Id_Membro2 string `json:"id_membro2"`
	Id_Membro3 string `json:"id_membro3"`
}

type projeto struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Equipe    string `json:"equipe"`
	ID_Equipe string `json:"id_equipe"`
}

var projetos = []projeto{
	{ID: "1", Nome: "Projeto API", Equipe: "Os Batutinhas", ID_Equipe: "1"},
	{ID: "2", Nome: "Projeto Front-End", Equipe: "Os Calvos", ID_Equipe: "2"},
}

var equipes = []equipe{
	{ID: "1", Nome: "Os Batutinhas", Id_Membro1: "1", Id_Membro2: "2", Id_Membro3: "3"},
	{ID: "1", Nome: "Os Calvos", Id_Membro1: "4", Id_Membro2: "5", Id_Membro3: "6"},
}

var pessoas = []pessoa{
	{ID: "1", Nome: "Bruno", Funcao: "Back-end"},
	{ID: "2", Nome: "Iara", Funcao: "Front-end"},
	{ID: "3", Nome: "Lucas Martins", Funcao: "Back-end"},
	{ID: "4", Nome: "Lucas Emmanuel", Funcao: "Back-end"},
	{ID: "5", Nome: "Larissa", Funcao: "Front-end"},
	{ID: "6", Nome: "Caio", Funcao: "Front-end"},
}

func main() {
	router := gin.Default()
	router.GET("/projetos", getProjetos)
	router.GET("/projetos/:id", getProjetoByID)
	router.GET("/projetos/equipe/id/:id", getEquipeByID)
	router.GET("/projetos/equipe/nome/:nome", getEquipeByNome)
	router.POST("/projetos", postProjetos)
	router.DELETE("/projetos/:id", deleteProjetoById)
	router.PUT("/projetos/:id", updateProjetoById)

	router.GET("/pessoas", getPessoas)
	router.GET("/pessoas/:id", getPessoaByID)
	router.POST("/pessoas", postPessoas)
	router.DELETE("/pessoas/:id", deletePessoaById)
	router.PUT("/pessoas/:id", updatePessoaById)

	router.GET("/equipes", getEquipes)
	router.GET("/equipes/id/:id", getEquipeByID)
	router.GET("/equipes/nome/:nome", getEquipeByNome)
	router.GET("/equipes/pessoa/:id", getPessoaByID)
	router.POST("/equipes", postEquipes)
	router.DELETE("/equipes/:id", deleteEquipeById)
	router.PUT("/equipes/:id", updateEquipeById)

	router.Run("localhost:8080")
}

func getProjetos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projetos)
}

func getPessoas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pessoas)
}

func getEquipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, equipes)
}

func postProjetos(c *gin.Context) {
	var newProjeto projeto
	if err := c.BindJSON(&newProjeto); err != nil {
		return
	}
	projetos = append(projetos, newProjeto)
	c.IndentedJSON(http.StatusCreated, newProjeto)
}

func postPessoas(c *gin.Context) {
	var newPessoa pessoa
	if err := c.BindJSON(&newPessoa); err != nil {
		return
	}
	pessoas = append(pessoas, newPessoa)
	c.IndentedJSON(http.StatusCreated, newPessoa)
}

func postEquipes(c *gin.Context) {
	var newEquipe equipe
	if err := c.BindJSON(&newEquipe); err != nil {
		return
	}
	equipes = append(equipes, newEquipe)
	c.IndentedJSON(http.StatusCreated, newEquipe)
}

func getProjetoByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range projetos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "projeto not found"})
}

func getPessoaByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range pessoas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
}

func getEquipeByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range equipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipe not found"})
}

func getEquipeByNome(c *gin.Context) {
	nome := c.Param("nome")
	for _, a := range equipes {
		if a.Nome == nome {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipe not found"})
}

func deletePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range pessoas {
		if a.ID == id {
			pessoas = append(pessoas[:i], pessoas[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro deletado"})
			return
		}
	}
}

func deleteEquipeById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range equipes {
		if a.ID == id {
			equipes = append(equipes[:i], equipes[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro deletado"})
			return
		}
	}
}

func deleteProjetoById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range projetos {
		if a.ID == id {
			projetos = append(projetos[:i], projetos[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro deletado"})
			return
		}
	}
}

func updatePessoaById(c *gin.Context) {
	id := c.Param("id")
	for i := range pessoas {
		if pessoas[i].ID == id {
			c.BindJSON(&pessoas[i])
			c.IndentedJSON(http.StatusOK, pessoas[i])
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro atualizado"})
			return
		}
	}
}

func updateEquipeById(c *gin.Context) {
	id := c.Param("id")
	for i := range equipes {
		if equipes[i].ID == id {
			c.BindJSON(&equipes[i])
			c.IndentedJSON(http.StatusOK, equipes[i])
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro atualizado"})
			return
		}
	}
}

func updateProjetoById(c *gin.Context) {
	id := c.Param("id")
	for i := range projetos {
		if projetos[i].ID == id {
			c.BindJSON(&projetos[i])
			c.IndentedJSON(http.StatusOK, projetos[i])
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Cadastro atualizado"})
			return
		}
	}
}
