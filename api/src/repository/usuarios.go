package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuario no banco de dados
func (repository Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repository.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos usuarios com nome ou nick parecidos
func (repository Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repository.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarPorID traz um usuario
func (repository Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, err := repository.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}
