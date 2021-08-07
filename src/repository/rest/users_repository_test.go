package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeOutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@email.com", "password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})

	respository := usersRepository{}

	user, err := respository.LoginUser("email@gmail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@email.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message":"invalid login credentals", "status":"404","error":"not_found"}`,
	})

	respository := usersRepository{}

	user, err := respository.LoginUser("email@gmail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)

}

func TestLoginUserInvalidUserResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@email.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":"1", "first_name":"test","last_name":"name", "email":"test@test.com"}`,
	})

	respository := usersRepository{}

	user, err := respository.LoginUser("email@gmail.com","password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@email.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":1, "first_name":"test","last_name":"name", "email":"test@test.com"}`,
	})

	respository := usersRepository{}

	user, err := respository.LoginUser("email@gmail.com","password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "test", user.FirstName)
	assert.EqualValues(t, "name", user.LastName)
	assert.EqualValues(t, "email", user.Email)

}
