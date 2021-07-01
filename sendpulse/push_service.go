package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
	"strings"
	"time"
)

type PushService struct {
	client *Client
}

func newPushService(cl *Client) *PushService {
	return &PushService{client: cl}
}

type PushListParams struct {
	Limit     int
	Offset    int
	From      time.Time
	To        time.Time
	WebsiteID int
}

type Push struct {
	ID        int                 `json:"id"`
	Title     string              `json:"title"`
	Body      string              `json:"body"`
	WebsiteID int                 `json:"website_id"`
	From      models.DateTimeType `json:"from"`
	To        models.DateTimeType `json:"to"`
	Status    int                 `json:"status"`
}

func (service *PushService) List(params PushListParams) ([]Push, error) {
	path := "/push/tasks/"
	var urlParts []string
	urlParts = append(urlParts, fmt.Sprintf("offset=%d", params.Offset))
	if params.Limit != 0 {
		urlParts = append(urlParts, fmt.Sprintf("limit=%d", params.Limit))
	}
	if !params.From.IsZero() {
		urlParts = append(urlParts, fmt.Sprintf("from=%s", params.From.Format("2006-01-02")))
	}
	if !params.To.IsZero() {
		urlParts = append(urlParts, fmt.Sprintf("to=%s", params.From.Format("2006-01-02")))
	}
	if params.WebsiteID != 0 {
		urlParts = append(urlParts, fmt.Sprintf("website_id=%d", params.WebsiteID))
	}

	if len(urlParts) != 0 {
		path += "?" + strings.Join(urlParts, "&")
	}

	var respData []Push
	_, err := service.client.NewRequest(http.MethodGet, path, nil, &respData, true)
	return respData, err
}

func (service *PushService) WebsitesTotal() (int, error) {
	path := "/push/websites/total"
	var respData struct {
		Total int `json:"total"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Total, err
}

type PushWebsite struct {
	ID      int                 `json:"id"`
	Url     string              `json:"url"`
	AddDate models.DateTimeType `json:"add_date"`
	Status  int                 `json:"status"`
}

func (service *PushService) WebsitesList(limit, offset int) ([]*PushWebsite, error) {
	path := fmt.Sprintf("/push/websites/?limit=%d&offset=%d", limit, offset)
	var respData []*PushWebsite
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

type PushWebsiteVariable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (service *PushService) WebsiteVariables(websiteID int) ([]*PushWebsiteVariable, error) {
	path := fmt.Sprintf("/push/websites/%d/variables", websiteID)
	var respData []*PushWebsiteVariable
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

type WebsiteSubscriptionsParams struct {
	Limit  int
	Offset int
	From   time.Time
	To     time.Time
}

type WebsiteSubscription struct {
	ID               int                   `json:"id"`
	Browser          string                `json:"browser"`
	Lang             string                `json:"lang"`
	Os               string                `json:"os"`
	CountryCode      string                `json:"country_code"`
	City             string                `json:"city"`
	Variables        []PushWebsiteVariable `json:"variables"`
	SubscriptionDate models.DateTimeType   `json:"subscription_date"`
	Status           int                   `json:"status"`
}

func (service *PushService) WebsiteSubscriptions(websiteID int, params WebsiteSubscriptionsParams) ([]*WebsiteSubscription, error) {
	path := fmt.Sprintf("/push/websites/%d/subscriptions", websiteID)

	var urlParts []string
	urlParts = append(urlParts, fmt.Sprintf("offset=%d", params.Offset))
	if params.Limit != 0 {
		urlParts = append(urlParts, fmt.Sprintf("limit=%d", params.Limit))
	}
	if !params.From.IsZero() {
		urlParts = append(urlParts, fmt.Sprintf("subscription_date_from=%s", params.From.Format("2006-01-02")))
	}
	if !params.To.IsZero() {
		urlParts = append(urlParts, fmt.Sprintf("subscription_date_to=%s", params.From.Format("2006-01-02")))
	}

	if len(urlParts) != 0 {
		path += "?" + strings.Join(urlParts, "&")
	}

	var respData []*WebsiteSubscription
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *PushService) SubscriptionsTotal(websiteID int) (int, error) {
	path := fmt.Sprintf("/push/websites/%d/subscriptions/total", websiteID)
	var respData struct {
		Total int `json:"total"`
	}
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Total, err
}

type WebsiteInfo struct {
	ID                int                 `json:"id"`
	Url               string              `json:"url"`
	Status            string              `json:"status"`
	Icon              string              `json:"icon"`
	AddDate           models.DateTimeType `json:"add_date"`
	TotalSubscribers  int                 `json:"total_subscribers"`
	Unsubscribed      int                 `json:"unsubscribed"`
	SubscribersToday  int                 `json:"subscribers_today"`
	ActiveSubscribers int                 `json:"active_subscribers"`
}

func (service *PushService) WebsiteInfo(websiteID int) (*WebsiteInfo, error) {
	path := fmt.Sprintf("/push/websites/info/%d", websiteID)
	var respData *WebsiteInfo
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *PushService) ActivateSubscription(subscriptionID int) error {
	path := "/push/subscriptions/state"
	type paramsFormat struct {
		ID    int `json:"id"`
		State int `json:"state"`
	}

	data := paramsFormat{ID: subscriptionID, State: 1}

	var respData struct {
		Result bool `json:"true"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), data, &respData, true)
	return err
}

