//go:build integrations

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	ctx         context.Context
	cleanUpTest func()
	client      *resty.Client
}

func (s *ServerTestSuite) TestRegisterAndLogin() {
	var Id struct {
		Id string `json:"user_id"`
	}
	var Token struct {
		Token string `json:"token"`
	}
	r, err := s.client.R().SetBody(map[string]interface{}{
		"email":     "sleeter@client.com",
		"password":  "12345",
		"user_type": "client",
	}).Post("/register")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Id)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), Id.Id)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /register client")

	r, err = s.client.R().SetBody(map[string]interface{}{
		"id":       Id.Id,
		"password": "12345",
	}).Post("/login")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Token)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), Token.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /login client")

	r, err = s.client.R().SetBody(map[string]interface{}{
		"email":     "sleeter@moderator.com",
		"password":  "12345",
		"user_type": "moderator",
	}).Post("/register")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Id)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), Id.Id)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /register moderator")

	r, err = s.client.R().SetBody(map[string]interface{}{
		"id":       Id.Id,
		"password": "12345",
	}).Post("/login")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Token)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), Token.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /login moderator")
}

func (s *ServerTestSuite) TestHousePipeline() {
	type Token struct {
		Token string `json:"token"`
	}
	var Flats struct {
		Flats []response.Flat `json:"flats"`
	}
	moderatorToken := Token{}
	clientToken := Token{}

	// client dummy login request
	r, err := s.client.R().SetQueryParam("user_type", "client").Get("/dummyLogin")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &clientToken)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), clientToken.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /dummyLogin")

	// client get house request
	r, err = s.client.R().SetBody(
		map[string]interface{}{
			"id": 1,
		}).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", clientToken.Token),
	}).Get(fmt.Sprintf("/house/%d", 1))
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Flats)
	require.NoError(s.T(), err)
	for _, flat := range Flats.Flats {
		require.Equal(s.T(), flat.Status, entity.FLATSTATUS_APPROVED)
	}
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /house/{id} with client request")

	// moderator dummy login request
	r, err = s.client.R().SetQueryParam("user_type", "moderator").Get("/dummyLogin")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &moderatorToken)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), moderatorToken.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid Valid GET /dummyLogin")

	// moderator get house request
	r, err = s.client.R().SetBody(
		map[string]interface{}{
			"id": "1",
		}).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", moderatorToken.Token),
	}).Get(fmt.Sprintf("/house/%d", 1))
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /house/{id} with moderator request")
}

func (s *ServerTestSuite) TestFlatPipeline() {
	developer := "developer"
	HouseRequest := request.House{
		Address:   "Address",
		Year:      2021,
		Developer: &developer,
	}
	HouseResponse := response.House{}
	FlatCreate := request.CreateFlat{
		HouseId: 0,
		Price:   2024,
		Rooms:   nil,
	}
	FlatResponse := response.Flat{}
	FlatUpdate1 := request.UpdateFlat{
		Id:     0,
		Status: entity.FLATSTATUS_ON_MODERATION,
	}
	FlatUpdate2 := request.UpdateFlat{
		Id:     0,
		Status: entity.FLATSTATUS_APPROVED,
	}
	type Token struct {
		Token string `json:"token"`
	}
	var Flats struct {
		Flats []response.Flat `json:"flats"`
	}
	var moderatorToken Token
	var clientToken Token
	// moderator dummy login request
	r, err := s.client.R().SetQueryParam("user_type", "moderator").Get("/dummyLogin")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &moderatorToken)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), moderatorToken.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid Valid GET /dummyLogin")

	// create house request
	r, err = s.client.R().SetBody(HouseRequest).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", moderatorToken.Token),
	}).Post("/house/create")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &HouseResponse)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), HouseResponse.Id)
	require.Equal(s.T(), HouseRequest.Year, HouseResponse.Year)
	require.Equal(s.T(), HouseRequest.Address, HouseResponse.Address)
	require.Equal(s.T(), HouseRequest.Developer, HouseResponse.Developer)
	require.NotNil(s.T(), HouseResponse.CreatedAt)
	require.NotNil(s.T(), HouseResponse.UpdateAt)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /house/create")

	// client dummy login request
	r, err = s.client.R().SetQueryParam("user_type", "client").Get("/dummyLogin")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &clientToken)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), clientToken.Token)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /dummyLogin")

	// create flat request
	FlatCreate.HouseId = HouseResponse.Id
	r, err = s.client.R().SetBody(FlatCreate).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", clientToken.Token),
	}).Post("/flat/create")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &FlatResponse)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), FlatResponse.Id)
	require.Equal(s.T(), FlatCreate.HouseId, FlatResponse.HouseId)
	require.Equal(s.T(), FlatCreate.Price, FlatResponse.Price)
	require.Equal(s.T(), entity.FLATSTATUS_CREATED, FlatResponse.Status)
	require.NotNil(s.T(), FlatResponse.Rooms)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /flat/create")

	// update flat request
	FlatUpdate1.Id = FlatResponse.Id
	r, err = s.client.R().SetBody(FlatUpdate1).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", moderatorToken.Token),
	}).Post("/flat/update")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &FlatResponse)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), FlatResponse.Id)
	require.Equal(s.T(), FlatUpdate1.Status, FlatResponse.Status)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /flat/update")

	// update flat request
	FlatUpdate2.Id = FlatResponse.Id
	r, err = s.client.R().SetBody(FlatUpdate2).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", moderatorToken.Token),
	}).Post("/flat/update")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &FlatResponse)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), FlatResponse.Id)
	require.Equal(s.T(), FlatUpdate2.Status, FlatResponse.Status)
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), "Valid POST /flat/update")

	// get house request
	r, err = s.client.R().SetBody(
		map[string]interface{}{
			"id": 1,
		}).SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("bearer: %s", clientToken.Token),
	}).Get(fmt.Sprintf("/house/%d", HouseResponse.Id))
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	err = json.Unmarshal(r.Body(), &Flats)
	require.NoError(s.T(), err)
	require.Len(s.T(), Flats.Flats, 1)
	for _, flat := range Flats.Flats {
		require.Equal(s.T(), flat.Status, entity.FLATSTATUS_APPROVED)
	}
	require.Equal(s.T(), http.StatusOK, r.StatusCode(), fmt.Sprintf("Valid GET /house/%d with client request", HouseResponse.Id))
}

func (s *ServerTestSuite) SetupTest() {
	time.Sleep(time.Second / 2)

	s.ctx, s.cleanUpTest = context.WithTimeout(context.Background(), time.Second)

	c := resty.New()
	c.SetBaseURL("http://localhost:8080")

	s.client = c
}

func (s *ServerTestSuite) TearDownTest() {
	s.cleanUpTest()
}

func TestBase(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
