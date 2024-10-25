package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/model"
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
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return

	}
	ctx.JSON(http.StatusOK, devices)
}

func (d *deviceController) CreateDevice(ctx *gin.Context) {
	var device model.Device
	err := ctx.BindJSON(&device)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedDevice, err := d.deviceUseCase.CreateDevice(device)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return

	}
	ctx.JSON(http.StatusCreated, insertedDevice)
}

func (d *deviceController) GetDeviceById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "Device ID cannot be null.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deviceId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Device ID is not valid.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	device, err := d.deviceUseCase.GetDeviceById(deviceId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if device == nil {
		response := model.Response{
			Message: "Device not found.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, device)
}

func (d *deviceController) UpdateDeviceById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "Device ID cannot be null.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var device model.Device
	device.ID = id
	err := ctx.BindJSON(&device)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	updatedDevice, err := d.deviceUseCase.UpdateDeviceById(&device)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if updatedDevice == nil {
		response := model.Response{
			Message: "Device not found.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, updatedDevice)
}
