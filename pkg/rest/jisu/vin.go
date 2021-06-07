package jisu

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/seosite/gcore/pkg/core/netx"
)

type QueryByVin struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Result QueryByVinInfo `json:"result"`
}

type QueryByVinInfo struct {
	Carid                  int64       `json:"carid"`
	Vin                    string      `json:"vin"`
	Name                   string      `json:"name"`
	Brand                  string      `json:"brand"`
	Typename               string      `json:"typename"`
	Logo                   string      `json:"logo"`
	Manufacturer           string      `json:"manufacturer"`
	Yeartype               string      `json:"yeartype"`
	Environmentalstandards string      `json:"environmentalstandards"`
	Comfuelconsumption     string      `json:"comfuelconsumption"`
	Engine                 string      `json:"engine"`
	Fueltype               string      `json:"fueltype"`
	Gearbox                string      `json:"gearbox"`
	Drivemode              string      `json:"drivemode"`
	Fronttiresize          string      `json:"fronttiresize"`
	Reartiresize           string      `json:"reartiresize"`
	Displacement           string      `json:"displacement"`
	Displacementml         string      `json:"displacementml"`
	Fuelgrade              string      `json:"fuelgrade"`
	Price                  string      `json:"price"`
	Chassis                interface{} `json:"chassis"`
	Frontbraketype         string      `json:"frontbraketype"`
	Rearbraketype          string      `json:"rearbraketype"`
	Parkingbraketype       string      `json:"parkingbraketype"`
	Maxpower               string      `json:"maxpower"`
	Sizetype               string      `json:"sizetype"`
	Gearnum                string      `json:"gearnum"`
	Geartype               string      `json:"geartype"`
	Seatnum                string      `json:"seatnum"`
	Bodystructure          string      `json:"bodystructure"`
	Maxhorsepower          string      `json:"maxhorsepower"`
	Iscorrect              int64       `json:"iscorrect"`
	Machineoil             struct {
		Volume    string `json:"volume"`
		Viscosity string `json:"viscosity"`
		Grade     string `json:"grade"`
		Level     string `json:"level"`
	} `json:"machineoil"`
	Listdate        string      `json:"listdate"`
	Model           string      `json:"model"`
	Marketprice     string      `json:"marketprice"`
	Version         string      `json:"version"`
	Groupid         string      `json:"groupid"`
	Groupname       string      `json:"groupname"`
	Isimport        int64       `json:"isimport"`
	Doornum         string      `json:"doornum"`
	Len             string      `json:"len"`
	Width           string      `json:"width"`
	Height          string      `json:"height"`
	Wheelbase       string      `json:"wheelbase"`
	Weight          string      `json:"weight"`
	Ratedloadweight interface{} `json:"ratedloadweight"`
	Bodytype        string      `json:"bodytype"`
	Enginemodel     string      `json:"enginemodel"`
	Cylindernum     string      `json:"cylindernum"`
	Fuelmethod      string      `json:"fuelmethod"`
	Carlist         []struct {
		Carid    int64  `json:"carid"`
		Name     string `json:"name"`
		Typeid   int64  `json:"typeid"`
		Typename string `json:"typename"`
	} `json:"carlist"`
}

func (j *JisuAPI) QueryByVin(vin string, c *gin.Context) (info QueryByVinInfo, err error) {
	url := "https://api.jisuapi.com/vin/query?appkey=" + j.AppKey + "&vin=" + vin

	var queryByVin QueryByVin

	res, err := netx.NewRetryClient().Post(url,"text/html",nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(res), &queryByVin); err != nil {
		return
	}

	return queryByVin.Result, nil
}
