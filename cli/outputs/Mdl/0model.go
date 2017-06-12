package Mdl

import "time"

type  struct{
	CompanyId int `json:"companyId" mysql:"company_id"`
	CompanyName string `json:"companyName" mysql:"company_name"`
	IsEnable int `json:"isEnable" mysql:"isEnable"`
	Address string `json:"address" mysql:"address"`
	SuburbDistrict string `json:"suburbDistrict" mysql:"suburb_district"`
	Ward string `json:"ward" mysql:"ward"`
	Postcode string `json:"postcode" mysql:"postcode"`
	StateProvince string `json:"stateProvince" mysql:"state_province"`
	Country string `json:"country" mysql:"country"`
	Description string `json:"description" mysql:"description"`
	Policy string `json:"policy" mysql:"policy"`
	ConditionToBook string `json:"conditionToBook" mysql:"condition_to_book"`
	LogoPath string `json:"logoPath" mysql:"logo_path"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	}

type s []*