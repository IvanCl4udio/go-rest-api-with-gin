package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivancl4udio/gin-api-rest/database"
	"github.com/ivancl4udio/gin-api-rest/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos @godoc
// @Summary Exibir todos os alunos
// @Schemes
// @Description Rota para exibir todo os alunos
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {

	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{"API diz": "E ai " + nome + ", tudo beleza?"})
}

// CriaNovoAluno @godoc
// @Summary Cadastra um novo aluno
// @Schemes
// @Description Rota para cadastrar um novo aluno
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		// informa o erro caso seja nulo
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidadaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	// cria registro
	database.DB.Create(&aluno)
	// envia status de sucesso e o registro que foi criado
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorId @godoc
// @Summary Pesquisa um aluno pelo seu identificador único
// @Schemes
// @Description Rota para pesquisar um aluno pelo ID
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 404 {object} httputil.HTTPError
// @Router /alunos [get]
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno @godoc
// @Summary Exclui um aluno a partir do seu ID
// @Schemes
// @Description Rota para excluir um aluno
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno @godoc
// @Summary Edita as informações de um aluno
// @Schemes
// @Description Rota para alterar as informações de um aluno
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidadaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

// BuscaAlunoPorCPF @godoc
// @Summary Pesquisa por alunos a partir do numero do CPF
// @Schemes
// @Description Rota para pesquisar um aluno pelo CPF
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} httputil.HTTPError
// @Router /alunos [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
