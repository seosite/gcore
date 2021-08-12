package jisu

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/seosite/gcore/pkg/core/netx"
	"github.com/spf13/cast"
)

type JisuAPI struct {
	AppKey    string
	AppSecret string
}

// QueryCarBrandListResp =============品牌================
type QueryCarBrandListResp struct {
	Status int64          `json:"status"`
	Msg    string         `json:"msg"`
	Result []CarBrandInfo `json:"result"`
}

type CarBrandInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Initial  string `json:"initial"`
	ParentID int64  `json:"parentid"`
	Logo     string `json:"logo"`
	Depth    int64  `json:"depth"`
}

func (j *JisuAPI) QueryCarBrandList(c *gin.Context) (list []CarBrandInfo, err error) {
	url := "https://api.jisuapi.com/car/brand?appkey=" + j.AppKey

	var carBrandList QueryCarBrandListResp
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(res), &carBrandList); err != nil {
		return nil, err
	}

	return carBrandList.Result, nil
}

// CarSeriesList =============车系================
type CarSeriesList struct {
	Status int64           `json:"status"`
	Msg    string          `json:"msg"`
	Result []CarSeriesInfo `json:"result"`
}

type CarSeriesInfo struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	Fullname interface{} `json:"fullname"`
	Initial  string      `json:"initial"`
	List     []struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Fullname  string `json:"fullname"`
		Logo      string `json:"logo"`
		Salestate string `json:"salestate"`
		Depth     int64  `json:"depth"`
	} `json:"list"`
}

func (j *JisuAPI) QueryCarSeriesByBrandIDList(brandID int64, c *gin.Context) (list []CarSeriesInfo, err error) {
	if brandID <= 0 {
		return
	}

	url := "https://api.jisuapi.com/car/type?appkey=" + j.AppKey + "&parentid=" + cast.ToString(brandID)

	var carSeriesList CarSeriesList
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(res), &carSeriesList); err != nil {
		return nil, err
	}

	return carSeriesList.Result, nil
}

// CarModelsList =============车型================
type CarModelsList struct {
	Status int64         `json:"status"`
	Msg    string        `json:"msg"`
	Result CarModelsInfo `json:"result"`
}

type CarModelsInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Initial   string `json:"initial"`
	Fullname  string `json:"fullname"`
	Logo      string `json:"logo"`
	Salestate string `json:"salestate"`
	Depth     int64  `json:"depth"`
	List      []struct {
		ID              int64  `json:"id"`
		Name            string `json:"name"`
		Logo            string `json:"logo"`
		Price           string `json:"price"`
		Yeartype        string `json:"yeartype"`
		Listdate        string `json:"listdate"`
		Productionstate string `json:"productionstate"`
		Salestate       string `json:"salestate"`
		Sizetype        string `json:"sizetype"`
		Displacement    string `json:"displacement"`
		Displacement2   string `json:"displacement2"`
		Geartype        string `json:"geartype"`
		Geartype2       int64  `json:"geartype2"`
		Groupid         string `json:"groupid"`
		Groupname       string `json:"groupname"`
	} `json:"list"`
}

func (j *JisuAPI) QueryCarModelsBySeriesIDList(seriesID int64, c *gin.Context) (info CarModelsInfo, err error) {
	if seriesID <= 0 {
		return
	}

	url := "https://api.jisuapi.com/car/car?appkey=" + j.AppKey + "&parentid=" + cast.ToString(seriesID)

	var carModelsList CarModelsList
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(res), &carModelsList); err != nil {
		return
	}

	return carModelsList.Result, nil
}

// CarInfo =============车详情================
type CarInfo struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Result CarInfoData `json:"result"`
}

