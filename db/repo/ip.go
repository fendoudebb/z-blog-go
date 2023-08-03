package repo

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
	"z-blog-go/db"
)

type Ip struct {
	Id        int
	Ip        net.IP
	Country   string
	Region    string
	City      string
	Isp       string
	CountryId string
	RegionId  string
	CityId    string
	IspId     string
	CreateTs  time.Time
	UpdateTs  time.Time
}

type IpAddress struct {
	Id      int
	Address string
}

type TaobaoRes struct {
	Code string
	Data TaobaoData
}

type TaobaoData struct {
	QueryIp      *string `json:"QUERY_IP"`
	IspCode      *string `json:"ISP_CODE"`
	IspCn        *string `json:"ISP_CN"`
	IspEn        *string `json:"ISP_EN"`
	Asn          *string `json:"ASN"`
	CountryCode  *string `json:"COUNTRY_CODE"`
	CountryCn    *string `json:"COUNTRY_CN"`
	CountryEn    *string `json:"COUNTRY_EN"`
	ProvinceCode *string `json:"PROVINCE_CODE"`
	ProvinceCn   *string `json:"PROVINCE_CN"`
	ProvinceEn   *string `json:"PROVINCE_EN"`
	CityCode     *string `json:"CITY_CODE"`
	CityCn       *string `json:"CITY_CN"`
	CityEn       *string `json:"CITY_EN"`
	AreaCode     *string `json:"AREA_CODE"`
	AreaCn       *string `json:"AREA_CN"`
	AreaEn       *string `json:"AREA_EN"`
	CountyCode   *string `json:"COUNTY_CODE"`
	CountyCn     *string `json:"COUNTY_CN"`
	CountyEn     *string `json:"COUNTY_EN"`
	Longitude    *string `json:"LONGITUDE"`
	Latitude     *string `json:"LATITUDE"`
	Tzone        *string `json:"TZONE"`
}

func QueryAddress(ip string) *IpAddress {
	ipAddress := new(IpAddress)
	row := db.DB.QueryRow("select id, concat(country, region, city, isp) as address from ip_pool where ip=$1::inet limit 1", ip)
	_ = row.Scan(&ipAddress.Id, &ipAddress.Address)
	if ipAddress.Address != "" {
		return ipAddress
	}
	resp, err := http.Get("https://ip.taobao.com/getIpInfo.php?ip=" + ip)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != 200 {
		log.Println("QueryAddress Http Get Request err:", ip, err, resp)
		if ipAddress.Id == 0 {
			err = db.DB.QueryRow("insert into ip_pool(ip) values($1::inet) on conflict(ip) do update set update_ts = current_timestamp returning id", ip).Scan(&ipAddress.Id)
		}
		_, err = db.DB.Exec(`insert into ip_unknown(ip) values($1) on conflict(ip) do update set update_ts = current_timestamp`, ip)
		if err != nil {
			log.Println("QueryAddress With Insert err:", ip, err)
		}
		return ipAddress
	}
	//{"code":"0","data":{"CITY_EN":"Shenzhen","QUERY_IP":"113.87.160.228","CITY_CODE":"440300","CITY_CN":"深圳","COUNTY_EN":"null","LONGITUDE":"","PROVINCE_CN":"广东","TZONE":"","PROVINCE_EN":"Guangdong","ISP_EN":"China-Telecom","AREA_CODE":"80","PROVINCE_CODE":"440000","ISP_CN":"电信","AREA_CN":"华南","COUNTRY_CN":"中国","AREA_EN":"HuaNan","COUNTRY_EN":"China","COUNTY_CN":"null","COUNTY_CODE":"null","ASN":"null","LATITUDE":"","COUNTRY_CODE":"CN","ISP_CODE":"100017"}}
	//{"code":"0","data":{"CITY_EN":"xx","QUERY_IP":"46.219.210.147","CITY_CODE":"xx","CITY_CN":"XX","COUNTY_EN":"null","LONGITUDE":"","PROVINCE_CN":"XX","TZONE":"","PROVINCE_EN":"xx","ISP_EN":"xx","AREA_CODE":"xx","PROVINCE_CODE":"xx","ISP_CN":"XX","AREA_CN":"XX","COUNTRY_CN":"乌克兰","AREA_EN":"xx","COUNTRY_EN":"Ukraine","COUNTY_CN":"null","COUNTY_CODE":"null","ASN":"null","LATITUDE":"","COUNTRY_CODE":"UA","ISP_CODE":"xx"}}
	taobaoRes := new(TaobaoRes)
	err = json.NewDecoder(resp.Body).Decode(taobaoRes)
	if err != nil {
		log.Println("QueryAddress Http Get Json Decode err:", ip, err)
		return ipAddress
	}
	data := taobaoRes.Data
	if *data.CountryCn == "XX" {
		data.CountryCn = nil
	}
	if *data.ProvinceCn == *data.CityCn {
		data.CityCn = nil
	}
	if *data.ProvinceCn == "XX" {
		data.ProvinceCn = nil
	}
	if *data.CityCn == "XX" {
		data.CityCn = nil
	}
	if *data.IspCn == "XX" {
		data.IspCn = nil
	}
	if *data.CountryCode == "xx" {
		data.CountryCode = nil
	}
	if *data.ProvinceCode == "xx" {
		data.ProvinceCode = nil
	}
	if *data.CityCode == "xx" {
		data.CityCode = nil
	}
	if *data.IspCode == "xx" {
		data.IspCode = nil
	}

	var e error
	if ipAddress.Id > 0 {
		e = db.DB.QueryRow(`update ip_pool set country=$1, region=$2, city=$3, isp=$4, country_id=$5, region_id=$6, city_id=$7, isp_id=$8, update_ts=current_timestamp where id=$8 returning id, concat(country, region, city, isp) as address`,
			data.CountryCn, data.ProvinceCn, data.CityCn, data.IspCn, data.CountryCode, data.ProvinceCode, data.CityCode, data.IspCode, ipAddress.Id).Scan(&ipAddress.Id, &ipAddress.Address)
	} else {
		e = db.DB.QueryRow(`insert into ip_pool(ip, country, region, city, isp, country_id, region_id, city_id, isp_id) values($1::inet, $2, $3, $4, $5, $6, $7,$8, $9) returning id, concat(country, region, city, isp) as address`,
			data.QueryIp, data.CountryCn, data.ProvinceCn, data.CityCn, data.IspCn, data.CountryCode, data.ProvinceCode, data.CityCode, data.IspCode).Scan(&ipAddress.Id, &ipAddress.Address)
	}
	if e != nil {
		log.Println("QueryAddress Insert or Update err:", ip, e)
		return QueryAddress(ip)
	}
	log.Println("QueryAddress Id:", ipAddress.Id, ip)
	return ipAddress
}

func QueryUnknownIp() {
	var ip string
	err := db.DB.QueryRow("select ip from ip_unknown limit 1").Scan(&ip)
	if err != nil {
		return
	}
	address := QueryAddress(ip)
	if address.Id > 0 {
		_, _ = db.DB.Exec("delete from ip_unknown where ip = $1", ip)
	}
}

func CountIps() int {
	var count int
	err := db.DB.QueryRow("select count(*) from ip_pool").Scan(&count)
	if err != nil {
		log.Println("CountIps err:", err)
	}
	return count
}
