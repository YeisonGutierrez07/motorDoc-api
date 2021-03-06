package company

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/motorDoc-api/v1/app/global"
	"github.com/motorDoc-api/v1/entities"
)

// GetAll funcion para traer toda la información de las empresas
func GetAll(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)
	name := c.Query("name")

	companies, err := GetCompanies(user, name)
	if err != nil {
		response := global.ResponseServices(user, "401", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(companies, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// CreateNewCompany registrar de nuevas empresas
func CreateNewCompany(c *gin.Context) {
	companyData := CreateCompany{}
	err := c.ShouldBind(&companyData)
	if err == nil {
		r, errRegister := RegisterNewCompany(companyData)
		if errRegister != nil {
			response := global.ResponseServices(companyData, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(r, "200", "Se he creado la empresa con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(companyData, "400", err.Error())
	c.JSON(400, response)
}

// ChangeStatusCompany Cambiar el estado de las empresas
func ChangeStatusCompany(c *gin.Context) {
	changeStatus := entities.ChangeStatus{}
	err := c.ShouldBind(&changeStatus)
	if err == nil {
		r, errRegister := ChangeStatusService(changeStatus)
		if errRegister != nil {
			response := global.ResponseServices(changeStatus, "400", errRegister.Error())
			c.JSON(400, response)
			return
		}
		response := global.ResponseServices(r, "200", "Se he cambiado el estado con exito")
		c.JSON(http.StatusOK, response)
		return
	}
	response := global.ResponseServices(changeStatus, "400", err.Error())
	c.JSON(400, response)
}

// DeleteCompany funcion para eliminar empresas
func DeleteCompany(c *gin.Context) {
	idCompany, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	changeStatus := entities.ChangeStatus{}
	changeStatus.CompanyID = idCompany
	changeStatus.Status = 2

	r, errRegister := ChangeStatusService(changeStatus)
	if errRegister != nil {
		response := global.ResponseServices(changeStatus, "400", errRegister.Error())
		c.JSON(400, response)
		return
	}
	response := global.ResponseServices(r, "200", "Se he eliminado la empresa con exito")
	c.JSON(http.StatusOK, response)
	return
}

// MyCompany funcion para traerla  información de mi compañia
func MyCompany(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	company, err := GetCompanyByUser(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(company, "200", "Exito")
	c.JSON(http.StatusOK, response)
}

// MyWorkShop funcion para traerla  información de los mecanicos talleres
func MyWorkShop(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	mechanic, err := GetMyWorkshops(user)
	if err != nil {
		response := global.ResponseServices(user, "400", err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := global.ResponseServices(mechanic, "200", "Exito")
	c.JSON(http.StatusOK, response)
}