type CarInfoData struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name"`
	Parentname              string `json:"parentname"`
	Brandname               string `json:"brandname"`
	Initial                 string `json:"initial"`
	Parentid                int64  `json:"parentid"`
	Logo                    string `json:"logo"`
	Price                   string `json:"price"`
	Yeartype                string `json:"yeartype"`
	Listdate                string `json:"listdate"`
	Productionstate         string `json:"productionstate"`
	Salestate               string `json:"salestate"`
	Sizetype                string `json:"sizetype"`
	Depth                   int64  `json:"depth"`
	Displacement            string `json:"displacement"`
	Displacement2           string `json:"displacement2"`
	Gearnum                 string `json:"gearnum"`
	Geartype                string `json:"geartype"`
	Geartype2               int64  `json:"geartype2"`
	Seatnum                 string `json:"seatnum"`
	Drivemode               string `json:"drivemode"`
	Drivemode2              int64  `json:"drivemode2"`
	Environmentalstandards  string `json:"environmentalstandards"`
	Environmentalstandards2 string `json:"environmentalstandards2"`
	Compartnum              int64  `json:"compartnum"`
	Groupid                 string `json:"groupid"`
	Groupname               string `json:"groupname"`
	Basic                   struct {
		Price                       string `json:"price"`
		Saleprice                   string `json:"saleprice"`
		Warrantypolicy              string `json:"warrantypolicy"`
		Vechiletax                  string `json:"vechiletax"`
		Displacement                string `json:"displacement"`
		Gearbox                     string `json:"gearbox"`
		Gearnum                     string `json:"gearnum"`
		Geartype                    string `json:"geartype"`
		Comfuelconsumption          string `json:"comfuelconsumption"`
		Userfuelconsumption         string `json:"userfuelconsumption"`
		Officialaccelerationtime100 string `json:"officialaccelerationtime100"`
		Testaccelerationtime100     string `json:"testaccelerationtime100"`
		Maxspeed                    string `json:"maxspeed"`
		Seatnum                     string `json:"seatnum"`
		Mixfuelconsumption          string `json:"mixfuelconsumption"`
	} `json:"basic"`
	Body struct {
		Color              string      `json:"color"`
		Len                string      `json:"len"`
		Width              string      `json:"width"`
		Height             string      `json:"height"`
		Wheelbase          string      `json:"wheelbase"`
		Fronttrack         string      `json:"fronttrack"`
		Reartrack          string      `json:"reartrack"`
		Weight             string      `json:"weight"`
		Fullweight         string      `json:"fullweight"`
		Mingroundclearance string      `json:"mingroundclearance"`
		Approachangle      string      `json:"approachangle"`
		Departureangle     string      `json:"departureangle"`
		Luggagevolume      string      `json:"luggagevolume"`
		Luggagemode        string      `json:"luggagemode"`
		Luggageopenmode    string      `json:"luggageopenmode"`
		Inductionluggage   string      `json:"inductionluggage"`
		Doornum            string      `json:"doornum"`
		Tooftype           string      `json:"tooftype"`
		Hoodtype           string      `json:"hoodtype"`
		Roofluggagerack    string      `json:"roofluggagerack"`
		Sportpackage       string      `json:"sportpackage"`
		Totalweight        interface{} `json:"totalweight"`
		Ratedloadweight    interface{} `json:"ratedloadweight"`
		Loadweightfactor   interface{} `json:"loadweightfactor"`
		Rampangle          string      `json:"rampangle"`
		Maxwadingdepth     string      `json:"maxwadingdepth"`
		Minturndiameter    string      `json:"minturndiameter"`
		Electricluggage    string      `json:"electricluggage"`
		Bodytype           string      `json:"bodytype"`
	} `json:"body"`
	Engine struct {
		Position               string      `json:"position"`
		Model                  string      `json:"model"`
		Modeleasyepc2          string      `json:"modeleasyepc2"`
		Modelsohu              string      `json:"modelsohu"`
		Displacement           string      `json:"displacement"`
		Displacementml         string      `json:"displacementml"`
		Intakeform             string      `json:"intakeform"`
		Cylinderarrangetype    string      `json:"cylinderarrangetype"`
		Cylindernum            string      `json:"cylindernum"`
		Valvetrain             string      `json:"valvetrain"`
		Valvestructure         string      `json:"valvestructure"`
		Compressionratio       string      `json:"compressionratio"`
		Bore                   string      `json:"bore"`
		Stroke                 string      `json:"stroke"`
		Maxhorsepower          string      `json:"maxhorsepower"`
		Maxpower               string      `json:"maxpower"`
		Maxpowerspeed          string      `json:"maxpowerspeed"`
		Maxtorque              string      `json:"maxtorque"`
		Maxtorquespeed         string      `json:"maxtorquespeed"`
		Fueltype               string      `json:"fueltype"`
		Fuelgrade              string      `json:"fuelgrade"`
		Fuelmethod             string      `json:"fuelmethod"`
		Fueltankcapacity       string      `json:"fueltankcapacity"`
		Cylinderheadmaterial   string      `json:"cylinderheadmaterial"`
		Cylinderbodymaterial   string      `json:"cylinderbodymaterial"`
		Environmentalstandards string      `json:"environmentalstandards"`
		Startstopsystem        string      `json:"startstopsystem"`
		Motorpower             string      `json:"motorpower"`
		Motortorque            string      `json:"motortorque"`
		Integratedpower        string      `json:"integratedpower"`
		Integratedtorque       string      `json:"integratedtorque"`
		Frontmaxpower          string      `json:"frontmaxpower"`
		Frontmaxtorque         string      `json:"frontmaxtorque"`
		Rearmaxpower           string      `json:"rearmaxpower"`
		Rearmaxtorque          interface{} `json:"rearmaxtorque"`
		Batterycapacity        string      `json:"batterycapacity"`
		Powerconsumption       string      `json:"powerconsumption"`
		Maxmileage             string      `json:"maxmileage"`
		Batterywarranty        interface{} `json:"batterywarranty"`
		Batteryfastchargetime  string      `json:"batteryfastchargetime"`
		Batteryslowchargetime  string      `json:"batteryslowchargetime"`
		Nedcmaxmileage         string      `json:"nedcmaxmileage"`
	} `json:"engine"`
	Gearbox struct {
		Gearbox      string `json:"gearbox"`
		Gearnum      string `json:"gearnum"`
		Geartype     string `json:"geartype"`
		Shiftpaddles string `json:"shiftpaddles"`
	} `json:"gearbox"`
	Chassisbrake struct {
		Bodystructure          string      `json:"bodystructure"`
		Powersteering          string      `json:"powersteering"`
		Frontbraketype         string      `json:"frontbraketype"`
		Rearbraketype          string      `json:"rearbraketype"`
		Parkingbraketype       string      `json:"parkingbraketype"`
		Drivemode              string      `json:"drivemode"`
		Airsuspension          string      `json:"airsuspension"`
		Adjustablesuspension   string      `json:"adjustablesuspension"`
		Frontsuspensiontype    string      `json:"frontsuspensiontype"`
		Rearsuspensiontype     string      `json:"rearsuspensiontype"`
		Centerdifferentiallock string      `json:"centerdifferentiallock"`
		Chassisid              interface{} `json:"chassisid"`
		Chassis                string      `json:"chassis"`
		Chassiscompany         interface{} `json:"chassiscompany"`
		Chassistype            interface{} `json:"chassistype"`
		Axlenum                interface{} `json:"axlenum"`
		Axleload               interface{} `json:"axleload"`
		Leafspringnum          interface{} `json:"leafspringnum"`
		Frontsuspension        interface{} `json:"frontsuspension"`
		Rearsuspension         interface{} `json:"rearsuspension"`
	} `json:"chassisbrake"`
	Safe struct {
		Airbagdrivingposition     string `json:"airbagdrivingposition"`
		Airbagfrontpassenger      string `json:"airbagfrontpassenger"`
		Airbagfrontside           string `json:"airbagfrontside"`
		Airbagfronthead           string `json:"airbagfronthead"`
		Airbagknee                string `json:"airbagknee"`
		Airbagrearside            string `json:"airbagrearside"`
		Airbagrearhead            string `json:"airbagrearhead"`
		Safetybeltprompt          string `json:"safetybeltprompt"`
		Safetybeltlimiting        string `json:"safetybeltlimiting"`
		Safetybeltpretightening   string `json:"safetybeltpretightening"`
		Frontsafetybeltadjustment string `json:"frontsafetybeltadjustment"`
		Rearsafetybelt            string `json:"rearsafetybelt"`
		Tirepressuremonitoring    string `json:"tirepressuremonitoring"`
		Zeropressurecontinued     string `json:"zeropressurecontinued"`
		Centrallocking            string `json:"centrallocking"`
		Childlock                 string `json:"childlock"`
		Remotekey                 string `json:"remotekey"`
		Keylessentry              string `json:"keylessentry"`
		Keylessstart              string `json:"keylessstart"`
		Engineantitheft           string `json:"engineantitheft"`
		Brakeassist               string `json:"brakeassist"`
		Sideaircurtain            string `json:"sideaircurtain"`
		Seatbeltairbag            string `json:"seatbeltairbag"`
		Rearcentralairbag         string `json:"rearcentralairbag"`
		Remotecontrol             string `json:"remotecontrol"`
		Smartkey                  string `json:"smartkey"`
	} `json:"safe"`
	Wheel struct {
		Fronttiresize string      `json:"fronttiresize"`
		Reartiresize  string      `json:"reartiresize"`
		Sparetiretype string      `json:"sparetiretype"`
		Hubmaterial   string      `json:"hubmaterial"`
		Fronttrack    interface{} `json:"fronttrack"`
		Reartrack     interface{} `json:"reartrack"`
		Tirenum       string      `json:"tirenum"`
	} `json:"wheel"`
	Drivingauxiliary struct {
		Abs                       string `json:"abs"`
		Ebd                       string `json:"ebd"`
		Brakeassist               string `json:"brakeassist"`
		Tractioncontrol           string `json:"tractioncontrol"`
		Esp                       string `json:"esp"`
		Eps                       string `json:"eps"`
		Automaticparking          string `json:"automaticparking"`
		Hillstartassist           string `json:"hillstartassist"`
		Hilldescent               string `json:"hilldescent"`
		Frontparkingradar         string `json:"frontparkingradar"`
		Reversingradar            string `json:"reversingradar"`
		Reverseimage              string `json:"reverseimage"`
		Panoramiccamera           string `json:"panoramiccamera"`
		Cruisecontrol             string `json:"cruisecontrol"`
		Adaptivecruise            string `json:"adaptivecruise"`
		Gps                       string `json:"gps"`
		Automaticparkingintoplace string `json:"automaticparkingintoplace"`
		Ldws                      string `json:"ldws"`
		Activebraking             string `json:"activebraking"`
		Integralactivesteering    string `json:"integralactivesteering"`
		Nightvisionsystem         string `json:"nightvisionsystem"`
		Blindspotdetection        string `json:"blindspotdetection"`
		Lanekeep                  string `json:"lanekeep"`
		Parallelaid               string `json:"parallelaid"`
		Fatiguereminder           string `json:"fatiguereminder"`
		Remoteparking             string `json:"remoteparking"`
		Autodriveassist           string `json:"autodriveassist"`
		Variablesteering          string `json:"variablesteering"`
		Drivemodechoose           string `json:"drivemodechoose"`
	} `json:"drivingauxiliary"`
	Doormirror struct {
		Openstyle                string `json:"openstyle"`
		Electricwindow           string `json:"electricwindow"`
		Uvinterceptingglass      string `json:"uvinterceptingglass"`
		Privacyglass             string `json:"privacyglass"`
		Antipinchwindow          string `json:"antipinchwindow"`
		Skylightopeningmode      string `json:"skylightopeningmode"`
		Skylightstype            string `json:"skylightstype"`
		Rearwindowsunshade       string `json:"rearwindowsunshade"`
		Rearsidesunshade         string `json:"rearsidesunshade"`
		Rearwiper                string `json:"rearwiper"`
		Sensingwiper             string `json:"sensingwiper"`
		Electricpulldoor         string `json:"electricpulldoor"`
		Rearmirrorwithturnlamp   string `json:"rearmirrorwithturnlamp"`
		Externalmirrormemory     string `json:"externalmirrormemory"`
		Externalmirrorheating    string `json:"externalmirrorheating"`
		Externalmirrorfolding    string `json:"externalmirrorfolding"`
		Externalmirroradjustment string `json:"externalmirroradjustment"`
		Rearviewmirrorantiglare  string `json:"rearviewmirrorantiglare"`
		Sunvisormirror           string `json:"sunvisormirror"`
		Autoheadlight            string `json:"autoheadlight"`
		Headlightfeature         string `json:"headlightfeature"`
		Rearviewmirrormedia      string `json:"rearviewmirrormedia"`
		Externalmirrormedia      string `json:"externalmirrormedia"`
		Externalmirrorantiglare  string `json:"externalmirrorantiglare"`
		Frontwiper               string `json:"frontwiper"`
		Electricsuctiondoor      string `json:"electricsuctiondoor"`
		Electricslidingdoor      string `json:"electricslidingdoor"`
		Roofrack                 string `json:"roofrack"`
		Rearwing                 string `json:"rearwing"`
		Frontelectricwindow      string `json:"frontelectricwindow"`
		Rearelectricwindow       string `json:"rearelectricwindow"`
	} `json:"doormirror"`
	Light struct {
		Headlighttype                   string `json:"headlighttype"`
		Optionalheadlighttype           string `json:"optionalheadlighttype"`
		Headlightautomaticopen          string `json:"headlightautomaticopen"`
		Headlightautomaticclean         string `json:"headlightautomaticclean"`
		Headlightdelayoff               string `json:"headlightdelayoff"`
		Headlightdynamicsteering        string `json:"headlightdynamicsteering"`
		Headlightilluminationadjustment string `json:"headlightilluminationadjustment"`
		Headlightdimming                string `json:"headlightdimming"`
		Frontfoglight                   string `json:"frontfoglight"`
		Readinglight                    string `json:"readinglight"`
		Interiorairlight                string `json:"interiorairlight"`
		Daytimerunninglight             string `json:"daytimerunninglight"`
		Ledtaillight                    string `json:"ledtaillight"`
		Lightsteeringassist             string `json:"lightsteeringassist"`
		Leddaytimerunninglight          string `json:"leddaytimerunninglight"`
	} `json:"light"`
	Internalconfig struct {
		Steeringwheelbeforeadjustment string `json:"steeringwheelbeforeadjustment"`
		Steeringwheelupadjustment     string `json:"steeringwheelupadjustment"`
		Steeringwheeladjustmentmode   string `json:"steeringwheeladjustmentmode"`
		Steeringwheelmemory           string `json:"steeringwheelmemory"`
		Steeringwheelmaterial         string `json:"steeringwheelmaterial"`
		Steeringwheelmultifunction    string `json:"steeringwheelmultifunction"`
		Steeringwheelheating          string `json:"steeringwheelheating"`
		Computerscreen                string `json:"computerscreen"`
		Interiorcolor                 string `json:"interiorcolor"`
		Rearcupholder                 string `json:"rearcupholder"`
		Supplyvoltage                 string `json:"supplyvoltage"`
		Interiormaterial              string `json:"interiormaterial"`
		Steeringwheelshift            string `json:"steeringwheelshift"`
		Activenoisereduction          string `json:"activenoisereduction"`
	} `json:"internalconfig"`
	Seat struct {
		Sportseat                           string `json:"sportseat"`
		Seatmaterial                        string `json:"seatmaterial"`
		Seatheightadjustment                string `json:"seatheightadjustment"`
		Driverseatadjustmentmode            string `json:"driverseatadjustmentmode"`
		Auxiliaryseatadjustmentmode         string `json:"auxiliaryseatadjustmentmode"`
		Driverseatlumbarsupportadjustment   string `json:"driverseatlumbarsupportadjustment"`
		Driverseatshouldersupportadjustment string `json:"driverseatshouldersupportadjustment"`
		Frontseatheadrestadjustment         string `json:"frontseatheadrestadjustment"`
		Rearseatadjustmentmode              string `json:"rearseatadjustmentmode"`
		Rearseatreclineproportion           string `json:"rearseatreclineproportion"`
		Rearseatangleadjustment             string `json:"rearseatangleadjustment"`
		Frontseatcenterarmrest              string `json:"frontseatcenterarmrest"`
		Rearseatcenterarmrest               string `json:"rearseatcenterarmrest"`
		Seatventilation                     string `json:"seatventilation"`
		Seatheating                         string `json:"seatheating"`
		Seatmassage                         string `json:"seatmassage"`
		Electricseatmemory                  string `json:"electricseatmemory"`
		Childseatfixdevice                  string `json:"childseatfixdevice"`
		Thirdrowseat                        string `json:"thirdrowseat"`
		Driverseatelectricadjustment        string `json:"driverseatelectricadjustment"`
		Auxiliaryseatelectricadjustment     string `json:"auxiliaryseatelectricadjustment"`
		Secondrowseatelectricadjustment     string `json:"secondrowseatelectricadjustment"`
		Frontseatfunction                   string `json:"frontseatfunction"`
		Rearseatfunction                    string `json:"rearseatfunction"`
		Secondrowseatadjustment             string `json:"secondrowseatadjustment"`
	} `json:"seat"`
	Entcom struct {
		Locationservice        string `json:"locationservice"`
		Bluetooth              string `json:"bluetooth"`
		Externalaudiointerface string `json:"externalaudiointerface"`
		Builtinharddisk        string `json:"builtinharddisk"`
		Cartv                  string `json:"cartv"`
		Speakernum             int64  `json:"speakernum"`
		Audiobrand             string `json:"audiobrand"`
		Dvd                    string `json:"dvd"`
		Cd                     string `json:"cd"`
		Consolelcdscreen       string `json:"consolelcdscreen"`
		Rearlcdscreen          string `json:"rearlcdscreen"`
		Lcdscreensize          string `json:"lcdscreensize"`
		Fulllcddashboard       string `json:"fulllcddashboard"`
		Huddisplay             string `json:"huddisplay"`
		Roadrescue             string `json:"roadrescue"`
		FourG                  string `json:"4g"`
		Carapp                 string `json:"carapp"`
		Voicecontrol           string `json:"voicecontrol"`
		Phoneconnect           string `json:"phoneconnect"`
		Wirelesscharge         string `json:"wirelesscharge"`
		Gesturecontrol         string `json:"gesturecontrol"`
		Drivingrecorder        string `json:"drivingrecorder"`
	} `json:"entcom"`
	Aircondrefrigerator struct {
		Airconditioningcontrolmode string `json:"airconditioningcontrolmode"`
		Tempzonecontrol            string `json:"tempzonecontrol"`
		Rearairconditioning        string `json:"rearairconditioning"`
		Reardischargeoutlet        string `json:"reardischargeoutlet"`
		Airconditioning            string `json:"airconditioning"`
		Airpurifyingdevice         string `json:"airpurifyingdevice"`
		Carrefrigerator            string `json:"carrefrigerator"`
		Frontairconditioning       string `json:"frontairconditioning"`
		Fragrance                  string `json:"fragrance"`
	} `json:"aircondrefrigerator"`
	Actualtest struct {
		Accelerationtime100 string `json:"accelerationtime100"`
		Brakingdistance     string `json:"brakingdistance"`
	} `json:"actualtest"`
}

