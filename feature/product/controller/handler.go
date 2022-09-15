package controller

import (
	"ecommerce-project/config"
	"ecommerce-project/feature/product/entities"
	"ecommerce-project/middlewares"
	"ecommerce-project/utils/helper"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	FromTo entities.ServiceInterface
}

func New(e *echo.Echo, data entities.ServiceInterface) {
	handler := &Delivery{
		FromTo: data,
	}

	e.GET("/profile/product", handler.GetMyProduct, middlewares.JWTMiddleware())
	e.PUT("/profile/product/:id", handler.UpdateProduct, middlewares.JWTMiddleware())
	e.POST("/product", handler.AddProduct, middlewares.JWTMiddleware())
	e.GET("/product", handler.Get8All)
	e.GET("/product/:id", handler.Detail)
	e.DELETE("/product/:id", handler.DeleteProduct, middlewares.JWTMiddleware())
}

func (user *Delivery) AddProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var request Request
	errbind := c.Bind(&request)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Failed Bind Data"))
	}
	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto")
	if fotoerr != http.ErrMissingFile || fotoerr == nil {
		format, errf := helper.CheckFile(infoFoto.Filename)
		if errf != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Format Error"))
		}
		//
		err_image_size := helper.CheckSize(infoFoto.Size)
		if err_image_size != nil {
			return c.JSON(http.StatusBadRequest, err_image_size)
		}
		//
		waktu := fmt.Sprintf("%v", time.Now())
		imageName := strconv.Itoa(userid) + "_" + request.Name + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("failed to upload file"))
		}

		request.Foto = imageaddress
	}

	core := request.ReqToCore(uint(userid))

	row, erri := user.FromTo.AddProductI(core)
	if erri != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Gagal ke database"))
	}

	return c.JSON(http.StatusCreated, helper.Success("Sukses menambahkan"))
}

func (user *Delivery) Get8All(c echo.Context) error {

	page, er := strconv.Atoi(c.QueryParam("page"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("query param must number"))
	}

	listcore, err := user.FromTo.GetAll(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Terjadi Kesalahan"))
	}
	listRes := CoreToResponseList(listcore)

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses mendapatkan data", listRes))
}

func (user *Delivery) DeleteProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	productid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, helper.Failed("Id msut nomor"))
	}
	msg, errs := user.FromTo.Delete(userid, productid)
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) GetMyProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	page, er := strconv.Atoi(c.QueryParam("page"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("page must number"))
	}

	mycore, err := user.FromTo.GetMyProduct(userid, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
	}
	myres := CoreToResponseList(mycore)

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses mendapatkan data", myres))
}

func (user *Delivery) UpdateProduct(c echo.Context) error {
	userid := middlewares.ExtractToken(c)

	var request Request

	errbind := c.Bind(&request)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal bind data"))
	}

	dataFoto, infoFoto, fotoerr := c.Request().FormFile("foto")

	if fotoerr != http.ErrMissingFile || fotoerr == nil {
		format, errf := helper.CheckFile(infoFoto.Filename)
		if errf != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Format Error"))
		}
		waktu := fmt.Sprintf("%v", time.Now())
		imageName := strconv.Itoa(userid) + "_" + waktu + "." + format

		imageaddress, errupload := helper.UploadFileToS3(config.FolderName, imageName, config.FileType, dataFoto)
		if errupload != nil {
			fmt.Println(errupload)
			return c.JSON(http.StatusInternalServerError, helper.Failed("failed to upload file"))
		}
		request.Foto = imageaddress
	}

	idproduct, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Failed("Parameter must be number"))
	}
	core := request.ReqToCore(uint(userid))
	msg, errs := user.FromTo.UpdateMyProduct(core, uint(idproduct))
	if errs != nil {
		c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) Detail(c echo.Context) error {
	idproduct, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Failed("Parameter must be number"))
	}
	core, er := user.FromTo.GetAProduct(idproduct)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Gagal Mendapatkan data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses Mendapatkan product", core))
}
