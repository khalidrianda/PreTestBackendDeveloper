package delivery

import (
	"net/http"
	"strconv"

	"testcase/features/main/entity"

	"github.com/labstack/echo"
)

type adminHandler struct {
	srv entity.Services
}

func New(e *echo.Echo, srv entity.Services) {
	handler := adminHandler{srv: srv}
	e.GET("/", GetHello()) //
	e.GET("/language", GetLanguage())
	e.GET("/palindrome", handler.GetPalindrom())
	e.POST("/language", SaveData())
	e.GET("/language/:id", GetLanguagebyID())
	e.GET("/languages", GetLanguages())
	e.PATCH("/language/:id", PatchLanguage())
	e.DELETE("/language/:id", DeleteLanguage())
}

// 2. Buatkan struct di Go dengan data.
var data []entity.Data = []entity.Data{{Language: "C", Appeared: 1972, Created: []string{"Dennis Ritchie"},
	Functional: true, ObjectOriented: false,
	Relation: struct {
		InfluencedBy []string "json:\"influenced-by\""
		Influences   []string "json:\"influences\""
	}{InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
		Influences: []string{"C++", "Objective-C", "C#", "Java", "Javascript", "PHP", "Go"}}}}

// 3. GET / —> return response body “Hello Go developers”
func GetHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, SuccessResponseNoData("Hello Go developers"))
	}
}

// 3. GET /language —> return response body berupa JSON data pada soal nomor 2
func GetLanguage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, SuccessResponseData(data[0]))
	}
}

// 4. Tambahkan endpoint GET /palindrome dengan input parameter bernama “text”
func (ah *adminHandler) GetPalindrom() echo.HandlerFunc {
	return func(c echo.Context) error {
		text := c.QueryParam("text")

		result := ah.srv.Palindrome(text)
		if result == "Not palindrome" {
			return c.JSON(http.StatusBadRequest, SuccessResponseNoData(result))
		}

		return c.JSON(http.StatusOK, SuccessResponseNoData(result))
	}
}

// 5. POST /language untuk menambahkan data bahasa pemrograman seperti pada soal nomor 2
func SaveData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.Data
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Wrong input id"))
		}
		data = append(data, input)

		return c.JSON(http.StatusOK, SuccessResponseData(input))
	}
}

// 5. Ubah endpoint GET /language menjadi GET /language/<id> dengan <id> adalah slice index.
func GetLanguagebyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Wrong input id"))
		}

		return c.JSON(http.StatusOK, SuccessResponseData(data[id]))
	}
}

// 5. GET /languages untuk mendapatkan semua data bahasa pemrograman yang tersimpan
func GetLanguages() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, SuccessResponseData(data))
	}
}

// 5. PATCH /language/<id> untuk mengubah data bahasa pemrograman
func PatchLanguage() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Wrong input id"))
		}

		var input entity.Data
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Wrong input id"))
		}
		data[id] = input

		return c.JSON(http.StatusOK, SuccessResponseData(input))
	}
}

// 5. DELETE /language/<id> untuk menghapus data bahasa pemrograman
func DeleteLanguage() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Wrong input id"))
		}

		data = RemoveIndex(data, id)

		return c.JSON(http.StatusOK, SuccessResponseData(data))
	}
}

func RemoveIndex(s []entity.Data, index int) []entity.Data {
	return append(s[:index], s[index+1:]...)
}
