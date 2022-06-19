package controlers

import (
	//"github.com/BogdanJas/Hospital/models"
	"net/http"
	"strconv"

	"github.com/BogdanJas/Hospital/models"
	"github.com/BogdanJas/Hospital/services"
	"github.com/gin-gonic/gin"
)

type PatientController struct {
	PatientService services.PatientService
}

func New(patientservice services.PatientService) PatientController {
	return PatientController{
		PatientService: patientservice,
	}
}

func (pc *PatientController) Createpatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	err := pc.PatientService.CreatePatient(&patient)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}

func (pc *PatientController) Getpatient(ctx *gin.Context) {
	patientname := ctx.Param("id")
	patientnameId,_ := strconv.Atoi(patientname)
	patient,err := pc.PatientService.GetPatient(&patientnameId)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

func (pc *PatientController) GetAll(ctx *gin.Context) {
	patients, err := pc.PatientService.GetAll()
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patients)
}

func (pc *PatientController) Updatepatient(ctx *gin.Context){
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	err := pc.PatientService.UpdatePatient(&patient)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}

func (pc *PatientController) Deletepatient(ctx *gin.Context){
	patientname := ctx.Param("name") 
	patientnameId,_ := strconv.Atoi(patientname)
	err := pc.PatientService.DeletePatient(&patientnameId)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
}

func (pc *PatientController) RegisterpatientRoutes(rg *gin.RouterGroup){
	patientroute := rg.Group("/patient")
	patientroute.POST("/create", pc.Createpatient)
	patientroute.GET("get/:name", pc.Getpatient)
	patientroute.GET("/get",pc.GetAll)
	patientroute.PATCH("/update",pc.Updatepatient)
	patientroute.DELETE("/delete/:name", pc.Deletepatient)
}
