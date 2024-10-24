package controller

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/usecase"
	"github.com/gin-gonic/gin"
)

type deviceController struct {
	deviceUseCase usecase.DeviceUsecase
}

func NewDeviceController(usecase usecase.DeviceUsecase) deviceController {
	return deviceController{
		deviceUseCase: usecase,
	}
}

func (d *deviceController) GetDevices(ctx *gin.Context) {
	devices, err := d.deviceUseCase.GetDevices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, devices)
}
