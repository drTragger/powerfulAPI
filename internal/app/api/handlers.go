package api

import (
	"encoding/json"
	"fmt"
	"github.com/drTragger/powerfulAPI/internal/app/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Message Helping struct to create messages
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func logEncode(err error) {
	if err != nil {
		log.Println("Error during encoding response message")
	}
}

func (api *API) GetAllArticles(writer http.ResponseWriter, _ *http.Request) {
	initHeaders(writer)
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing articles in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	api.logger.Info("Get All Articles GET /articles")
	writer.WriteHeader(http.StatusOK)
	logEncode(json.NewEncoder(writer).Encode(articles))
}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Create article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid JSON received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article")
		msg := Message{
			StatusCode: 501,
			Message:    "We are having some troubles accessing DataBase",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
	}
	writer.WriteHeader(201)
	logEncode(json.NewEncoder(writer).Encode(a))
}

func (api *API) GetArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get article by id GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Inappropriate id value",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles accessing articles:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We are having some troubles accessing DataBase. Try again later.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	if !ok {
		api.logger.Info(fmt.Sprintf("Could not find article with id %d", id))
		msg := Message{
			StatusCode: 404,
			Message:    "Article with this id does not exist",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	writer.WriteHeader(200)
	logEncode(json.NewEncoder(writer).Encode(article))
}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete article by id DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Inappropriate id value",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles accessing articles:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We are having some troubles accessing DataBase. Try again later.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	if !ok {
		api.logger.Info(fmt.Sprintf("Could not find article with id %d", id))
		msg := Message{
			StatusCode: 404,
			Message:    fmt.Sprintf("Article with id %d does not exist", id),
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles deleting article:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We are having some troubles accessing DataBase. Try again later.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with id %d has been deleted successfully", id),
		IsError:    false,
	}
	writer.WriteHeader(msg.StatusCode)
	logEncode(json.NewEncoder(writer).Encode(msg))
}

func (api *API) PostUserRegister(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	api.logger.Info("User register POST /api/v1/users/register")
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid JSON received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided JSON is invalid",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles accessing users:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We are having some troubles accessing DataBase. Try again later.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	if ok {
		api.logger.Info(fmt.Sprintf("User with login \"%s\" already exists", user.Login))
		msg := Message{
			StatusCode: 400,
			Message:    fmt.Sprintf("User %s already exists", user.Login),
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	newUser, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles accessing users:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We are having some troubles accessing DataBase. Try again later.",
			IsError:    true,
		}
		writer.WriteHeader(msg.StatusCode)
		logEncode(json.NewEncoder(writer).Encode(msg))
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User %s has been successfully registered", newUser.Login),
	}
	writer.WriteHeader(msg.StatusCode)
	logEncode(json.NewEncoder(writer).Encode(msg))
}