func (service *PushService) DeactivateSubscription(subscriptionID int) error {
	path := "/push/subscriptions/state"
	type paramsFormat struct {
		ID    int `json:"id"`
		State int `json:"state"`
	}

	data := paramsFormat{ID: subscriptionID, State: 0}

	var respData struct {
		Result bool `json:"true"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), data, &respData, true)
	return err
}

type PushTaskParams struct {
	Title                string    `json:"title"`
	WebsiteID            int       `json:"website_id"`
	Body                 string    `json:"body"`
	TtlSec               int       `json:"ttl"`
	Link                 string    `json:"link,omitempty"`
	FilterLang           string    `json:"filter_lang,omitempty"`
	FilterBrowser        string    `json:"filter_browser,omitempty"`
	FilterRegion         string    `json:"filter_region,omitempty"`
	FilterUrl            string    `json:"filter_url,omitempty"`
	SubscriptionDateFrom time.Time `json:"filter_subscription_date_from,omitempty"`
	SubscriptionDateTo   time.Time `json:"filter_subscription_date_to,omitempty"`
	Filter               *struct {
		VariableName string `json:"variable_name"`
		Operator     string `json:"operator"`
		Conditions   []struct {
			Condition string      `json:"condition"`
			Value     interface{} `json:"value"`
		} `json:"conditions"`
	} `json:"filter,omitempty"`
	StretchTimeSec int                 `json:"stretch_time"`
	SendDate       models.DateTimeType `json:"send_date"`
	Buttons        *struct {
		Text string `json:"text"`
		Link string `json:"link"`
	} `json:"buttons,omitempty"`
	Image *struct {
		Name       string `json:"name"`
		DataBase64 string `json:"data"`
	} `json:"image,omitempty"`
	Icon *struct {
		Name       string `json:"name"`
		DataBase64 string `json:"data"`
	} `json:"icon,omitempty"`
}

func (service *PushService) CreatePushTask(params PushTaskParams) (int, error) {
	path := "/push/tasks"

	var respData struct {
		ID     int  `json:"id"`
		Result bool `json:"true"`
	}
	_, err := service.client.NewRequest(http.MethodPost, fmt.Sprintf(path), params, &respData, true)
	return respData.ID, err
}

type PushTasksStatistics struct {
	ID      int `json:"id"`
	Message struct {
		Title string `json:"title"`
		Text  string `json:"text"`
		Link  string `json:"link"`
	}
	Website   string `json:"website"`
	WebsiteID int    `json:"website_id"`
	Status    int    `json:"status"`
	Send      int    `json:"send,string"`
	Delivered int    `json:"delivered"`
	Redirect  int    `json:"redirect"`
}

func (service *PushService) PushTaskStatistics(taskID int) (*PushTasksStatistics, error) {
	path := fmt.Sprintf("/push/tasks/%d", taskID)

	var respData *PushTasksStatistics
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}
