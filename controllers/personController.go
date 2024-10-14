package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mereb/v4/database"
	"github.com/mereb/v4/models"
	"github.com/mereb/v4/services"
)

func CreatePersonController(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Binding Error",
			"Status":  "Failure",

			"Error": err.Error()})
		return
	}

	var count int64
	if err := database.GetDB().Model(&models.Person{}).Where("name = ?", person.Name).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failure",
			"Message": "Error In Counting Person Uniqueness",
			"Error":   err.Error()})
		return
	}
	if count >= 1 {
		c.JSON(http.StatusConflict, gin.H{
			"Status":  "Failure",
			"Message": "Person Already Exist",
		})
		return
	}

	createdPerson, err := services.CreatePersonService(&person)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failure",
			"Message": "Internal Server Error",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Person Created Successfully",
		"Data":    createdPerson,
	})
}

func GetAllPersonsController(c *gin.Context) {

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	perPage := 3
	page := 0

	if currValue, err := strconv.Atoi(limitStr); err == nil && currValue >= 1 {
		perPage = currValue
	}

	if currValue, err := strconv.Atoi(offsetStr); err == nil && currValue >= 1 {
		page = currValue
	}
	persons, err := services.GetAllPersonsService(perPage, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failure",
			"Message": "Internal Server Error",
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Page":    page,
		"PerPage": perPage,
		"Status":  "Success",
		"Message": "Person Fetched Successfully ",
		"Data":    persons,
	})
}

func GetPersonController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failure",
			"Message": "Invalid Query Parameter",
			"Error":   err.Error()})
		return
	}

	person, err := services.GetPersonService(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status": "Failure",

			"Message": " Internal Server Error",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Person Fetched Successfully !!!",
		"Data":    person,
	})
}

func UpdatePersonController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failure",
			"Message": "Invalid ID",
			"Error":   err.Error()})
		return
	}

	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failure",
			"Message": "Binding Error",
			"Error":   err.Error()})
		return
	}

	person.PersonID = uint(id)

	if err := services.UpdatePersonService(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failure",
			"Message": "Internal Servaer Error",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Person Updated Successfully !",
		"Data":    person,
	})
}

func DeletePersonController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failure",
			"Message": "Invalid ID",
			"Error":   err.Error()})
		return
	}

	if err := services.DeletePersonServices(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failure",
			"Message": "Internal Server Error",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Person deleted successfully"})
}
