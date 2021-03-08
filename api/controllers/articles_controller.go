package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haryanapnx/go-blog-crud/api/auth"
	"github.com/haryanapnx/go-blog-crud/api/models"
	"github.com/haryanapnx/go-blog-crud/api/responses"
	"github.com/haryanapnx/go-blog-crud/api/utils/formaterror"
)

func (server *Server) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll((r.Body))

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	article := models.Article{}
	err = json.Unmarshal(body, &article)

	//validate if error
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	article.Prepare()
	err = article.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	aid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if aid != article.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	articleCreated, err := article.SaveArticle(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, articleCreated.ID))
	responses.JSON(w, http.StatusCreated, articleCreated)
}

func (server *Server) GetAllArticle(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.FindAllUser(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	article := models.Article{}
	articleItem, err := article.FindArticleByID(server.DB, uint64(aid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, articleItem)
}

func (server *Server) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Check if the post id is valid
	aid, err := strconv.ParseUint(vars["id"], 10, 12)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	// Check if article exist
	article := models.Article{}
	err = server.DB.Debug().Model(models.Article{}).Where("id = ?", aid).Take(&article).Error

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Article Not found!"))
		return
	}

	// If a user attempt to update a post not belonging to him
	if uid != article.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	articleUpdate := models.Article{}
	err = json.Unmarshal(body, &articleUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != articleUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	articleUpdate.Prepare()
	err = articleUpdate.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	articleUpdate.ID = article.ID

	articleUpdated, err := articleUpdate.UpdateArticle(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, articleUpdated)
}

func (server *Server) RemoveArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	post := models.Article{}
	err = server.DB.Debug().Model(models.Article{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.DeleteArticle(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")

}
