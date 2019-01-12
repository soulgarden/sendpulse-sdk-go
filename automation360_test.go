package sendpulse

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
	"net/http"
	"testing"
)

func TestStartEventEmptyEventName(t *testing.T) {
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent("", variables)
	assert.Error(t, err)
	_, isResponseError := err.(*ResponseError)
	assert.False(t, isResponseError)
}

func TestStartEventNoPhoneAndEmail(t *testing.T) {
	eventName := fake.Word()
	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))
	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	_, isResponseError := err.(*ResponseError)
	assert.False(t, isResponseError)
}

func TestStartEventNotExists(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{"error_code": 102,"message": "Event not exists"}`
	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusBadRequest,
			respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	ResponseError, isResponseError := err.(*ResponseError)
	assert.True(t, isResponseError)
	assert.Equal(t, http.StatusBadRequest, ResponseError.HttpCode)
	assert.Equal(t, respBody, ResponseError.Body)
	assert.Equal(t, url, ResponseError.Url)
}

func TestStartEventDublicateData(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	respBody := `{"result": false,"message": "Dublicate data"}`

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK, respBody))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.Error(t, err)
	_, isResponseError := err.(*ResponseError)
	assert.False(t, isResponseError)
	assert.Equal(t, respBody, err.Error())
}

func TestStartEventWithPhone(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": true}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["phone"] = fake.Phone()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.NoError(t, err)
}

func TestStartEventWithEmail(t *testing.T) {
	eventName := fake.Word()

	url := fmt.Sprintf("%s/events/name/%s", apiBaseUrl, eventName)

	apiUid := fake.CharactersN(50)
	apiSecret := fake.CharactersN(50)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", apiBaseUrl+"/oauth/access_token",
		httpmock.NewStringResponder(http.StatusOK,
			`{"access_token": "testtoken","token_type": "Bearer","expires_in": 3600}`))

	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(http.StatusOK,
			`{"result": true}`))

	spClient, _ := ApiClient(apiUid, apiSecret, 5)

	variables := make(map[string]string)
	variables["email"] = fake.EmailAddress()
	variables["name"] = fake.FullName()
	err := spClient.Automation360.StartEvent(eventName, variables)
	assert.NoError(t, err)
}
