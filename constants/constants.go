package constants

const(
	CoWINBaseURL = "https://cdn-api.co-vin.in/api"
	GetSessionsByDistrictIdAndDate = CoWINBaseURL + "/v2/appointment/sessions/public/findByDistrict?"
	GetSessionsByPinIdAndDate = CoWINBaseURL + "/v2/appointment/sessions/public/findByPin?"
	StatesList = CoWINBaseURL+ "/v2/admin/location/states"
    DistrictList = CoWINBaseURL+ "/v2/admin/location/districts"
)
