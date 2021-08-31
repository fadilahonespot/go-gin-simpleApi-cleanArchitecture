package handler

import (
	"net/http"
	"simpleApic/model"
	"simpleApic/person"
	"simpleApic/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	personUsecase person.PersonUsecase
}

func CreatePersonHandler(r *gin.Engine, personUsecase person.PersonUsecase) {
	personHandler := PersonHandler{personUsecase}

	r.POST("/person", personHandler.addPerson)
	r.GET("/person", personHandler.viewPersons)
	r.GET("person/:id", personHandler.viewPersonId)
	r.PUT("/person/:id", personHandler.editPerson)
	r.DELETE("/person/:id", personHandler.deletePerson)
}

func (e *PersonHandler) addPerson(c *gin.Context) {
	var person = model.Person{}
	err := c.Bind(&person)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if person.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}

	if person.FirstName == "" || person.LastName == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	newPerson, err := e.personUsecase.Create(&person)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, newPerson)
}

func (e *PersonHandler) viewPersons(c *gin.Context) {
	persons, err := e.personUsecase.ReadAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*persons) == 0 {
		utils.HandleError(c, http.StatusNotFound, "list person is empty")
		return
	}
	utils.HandleSucces(c, persons)
}

func (e *PersonHandler) viewPersonId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	person, err := e.personUsecase.ReadById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, person)
}

func (e *PersonHandler) editPerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.personUsecase.ReadById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var tempPerson = model.Person{}
	err = c.Bind(&tempPerson)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if tempPerson.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "input not permitted")
		return
	}
	if tempPerson.FirstName == "" || tempPerson.LastName == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	person, err := e.personUsecase.Update(id, &tempPerson)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, person)
}

func (e *PersonHandler) deletePerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	err = e.personUsecase.Delete(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, "success delete data")
}
