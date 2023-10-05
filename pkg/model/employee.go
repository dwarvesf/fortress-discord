package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Employee struct {
	ID               string      `json:"id"`
	CreatedAt        time.Time   `json:"createdAt"`
	UpdatedAt        time.Time   `json:"updatedAt"`
	FullName         string      `json:"fullName"`
	DisplayName      string      `json:"displayName"`
	TeamEmail        string      `json:"teamEmail"`
	PersonalEmail    string      `json:"personalEmail"`
	Avatar           string      `json:"avatar"`
	PhoneNumber      string      `json:"phoneNumber"`
	Address          string      `json:"address"`
	PlaceOfResidence string      `json:"placeOfResidence"`
	Country          string      `json:"country"`
	City             string      `json:"city"`
	MBTI             string      `json:"mbti"`
	Gender           string      `json:"gender"`
	Horoscope        string      `json:"horoscope"`
	Birthday         time.Time   `json:"birthday"`
	Username         string      `json:"username"`
	GithubID         string      `json:"githubID"`
	NotionID         string      `json:"notionID"`
	NotionName       string      `json:"notionName"`
	DiscordID        string      `json:"discordID"`
	DiscordName      string      `json:"discordName"`
	LinkedInName     string      `json:"linkedInName"`
	Status           string      `json:"status"`
	JoinedDate       time.Time   `json:"joinedDate"`
	LeftDate         interface{} `json:"leftDate"`
	Seniority        struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Name      string    `json:"name"`
		Code      string    `json:"code"`
		Level     int       `json:"level"`
	} `json:"seniority"`
	Positions []struct {
		Id   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"positions"`
	Stacks []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Avatar string `json:"avatar"`
	} `json:"stacks"`
	Projects []struct {
		Id             string        `json:"id"`
		Name           string        `json:"name"`
		DeploymentType string        `json:"deploymentType"`
		Status         string        `json:"status"`
		Positions      []interface{} `json:"positions"`
		Code           string        `json:"code"`
		Avatar         string        `json:"avatar"`
		StartDate      interface{}   `json:"startDate"`
		EndDate        interface{}   `json:"endDate"`
	} `json:"projects"`
	MmaScore struct {
		MasteryScore  string    `json:"masteryScore"`
		AutonomyScore string    `json:"autonomyScore"`
		MeaningScore  string    `json:"meaningScore"`
		RatedAt       time.Time `json:"ratedAt"`
	} `json:"mmaScore"`
}

type FortressEmployeeResponse struct {
	Data Employee `json:"data"`
}

type EmployeeMMAScore struct {
	EmployeeID    string          `json:"employeeID"`
	FullName      string          `json:"fullName"`
	MMAID         string          `json:"mmaID"`
	MasteryScore  decimal.Decimal `json:"masteryScore"`
	AutonomyScore decimal.Decimal `json:"autonomyScore"`
	MeaningScore  decimal.Decimal `json:"meaningScore"`
	RatedAt       *time.Time      `json:"ratedAt"`
}

type FortressEmployeeMMAScoreResponse struct {
	Data []EmployeeMMAScore `json:"data"`
}
