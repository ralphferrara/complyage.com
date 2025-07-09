package handlers

import (
	"encoding/binary"
	"net"
	"net/http"
	"os"
	"strings"

	"api/client"
	"api/db"
	"api/geo"
	"api/keys"
	"api/models"
	"api/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| In Memory
//||------------------------------------------------------------------------------------------------||

var UseInMemory = os.Getenv("ENV_MODE") == "production"

//||------------------------------------------------------------------------------------------------||
//|| Request Struct
//||------------------------------------------------------------------------------------------------||

type LookupRequest struct {
	APIKey    string `json:"api_key"`
	IPAddress string `json:"ip_address"`
}

//||------------------------------------------------------------------------------------------------||
//|| Response Struct
//||------------------------------------------------------------------------------------------------||

type LookupResponse struct {
	Requirements string `json:"requirements"`
	Country      string `json:"country"`
	State        string `json:"state"`
}

//||------------------------------------------------------------------------------------------------||
//|| Helper :: Convert IPv4 string to uint32
//||------------------------------------------------------------------------------------------------||

func ipToUint32(ipStr string) uint32 {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip)
}

//||------------------------------------------------------------------------------------------------||
//|| Handle Loading of IP ranges into memory
//||------------------------------------------------------------------------------------------------||

func LookupRequirements(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||
	apiKey := r.URL.Query().Get("api")
	if apiKey == "" {
		responses.Error(w, http.StatusUnauthorized, "API key is required")
		return
	}
	ipAddress := client.GetClientIP(r)
	//||------------------------------------------------------------------------------------------------||
	//|| Get the API Key Data
	//||------------------------------------------------------------------------------------------------||
	var api models.API
	if UseInMemory {
		var ok bool
		api, ok = keys.GetKey(apiKey)
		if !ok {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
	} else {
		if result := db.DB.Where("api_public = ?", apiKey).First(&api); result.Error != nil {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	ipNum := ipToUint32(ipAddress)
	if ipNum == 0 {
		http.Error(w, "Invalid IP address - "+ipAddress, http.StatusBadRequest)
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Country and State
	//||------------------------------------------------------------------------------------------------||
	var country, state string
	if UseInMemory {
		ipBlock, found := geo.FindIPRange(ipNum)
		if !found {
			http.Error(w, "IP not found", http.StatusNotFound)
			return
		}
		country = ipBlock.Country
		state = ipBlock.State
	} else {
		var ipRecord models.IP
		if result := db.DB.
			Where("start_ip <= ? AND end_ip >= ?", ipNum, ipNum).
			Order("start_ip DESC").
			Limit(1).
			First(&ipRecord); result.Error != nil {
			http.Error(w, "IP not found", http.StatusNotFound)
			return
		}
		country = ipRecord.Country
		state = ipRecord.State
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if in a restricted territory
	//||------------------------------------------------------------------------------------------------||
	allowed := false
	territories := strings.Split(api.APITerritories, ",")
	for _, t := range territories {
		t = strings.TrimSpace(t)
		if strings.EqualFold(t, country) || strings.EqualFold(t, state) {
			allowed = true
			break
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Not allowed
	//||------------------------------------------------------------------------------------------------||
	if !allowed {
		responses.Success(w, http.StatusOK, LookupResponse{
			Requirements: api.APIRequirements,
			Country:      country,
			State:        state,
		})
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Allowed
	//||------------------------------------------------------------------------------------------------||
	resp := LookupResponse{
		Requirements: api.APIRequirements,
		Country:      country,
		State:        state,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Give the data
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, resp)
}
