// CheckID_card project main.go
package main

/*
 * ai -> a1 , a2, a3, a4, a5, a6... a17 (a18 是校验码) 身份证前17位对应(ai)
 * wi -> 7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 (17位)
 *
 *  int  res = 0;
 *  for (i = 1; i < 17; i++) {
 *          res += (a[i] * w[i])
 *          }
 *     int  y = res % 11；
 *
 *
 * y 与 a18的对应关系
 *
 *  y    { 0, 1,  2,  3, 4, 5, 6, 7, 8, 9, 10}
 *  a18  { 1, 0, 'X', 9, 8, 7, 6, 5, 4, 3,  2 } -> vefiry[18] = { 1, 0, 'X', 9, 8, 7, 6, 5, 4, ,3, 2};
 */

/*
 * 添加检测身份证年月日是否有问题功能
func checkYMD(ymd string) string { // 19910821 -> 1991 ,08, 21
	var y, m, d string
	y = ymd[:4]
	m = ymd[4:6]
	d = ymd[6:8]

	type YearMonthDay struct{
		Year  string
		Month string
		Day   string
	}

	YMD YearMonthDay = new YearMonthDay
	YMD.Year  = y
	YMD.Month = m
	YMD.Day   = d

	fmt.Println(YMD)

	return ""

}

*/

import (
	"fmt"
	"strconv"
	//"os"
	//"strings"
)

func PrintDate(date string) (string, string, string) {
	y := date[:4]
	m := date[4:6]
	d := date[6:8]

	/*
		fmt.Printf("年 -> %s\n", y)
		fmt.Printf("月 -> %s\n", m)
		fmt.Printf("日 -> %s\n", d)
	*/

	return y, m, d
}

func IsLeapYear(y string) bool { //y == 2000, 2004
	//判断是否为闰年
	year, _ := strconv.Atoi(y)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

func CheckYMD(y, m, d string) (bool, string) {
	//检查年份，假设现在最大的时间为2015年，当超过这个时间点，就显示错误
	/* 月份最大为12， 日期最大为 31，
	如果是2月，最大为29，最小为28
	*/
	year, _ := strconv.Atoi(y)
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)

	if year > 2015 {
		return false, "out of year "
	}

	if month > 12 {
		return false, "out of month"
	}

	if IsLeapYear(y) { //如果返回true,即是闰年
		if month == 2 && day > 29 {
			return false, "闰年，但是日期错误"
		}
	} else {
		if month == 2 && day > 28 {
			return false, "2月份，日期不为28日"
		}
	}

	return true, " "

}
func byte2int(x string) int {
	if x == "X" {
		return 88
	}

	res, _ := strconv.Atoi(x)

	return res
}

func check_id(id string) int { // len(id)= 17
	arry := make([]int, 17)

	//强制类型转换，将[]byte转换成[]int ,变化过程
	// []byte -> byte -> string -> int
	//将通过range 将[]byte转换成单个byte,再用强制类型转换string()，将byte转换成string
	//再通过strconv.Atoi()将string 转换成int 类型
	for i := 0; i < 17; i++ {
		arry[i], _ = strconv.Atoi(string(id[i]))
	}
	/*
		for k, v := range id {
			arry[k], _ = strconv.Atoi(string(v))
		}
	*/

	/*
		for p := 0; p < len(arry); p++ {
			fmt.Println("arry[", p, "]", "=", arry[p])
		}
	*/

	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		//fmt.Println("id =", i, byte2int(id[i]), wi[i])
		res += arry[i] * wi[i]
	}

	//fmt.Println("res = ", res)

	return (res % 11)
}

func verify_id(verify int, id_v int) (bool, string) {
	var temp int
	var i int
	a18 := [11]int{1, 0, 88 /* 'X' */, 9, 8, 7, 6, 5, 4, 3, 2}

	for i = 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			//fmt.Println("verify_id id",)
			// if a18[i] == 'X' ,let convert it to type string
			if a18[i] == 88 {
				fmt.Println("计算得到身份证最后一位是 ", string(a18[i]))
			} else {
				fmt.Println("计算得到身份证最后一位是 ", a18[i])
			}
			//fmt.Println(i, temp)
			break
		}
	}
	//if id_v == 'X', let's convert it to type string
	if id_v == 88 {
		fmt.Println("身份证最后一位是  X ")
	} else {
		fmt.Println("身份证最后一位是  ", id_v) // id_v是身份证的最后一位
	}

	if temp == id_v {

		return true, "验证成功"
	}

	return false, "验证失败"
}

func usage() {
	fmt.Println("请输入要检查的身份证号码18位\n")
}

func main() {
	/*
		var id_card [18]byte // 'X' == byte(88)， 'X'在byte中表示为88
		var id_card_copy [17]byte
	*/
	var id_card_string string
	fmt.Scanf("%s", &id_card_string)
	fmt.Printf("身份证号码是 = %s\n", id_card_string)
	//fmt.Println("id_card_string len = ", len(id_card_string))

	if len(id_card_string) < 18 {
		panic("必须要输入18位的身份证号码")
	}

	// 将字符串，转换成[]byte,并保存到id_card[]数组当中
	/*
		for k, v := range []byte(id_card_string) {
			id_card[k] = byte(v)
		}
	*/

	YearMonthDay := id_card_string[6:14]
	fmt.Println("身份证包含的日期: ", YearMonthDay)
	PrintDate(YearMonthDay)

	//复制id_card[18]前17位元素到id_card_copy[]数组当中
	/*
		for j := 0; j < 17; j++ {
			id_card_copy[j] = id_card[j]

			//fmt.Println(byte2int(id_card[j]))
		}
	*/

	/*
		fmt.Println(byte2int(id_card[17]))
		fmt.Println(string(id_card[17]))
	*/
	/*
		y := check_id(id_card_copy)
		fmt.Println(y)
	*/

	CheckYMD(PrintDate(YearMonthDay))
	fmt.Println(verify_id(check_id(id_card_string[:17]), byte2int(id_card_string[17:])))
}

//测试身份证号码：34052419800101001X
//测试身份证号码：511028199507215915
