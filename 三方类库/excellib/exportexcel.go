package excellib

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"time"
)

func ExportExcel() {
	s := "date：" + time.Now().Format("2006-01-02 15:04:05") + "\n"
	filePath := "./1.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		s += `
-----------------------------------------------------
| rename excel to 1.xlsx and move to current direct |
-----------------------------------------------------
`
		fmt.Println(s)
		time.Sleep(1000 * time.Minute)
		return
	}
	rows, err := f.GetRows("Sheet1")
	m := make(map[int]int)
	num := 1
	flag := 0 //标识是第一次统计
	for i, row := range rows {
		if i <= 4 {
			continue
		}
		//只要包含over就退出
		for _, v := range row {
			if v == "over" {
				//补上上一次丢失的数据
				m[num] += 1
				total := 0
				for i := 0; i < 200; i++ {
					if item, ok := m[i]; ok {
						total += i * item
						s += fmt.Sprintf("%d %d total %d\n", i, item, item*i)
					}
				}
				//最终总数
				s += fmt.Sprintf("total:%+v\n", total)
				fmt.Println(s)
				time.Sleep(1000 * time.Minute)
				return
			}
		}
		if row[3] == "" {
			continue
		}
		if row[2] != "" {
			flag++
			if flag == 1 { //如果是第一次、直接跳出
				continue
			}
			m[num] += 1
			num = 1
		}
		if row[2] == "" {
			num++
		}
	}
}
