package file

import (
	"fmt"
	"sort"
	"testing"
)

var source = map[string]map[string]string{
	"mini": {
		"shareholder":            "shareholder",
		"KeyPerson":              "KeyPerson",
		"company_base_info":      "entbaseInfo",
		"branch":                 "branch",
		"alter":                  "alter",
		"spotCheckInfo":          "spotCheckInfo",
		"licenceinfoDetail":      "licenceinfoDetail",
		"insLicenceinfo":         "insLicenceinfo",
		"punishmentdetail":       "punishmentdetail",
		"illInfo":                "illInfo",
		"liquidation":            "liquidation",
		"insinvinfo":             "insinvinfo",
		"insAlterstockinfo":      "insAlterstockinfo",
		"entBusExcep":            "entBusExcep",
		"stakqualitinfo":         "stakqualitinfo",
		"mortreginfo":            "mortreginfo",
		"assistInfo":             "assistInfo",
		"insProPledgeRegInfo":    "insProPledgeRegInfo",
		"report_base":            "report_base",
		"report_webinfo":         "report_webinfo",
		"report_investor":        "report_investor",
		"report_investment":      "report_investment",
		"report_change":          "report_change",
		"report_social_security": "report_social_security",
		"report_guarantee":       "report_guarantee",
		"report_alter_stockinfo": "report_alter_stockinfo",
	},
	"pc": {
		"shareholder":            "shareholder",
		"KeyPerson":              "KeyPerson",
		"company_base_info":      "base",
		"shareholderDetail":      "shareholderDetail",
		"branch":                 "branch",
		"alter":                  "alter",
		"spotCheckInfo":          "spotCheckInfo",
		"licenceinfoDetail":      "licenceinfoDetail",
		"insLicenceinfo":         "insLicenceinfo",
		"punishmentdetail":       "punishmentdetail",
		"illInfo":                "illInfo",
		"liquidation":            "liquidation",
		"insInvinfo":             "insInvinfo",
		"insAlterstockinfo":      "insAlterstockinfo",
		"entBusExcep":            "entBusExcep",
		"stakqualitinfo":         "stakqualitinfo",
		"mortreginfo":            "mortreginfo",
		"mortregpersoninfo":      "mortregpersoninfo",
		"mortCreditorRightInfo":  "mortCreditorRightInfo",
		"mortGuaranteeInfo":      "mortGuaranteeInfo",
		"getMortAltItemInfo":     "getMortAltItemInfo",
		"getMortRegCancelInfo":   "getMortRegCancelInfo",
		"assistInfo":             "assistInfo",
		"assistInfoDetail":       "assistInfoDetail",
		"insProPledgeRegInfo":    "insProPledgeRegInfo",
		"report_base":            "report_base",
		"report_webinfo":         "report_webinfo",
		"report_investor":        "report_investor",
		"report_investment":      "report_investment",
		"report_change":          "report_change",
		"report_social_security": "report_social_security",
		"report_guarantee":       "report_guarantee",
		"report_alter_stockinfo": "report_alter_stockinfo",
	},
	"diqu": {
		"shareholder":            "shareholder",
		"KeyPerson":              "KeyPerson",
		"company_base_info":      "base",
		"shareholderDetail":      "shareholderDetail",
		"branch":                 "branch",
		"alter":                  "alter",
		"spotCheckInfo":          "spotCheckInfo",
		"licenceinfoDetail":      "licenceinfoDetail",
		"insLicenceinfo":         "insLicenceinfo",
		"punishmentdetail":       "punishmentdetail",
		"illInfo":                "illInfo",
		"liquidation":            "liquidation",
		"insInvinfo":             "insInvinfo",
		"insAlterstockinfo":      "insAlterstockinfo",
		"entBusExcep":            "entBusExcep",
		"stakqualitinfo":         "stakqualitinfo",
		"mortreginfo":            "mortreginfo",
		"mortregpersoninfo":      "mortregpersoninfo",
		"mortCreditorRightInfo":  "mortCreditorRightInfo",
		"mortGuaranteeInfo":      "mortGuaranteeInfo",
		"getMortAltItemInfo":     "getMortAltItemInfo",
		"getMortRegCancelInfo":   "getMortRegCancelInfo",
		"assistInfo":             "assistInfo",
		"assistInfoDetail":       "assistInfoDetail",
		"insProPledgeRegInfo":    "insProPledgeRegInfo",
		"report_base":            "report_base",
		"report_webinfo":         "report_webinfo",
		"report_investor":        "report_investor",
		"report_investment":      "report_investment",
		"report_change":          "report_change",
		"report_social_security": "report_social_security",
		"report_guarantee":       "report_guarantee",
		"report_alter_stockinfo": "report_alter_stockinfo",
	},
}

func TestFileExist(t *testing.T) {
	path := "./a.txt"
	f, err := Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b []byte
	n, err := f.Read(b)
	if err != nil {
		fmt.Println(err, n)
	}
	defer f.Close()
}

func TestFindNotExistKeys(t *testing.T) {
	//path := "/home/learnGoroutine/file/src.log"
	path := "/home/learnGoroutine/file/gongshang_pc_test.log"
	var keys []string
	for _, v := range source["pc"] {
		keys = append(keys, v)
	}
	//sort.Strings(keys)
	output, err := FindNotExistKeys(path, keys)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 对字符串列表进行排序, 从而保证输出结果总是一样的
	sort.Strings(output)
	fmt.Println(output)
}

func BenchmarkFindNotExistKeys(b *testing.B) {
	path := "/home/learnGoroutine/file/gongshang_pc_test.log"
	var keys []string
	for _, v := range source["pc"] {
		keys = append(keys, v)
	}
	for i := 0; i < b.N; i++ {
		_, err := FindNotExistKeys(path, keys)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