func (j *JisuAPI) QueryCarInfoByModelsIDList(modelsID int64, c *gin.Context) (info CarInfoData, err error) {
	if modelsID <= 0 {
		return
	}

	url := "https://api.jisuapi.com/car/car?appkey=" + j.AppKey + "&carid=" + cast.ToString(modelsID)

	var carInfo CarInfo
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(res), &carInfo); err != nil {
		return
	}

	return carInfo.Result, nil
}

// CarSearchRes =============搜索详情================
type CarSearchRes struct {
	Status int64         `json:"status"`
	Msg    string        `json:"msg"`
	Result CarSearchInfo `json:"result"`
}

type CarSearchInfo struct {
	Total   int64  `json:"total"`
	Keyword string `json:"keyword"`
	List    []struct {
		ID              int64  `json:"id"`
		Name            string `json:"name"`
		Brandname       string `json:"brandname"`
		Parentname      string `json:"parentname"`
		Logo            string `json:"logo"`
		Price           string `json:"price"`
		Yeartype        string `json:"yeartype"`
		Listdate        string `json:"listdate"`
		Productionstate string `json:"productionstate"`
		Salestate       string `json:"salestate"`
		Sizetype        string `json:"sizetype"`
	} `json:"list"`
}

func (j *JisuAPI) QueryCarByKeywordList(keyword string, c *gin.Context) (info CarSearchInfo, err error) {
	if keyword == "" {
		return
	}

	url := "https://api.jisuapi.com/car/car?appkey=" + j.AppKey + "&keyword=" + keyword

	var carInfo CarSearchRes
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(res), &carInfo); err != nil {
		return
	}

	return carInfo.Result, nil
}

