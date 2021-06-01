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
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func TestFindNotExistKeys(t *testing.T) {
	//path := "/home/learnGoroutine/file/src.log"
	//path := "/home/learnGoroutine/file/gongshang_pc.log"
	//dataSource := "pc"
	path := "/home/learnGoroutine/file/gongshang_diqu.log"
	dataSource := "diqu"
	var keys []string
	for k := range source[dataSource] {
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	output, err := FindNotExistKeys(path, dataSource, keys)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 对字符串列表进行排序, 从而保证输出结果总是一样的
	sort.Strings(output)
	fmt.Println(output)
}

// go test -bench=. -run=none -benchmem -benchtime=3s
// -bench=.表示运行所有的基准测试，-bench=BenchmarkFindNotExistKeys表示只允许该基准函数
// -run=none表示不运行TestFindNotExistKeys函数，-benchtime=3s表示设置基准测试运行的时间（默认为1s）
// -benchmem表示显示memory指标（也就是再PASS这一行多显示 B/op 和 allocs/op）
// 15表示3s内运行的次数（b.N的数值），220301926 ns/op表示执行一次循环耗时220301926纳秒
// B/op 表示执行一次循环分配的内存数（字节，Byte）；allocs/op表示执行一次操作分配的内存次数
//
// goos表示操作系统，goarch表示平台的体系架构，pkg表示运行的文件所在的包
// BenchmarkFindNotExistKeys表示基准测试函数名
//
//goos: linux
//goarch: amd64
//pkg: GopherUtils/file
//BenchmarkFindNotExistKeys             15         220301926 ns/op         1485740 B/op       2629 allocs/op
//PASS
//ok      GopherUtils/file        3.525s
func BenchmarkFindNotExistKeys(b *testing.B) {
	path := "/home/learnGoroutine/file/gongshang_pc_test.log"
	var keys []string
	dataSource := "pc"
	for k := range source[dataSource] {
		keys = append(keys, k)
	}

	// 重置计时器，会将运行事件归零，然后从下面的代码运行开始重新计时;过滤掉之前运行的代码所消耗的时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := FindNotExistKeys(path, dataSource, keys)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
