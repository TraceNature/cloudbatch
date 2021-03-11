package cmd

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/waf/models"
)

func TestExcelWaf2interface(t *testing.T) {
    domainConfig := &DomainsConfig{}
    domainConfig.Domains = make([]models.AddDomain, 0)
    f, err := excelize.OpenFile("E:\\User\\luxinyu1\\Desktop\\项目\\cloudbatch\\waf.xlsx")
    if err != nil {
        panic(err)
    }
    rows, err := f.GetRows("WAF")
    if err != nil {
        panic(err)
    }
    for index, row := range rows {
        if index == 0 {
            continue
        }
        domain := excelWaf2interface(row)
        if domain == nil {
            panic(domain)
        }
        domainConfig.Domains = append(domainConfig.Domains, *domain)
    }
    domainJsonStr, err := json.Marshal(domainConfig.Domains)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(domainJsonStr))
}