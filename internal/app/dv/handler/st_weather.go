package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/chenxinlong/dv/configs"

	aw "github.com/deanishe/awgo"
)

const (
	HandlerWeather = "weather"
)

type WeatherHandler struct {
	typ HandlerType
	key string
}

type WeatherResult struct {
	FxLink string `json:"fxLink"`
	Now    struct {
		ObsTime   string `json:"obsTime"` // observe time
		Temp      string `json:"temp"`    // temperature
		Text      string `json:"text"`
		WindScale string `json:"windScale"` // wind scale
		Humidity  string `json:"humidity"`
		Precip    string `json:"precip"` // unit(mm)
	} `json:"now"`
}

func NewHandlerWeather() *WeatherHandler {
	return &WeatherHandler{
		typ: HandlerTypShortTag,
		key: configs.Conf.Weather.Key,
	}
}

func (g WeatherHandler) GetType() HandlerType {
	return g.typ
}

func (g WeatherHandler) Handle(e Event) (item *aw.Item) {
	arr := strings.Split(e.input, "-w ")
	if len(arr) != 2 {
		return
	}
	location := arr[1]

	// eg.厦门，locationID = 101230201
	locationID, err := g.GetLocationID(location)
	if err != nil {
		return
	}
	result, err := g.GetWeather(locationID)
	if err != nil {
		return
	}

	// construct items
	e.wf.NewItem(fmt.Sprintf("在和风天气上查看 %q 的天气详情", location)).Valid(true).Arg(result.FxLink)

	// overview
	title := fmt.Sprintf("[%s] 天气:%s, 温度:%s", location, result.Now.Text, result.Now.Temp)
	e.wf.NewItem(title).Valid(true).Arg(result.FxLink)

	title = fmt.Sprintf("[%s] 当前小时降水 : %s 毫米", location, result.Now.Precip)
	e.wf.NewItem(title).Valid(true).Arg(result.FxLink)

	return nil
}

func (g WeatherHandler) GetLocationID(name string) (id string, err error) {
	resp, err := http.Get(fmt.Sprintf("https://geoapi.qweather.com/v2/city/lookup?location=%s&key=%s", name, g.key))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	type respGetLocation struct {
		Location []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"location"`
	}
	var _resp respGetLocation
	if err = json.Unmarshal(b, &_resp); err != nil || len(_resp.Location) == 0 {
		return
	}

	return _resp.Location[0].ID, nil
}

func (g WeatherHandler) GetWeather(locationID string) (result WeatherResult, err error) {
	resp, err := http.Get(fmt.Sprintf("https://devapi.qweather.com/v7/weather/now?location=%s&key=%s", locationID, g.key))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		return
	}

	return
}