// CarDetailRes =============搜索详情================
type CarDetailRes struct {
	Status int64         `json:"status"`
	Msg    string        `json:"msg"`
	Result CarDetailInfo `json:"result"`
}

type CarDetailInfo struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name"`
	Parentname              string `json:"parentname"`
	Brandname               string `json:"brandname"`
	Initial                 string `json:"initial"`
	Parentid                int64  `json:"parentid"`
	Logo                    string `json:"logo"`
	Price                   string `json:"price"`
	Yeartype                string `json:"yeartype"`
	Listdate                string `json:"listdate"`
	Productionstate         string `json:"productionstate"`
	Salestate               string `json:"salestate"`
	Sizetype                string `json:"sizetype"`
	Depth                   int64  `json:"depth"`
	Displacement            string `json:"displacement"`
	Displacement2           string `json:"displacement2"`
	Gearnum                 string `json:"gearnum"`
	Geartype                string `json:"geartype"`
	Geartype2               int64  `json:"geartype2"`
	Seatnum                 string `json:"seatnum"`
	Drivemode               string `json:"drivemode"`
	Drivemode2              int64  `json:"drivemode2"`
	Environmentalstandards  string `json:"environmentalstandards"`
	Environmentalstandards2 string `json:"environmentalstandards2"`
	Compartnum              int    `json:"compartnum"`
	Groupid                 string `json:"groupid"`
	Groupname               string `json:"groupname"`
	Basic                   struct {
		Price                       string `json:"price"`
		Saleprice                   string `json:"saleprice"`
		Warrantypolicy              string `json:"warrantypolicy"`
		Vechiletax                  string `json:"vechiletax"`
		Displacement                string `json:"displacement"`
		Gearbox                     string `json:"gearbox"`
		Gearnum                     string `json:"gearnum"`
		Geartype                    string `json:"geartype"`
		Comfuelconsumption          string `json:"comfuelconsumption"`
		Userfuelconsumption         string `json:"userfuelconsumption"`
		Officialaccelerationtime100 string `json:"officialaccelerationtime100"`
		Testaccelerationtime100     string `json:"testaccelerationtime100"`
		Maxspeed                    string `json:"maxspeed"`
		Seatnum                     string `json:"seatnum"`
		Mixfuelconsumption          string `json:"mixfuelconsumption"`
	} `json:"basic"`
	Body struct {
		Color              string      `json:"color"`
		Len                string      `json:"len"`
		Width              string      `json:"width"`
		Height             string      `json:"height"`
		Wheelbase          string      `json:"wheelbase"`
		Fronttrack         string      `json:"fronttrack"`
		Reartrack          string      `json:"reartrack"`
		Weight             string      `json:"weight"`
		Fullweight         string      `json:"fullweight"`
		Mingroundclearance string      `json:"mingroundclearance"`
		Approachangle      string      `json:"approachangle"`
		Departureangle     string      `json:"departureangle"`
		Luggagevolume      string      `json:"luggagevolume"`
		Luggagemode        string      `json:"luggagemode"`
		Luggageopenmode    string      `json:"luggageopenmode"`
		Inductionluggage   string      `json:"inductionluggage"`
		Doornum            string      `json:"doornum"`
		Tooftype           string      `json:"tooftype"`
		Hoodtype           string      `json:"hoodtype"`
		Roofluggagerack    string      `json:"roofluggagerack"`
		Sportpackage       string      `json:"sportpackage"`
		Totalweight        interface{} `json:"totalweight"`
		Ratedloadweight    interface{} `json:"ratedloadweight"`
		Loadweightfactor   interface{} `json:"loadweightfactor"`
		Rampangle          interface{} `json:"rampangle"`
		Maxwadingdepth     interface{} `json:"maxwadingdepth"`
		Minturndiameter    interface{} `json:"minturndiameter"`
		Electricluggage    string      `json:"electricluggage"`
		Bodytype           string      `json:"bodytype"`
	} `json:"body"`
	Engine struct {
		Position               string      `json:"position"`
		Model                  string      `json:"model"`
		Displacement           string      `json:"displacement"`
		Displacementml         string      `json:"displacementml"`
		Intakeform             string      `json:"intakeform"`
		Cylinderarrangetype    string      `json:"cylinderarrangetype"`
		Cylindernum            string      `json:"cylindernum"`
		Valvetrain             string      `json:"valvetrain"`
		Valvestructure         string      `json:"valvestructure"`
		Compressionratio       string      `json:"compressionratio"`
		Bore                   string      `json:"bore"`
		Stroke                 string      `json:"stroke"`
		Maxhorsepower          string      `json:"maxhorsepower"`
		Maxpower               string      `json:"maxpower"`
		Maxpowerspeed          string      `json:"maxpowerspeed"`
		Maxtorque              string      `json:"maxtorque"`
		Maxtorquespeed         string      `json:"maxtorquespeed"`
		Fueltype               string      `json:"fueltype"`
		Fuelgrade              string      `json:"fuelgrade"`
		Fuelmethod             string      `json:"fuelmethod"`
		Fueltankcapacity       string      `json:"fueltankcapacity"`
		Cylinderheadmaterial   string      `json:"cylinderheadmaterial"`
		Cylinderbodymaterial   string      `json:"cylinderbodymaterial"`
		Environmentalstandards string      `json:"environmentalstandards"`
		Startstopsystem        string      `json:"startstopsystem"`
		Motorpower             interface{} `json:"motorpower"`
		Motortorque            interface{} `json:"motortorque"`
		Integratedpower        interface{} `json:"integratedpower"`
		Integratedtorque       interface{} `json:"integratedtorque"`
		Frontmaxpower          interface{} `json:"frontmaxpower"`
		Frontmaxtorque         interface{} `json:"frontmaxtorque"`
		Rearmaxpower           interface{} `json:"rearmaxpower"`
		Rearmaxtorque          interface{} `json:"rearmaxtorque"`
		Batterycapacity        interface{} `json:"batterycapacity"`
		Powerconsumption       interface{} `json:"powerconsumption"`
		Maxmileage             interface{} `json:"maxmileage"`
		Batterywarranty        interface{} `json:"batterywarranty"`
		Batteryfastchargetime  interface{} `json:"batteryfastchargetime"`
		Batteryslowchargetime  interface{} `json:"batteryslowchargetime"`
		Nedcmaxmileage         interface{} `json:"nedcmaxmileage"`
	} `json:"engine"`
	Gearbox struct {
		Gearbox      string `json:"gearbox"`
		Gearnum      string `json:"gearnum"`
		Geartype     string `json:"geartype"`
		Shiftpaddles string `json:"shiftpaddles"`
	} `json:"gearbox"`
	Chassisbrake struct {
		Bodystructure          string      `json:"bodystructure"`
		Powersteering          string      `json:"powersteering"`
		Frontbraketype         string      `json:"frontbraketype"`
		Rearbraketype          string      `json:"rearbraketype"`
		Parkingbraketype       string      `json:"parkingbraketype"`
		Drivemode              string      `json:"drivemode"`
		Airsuspension          string      `json:"airsuspension"`
		Adjustablesuspension   string      `json:"adjustablesuspension"`
		Frontsuspensiontype    string      `json:"frontsuspensiontype"`
		Rearsuspensiontype     string      `json:"rearsuspensiontype"`
		Centerdifferentiallock string      `json:"centerdifferentiallock"`
		Chassisid              interface{} `json:"chassisid"`
		Chassis                string      `json:"chassis"`
		Chassiscompany         interface{} `json:"chassiscompany"`
		Chassistype            interface{} `json:"chassistype"`
		Axlenum                interface{} `json:"axlenum"`
		Axleload               interface{} `json:"axleload"`
		Leafspringnum          interface{} `json:"leafspringnum"`
		Frontsuspension        interface{} `json:"frontsuspension"`
		Rearsuspension         interface{} `json:"rearsuspension"`
	} `json:"chassisbrake"`
	Safe struct {
		Airbagdrivingposition     string `json:"airbagdrivingposition"`
		Airbagfrontpassenger      string `json:"airbagfrontpassenger"`
		Airbagfrontside           string `json:"airbagfrontside"`
		Airbagfronthead           string `json:"airbagfronthead"`
		Airbagknee                string `json:"airbagknee"`
		Airbagrearside            string `json:"airbagrearside"`
		Airbagrearhead            string `json:"airbagrearhead"`
		Safetybeltprompt          string `json:"safetybeltprompt"`
		Safetybeltlimiting        string `json:"safetybeltlimiting"`
		Safetybeltpretightening   string `json:"safetybeltpretightening"`
		Frontsafetybeltadjustment string `json:"frontsafetybeltadjustment"`
		Rearsafetybelt            string `json:"rearsafetybelt"`
		Tirepressuremonitoring    string `json:"tirepressuremonitoring"`
		Zeropressurecontinued     string `json:"zeropressurecontinued"`
		Centrallocking            string `json:"centrallocking"`
		Childlock                 string `json:"childlock"`
		Remotekey                 string `json:"remotekey"`
		Keylessentry              string `json:"keylessentry"`
		Keylessstart              string `json:"keylessstart"`
		Engineantitheft           string `json:"engineantitheft"`
		Brakeassist               string `json:"brakeassist"`
		Sideaircurtain            string `json:"sideaircurtain"`
		Seatbeltairbag            string `json:"seatbeltairbag"`
		Rearcentralairbag         string `json:"rearcentralairbag"`
		Remotecontrol             string `json:"remotecontrol"`
		Smartkey                  string `json:"smartkey"`
	} `json:"safe"`
	Wheel struct {
		Fronttiresize string      `json:"fronttiresize"`
		Reartiresize  string      `json:"reartiresize"`
		Sparetiretype string      `json:"sparetiretype"`
		Hubmaterial   string      `json:"hubmaterial"`
		Fronttrack    interface{} `json:"fronttrack"`
		Reartrack     interface{} `json:"reartrack"`
		Tirenum       interface{} `json:"tirenum"`
	} `json:"wheel"`
	Drivingauxiliary struct {
		Abs                       string `json:"abs"`
		Ebd                       string `json:"ebd"`
		Brakeassist               string `json:"brakeassist"`
		Tractioncontrol           string `json:"tractioncontrol"`
		Esp                       string `json:"esp"`
		Eps                       string `json:"eps"`
		Automaticparking          string `json:"automaticparking"`
		Hillstartassist           string `json:"hillstartassist"`
		Hilldescent               string `json:"hilldescent"`
		Frontparkingradar         string `json:"frontparkingradar"`
		Reversingradar            string `json:"reversingradar"`
		Reverseimage              string `json:"reverseimage"`
		Panoramiccamera           string `json:"panoramiccamera"`
		Cruisecontrol             string `json:"cruisecontrol"`
		Adaptivecruise            string `json:"adaptivecruise"`
		Gps                       string `json:"gps"`
		Automaticparkingintoplace string `json:"automaticparkingintoplace"`
		Ldws                      string `json:"ldws"`
		Activebraking             string `json:"activebraking"`
		Integralactivesteering    string `json:"integralactivesteering"`
		Nightvisionsystem         string `json:"nightvisionsystem"`
		Blindspotdetection        string `json:"blindspotdetection"`
		Lanekeep                  string `json:"lanekeep"`
		Parallelaid               string `json:"parallelaid"`
		Fatiguereminder           string `json:"fatiguereminder"`
		Remoteparking             string `json:"remoteparking"`
		Autodriveassist           string `json:"autodriveassist"`
		Variablesteering          string `json:"variablesteering"`
		Drivemodechoose           string `json:"drivemodechoose"`
	} `json:"drivingauxiliary"`
	Doormirror struct {
		Openstyle                string `json:"openstyle"`
		Electricwindow           string `json:"electricwindow"`
		Uvinterceptingglass      string `json:"uvinterceptingglass"`
		Privacyglass             string `json:"privacyglass"`
		Antipinchwindow          string `json:"antipinchwindow"`
		Skylightopeningmode      string `json:"skylightopeningmode"`
		Skylightstype            string `json:"skylightstype"`
		Rearwindowsunshade       string `json:"rearwindowsunshade"`
		Rearsidesunshade         string `json:"rearsidesunshade"`
		Rearwiper                string `json:"rearwiper"`
		Sensingwiper             string `json:"sensingwiper"`
		Electricpulldoor         string `json:"electricpulldoor"`
		Rearmirrorwithturnlamp   string `json:"rearmirrorwithturnlamp"`
		Externalmirrormemory     string `json:"externalmirrormemory"`
		Externalmirrorheating    string `json:"externalmirrorheating"`
		Externalmirrorfolding    string `json:"externalmirrorfolding"`
		Externalmirroradjustment string `json:"externalmirroradjustment"`
		Rearviewmirrorantiglare  string `json:"rearviewmirrorantiglare"`
		Sunvisormirror           string `json:"sunvisormirror"`
		Autoheadlight            string `json:"autoheadlight"`
		Headlightfeature         string `json:"headlightfeature"`
		Rearviewmirrormedia      string `json:"rearviewmirrormedia"`
		Externalmirrormedia      string `json:"externalmirrormedia"`
		Externalmirrorantiglare  string `json:"externalmirrorantiglare"`
		Frontwiper               string `json:"frontwiper"`
		Electricsuctiondoor      string `json:"electricsuctiondoor"`
		Electricslidingdoor      string `json:"electricslidingdoor"`
		Roofrack                 string `json:"roofrack"`
		Rearwing                 string `json:"rearwing"`
		Frontelectricwindow      string `json:"frontelectricwindow"`
		Rearelectricwindow       string `json:"rearelectricwindow"`
	} `json:"doormirror"`
	Light struct {
		Headlighttype                   string `json:"headlighttype"`
		Optionalheadlighttype           string `json:"optionalheadlighttype"`
		Headlightautomaticopen          string `json:"headlightautomaticopen"`
		Headlightautomaticclean         string `json:"headlightautomaticclean"`
		Headlightdelayoff               string `json:"headlightdelayoff"`
		Headlightdynamicsteering        string `json:"headlightdynamicsteering"`
		Headlightilluminationadjustment string `json:"headlightilluminationadjustment"`
		Headlightdimming                string `json:"headlightdimming"`
		Frontfoglight                   string `json:"frontfoglight"`
		Readinglight                    string `json:"readinglight"`
		Interiorairlight                string `json:"interiorairlight"`
		Daytimerunninglight             string `json:"daytimerunninglight"`
		Ledtaillight                    string `json:"ledtaillight"`
		Lightsteeringassist             string `json:"lightsteeringassist"`
		Leddaytimerunninglight          string `json:"leddaytimerunninglight"`
	} `json:"light"`
	Internalconfig struct {
		Steeringwheelbeforeadjustment string `json:"steeringwheelbeforeadjustment"`
		Steeringwheelupadjustment     string `json:"steeringwheelupadjustment"`
		Steeringwheeladjustmentmode   string `json:"steeringwheeladjustmentmode"`
		Steeringwheelmemory           string `json:"steeringwheelmemory"`
		Steeringwheelmaterial         string `json:"steeringwheelmaterial"`
		Steeringwheelmultifunction    string `json:"steeringwheelmultifunction"`
		Steeringwheelheating          string `json:"steeringwheelheating"`
		Computerscreen                string `json:"computerscreen"`
		Interiorcolor                 string `json:"interiorcolor"`
		Rearcupholder                 string `json:"rearcupholder"`
		Supplyvoltage                 string `json:"supplyvoltage"`
		Interiormaterial              string `json:"interiormaterial"`
		Steeringwheelshift            string `json:"steeringwheelshift"`
		Activenoisereduction          string `json:"activenoisereduction"`
	} `json:"internalconfig"`
	Seat struct {
		Sportseat                           string `json:"sportseat"`
		Seatmaterial                        string `json:"seatmaterial"`
		Seatheightadjustment                string `json:"seatheightadjustment"`
		Driverseatadjustmentmode            string `json:"driverseatadjustmentmode"`
		Auxiliaryseatadjustmentmode         string `json:"auxiliaryseatadjustmentmode"`
		Driverseatlumbarsupportadjustment   string `json:"driverseatlumbarsupportadjustment"`
		Driverseatshouldersupportadjustment string `json:"driverseatshouldersupportadjustment"`
		Frontseatheadrestadjustment         string `json:"frontseatheadrestadjustment"`
		Rearseatadjustmentmode              string `json:"rearseatadjustmentmode"`
		Rearseatreclineproportion           string `json:"rearseatreclineproportion"`
		Rearseatangleadjustment             string `json:"rearseatangleadjustment"`
		Frontseatcenterarmrest              string `json:"frontseatcenterarmrest"`
		Rearseatcenterarmrest               string `json:"rearseatcenterarmrest"`
		Seatventilation                     string `json:"seatventilation"`
		Seatheating                         string `json:"seatheating"`
		Seatmassage                         string `json:"seatmassage"`
		Electricseatmemory                  string `json:"electricseatmemory"`
		Childseatfixdevice                  string `json:"childseatfixdevice"`
		Thirdrowseat                        string `json:"thirdrowseat"`
		Driverseatelectricadjustment        string `json:"driverseatelectricadjustment"`
		Auxiliaryseatelectricadjustment     string `json:"auxiliaryseatelectricadjustment"`
		Secondrowseatelectricadjustment     string `json:"secondrowseatelectricadjustment"`
		Frontseatfunction                   string `json:"frontseatfunction"`
		Rearseatfunction                    string `json:"rearseatfunction"`
		Secondrowseatadjustment             string `json:"secondrowseatadjustment"`
	} `json:"seat"`
	Entcom struct {
		Locationservice        string `json:"locationservice"`
		Bluetooth              string `json:"bluetooth"`
		Externalaudiointerface string `json:"externalaudiointerface"`
		Builtinharddisk        string `json:"builtinharddisk"`
		Cartv                  string `json:"cartv"`
		Speakernum             int    `json:"speakernum"`
		Audiobrand             string `json:"audiobrand"`
		Dvd                    string `json:"dvd"`
		Cd                     string `json:"cd"`
		Consolelcdscreen       string `json:"consolelcdscreen"`
		Rearlcdscreen          string `json:"rearlcdscreen"`
		Lcdscreensize          string `json:"lcdscreensize"`
		Fulllcddashboard       string `json:"fulllcddashboard"`
		Huddisplay             string `json:"huddisplay"`
		Roadrescue             string `json:"roadrescue"`
		FourG                  string `json:"4g"`
		Carapp                 string `json:"carapp"`
		Voicecontrol           string `json:"voicecontrol"`
		Phoneconnect           string `json:"phoneconnect"`
		Wirelesscharge         string `json:"wirelesscharge"`
		Gesturecontrol         string `json:"gesturecontrol"`
		Drivingrecorder        string `json:"drivingrecorder"`
	} `json:"entcom"`
	Aircondrefrigerator struct {
		Airconditioningcontrolmode string `json:"airconditioningcontrolmode"`
		Tempzonecontrol            string `json:"tempzonecontrol"`
		Rearairconditioning        string `json:"rearairconditioning"`
		Reardischargeoutlet        string `json:"reardischargeoutlet"`
		Airconditioning            string `json:"airconditioning"`
		Airpurifyingdevice         string `json:"airpurifyingdevice"`
		Carrefrigerator            string `json:"carrefrigerator"`
		Frontairconditioning       string `json:"frontairconditioning"`
		Fragrance                  string `json:"fragrance"`
	} `json:"aircondrefrigerator"`
	Actualtest struct {
		Accelerationtime100 string `json:"accelerationtime100"`
		Brakingdistance     string `json:"brakingdistance"`
	} `json:"actualtest"`
}

func (j *JisuAPI) QueryCarInfoByID(id string, c *gin.Context) (info CarDetailInfo, err error) {
	if id == "" {
		return
	}

	url := "https://api.jisuapi.com/car/detail?appkey=" + j.AppKey + "&carid=" + id

	var carInfo CarDetailRes
	res, err := netx.NewRetryClient().Get(url, c)
	if err != nil {
		return
	}

	if err = json.Unmarshal([]byte(res), &carInfo); err != nil {
		return
	}

	return carInfo.Result, nil
}
