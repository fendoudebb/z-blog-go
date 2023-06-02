package repo

import (
	"fmt"
	"net"
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

func QueryAddress(ip string) (IpAddress, error) {
	var ipAddress IpAddress
	row := db.DB.QueryRow("select id, concat(country, region, city, isp) as address from ip_pool where ip=$1::inet limit 1", ip)
	if err := row.Scan(&ipAddress.Id, &ipAddress.Address); err != nil {
		return ipAddress, fmt.Errorf("GetIpAddress %q: %v", ip, err)
	}
	return ipAddress, nil
}

func CountIps() (int, error) {
	var count int
	row := db.DB.QueryRow("select count(*) from ip_pool")
	if err := row.Scan(&count); err != nil {
		return count, fmt.Errorf("CountIps: %v", err)
	}
	return count, nil
}
