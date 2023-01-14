package rotas

import (
	"net/http"
	"api/src/crontrollers"
)

var rotasUsuarios = []Rota{
	//register users  
	{
		URI:    			"/usuarios",
		Metodo: 			http.MethodPost,
		Funcao: 			controllers.CriarUsuarios
		RequerAutenticao: 	false,
	},
	//fetch all users 
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticao: false,
	},
	//search for a user
	{
		URI:    "/usuarios/{usuarioid}",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticao: false,
	},
	//update the user
	{
		URI:    "/usuarios/usuarioid",
		Metodo: http.MethodPut,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticao: false,
	},
	//delete a user 
	{
		URI:    "/usuarios/{usuarioid}",
		Metodo: http.MethodDelete,
		Funcao: func(w http.ResponseWriter, r *http.Request) {

		},
		RequerAutenticao: false,
	},
}