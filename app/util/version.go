package util

import "strings"

// 该函数比较两个版本号是否相等，是否大于或小于的关系
// 返回值：0表示v1与v2相等；1表示v1大于v2；2表示v1小于v2
func Compare(v1, v2 string) int {
	// 替换一些常见的版本符号
	replaceMap := map[string]string{"V": "", "v": "", "-": "."}
	for k, v := range replaceMap {
		if strings.Contains(v1, k) {
			strings.Replace(v1, k, v, -1)
		}
		if strings.Contains(v2, k) {
			strings.Replace(v2, k, v, -1)
		}
	}
	ver1 := strings.Split(v1, ".")
	ver2 := strings.Split(v2, ".")
	// 找出v1和v2哪一个最短
	var shorter int
	if len(ver1) > len(ver2) {
		shorter = len(ver2)
	} else {
		shorter = len(ver1)
	}
	// 循环比较
	for i := 0; i < shorter; i++ {
		if ver1[i] == ver2[i] {
			if shorter-1 == i {
				if len(ver1) == len(ver2) {
					return 0
				} else {
					if len(ver1) > len(ver2) {
						return 1
					} else {
						return 2
					}
				}
			}
		} else if ver1[i] > ver2[i] {
			return 1
		} else {
			return 2
		}
	}
	return -1
}

func VersionCompare(v1, v2, operator string) bool {
	com := Compare(v1, v2)
	switch operator {
	case "==":
		if com == 0 {
			return true
		}
	case "<":
		if com == 2 {
			return true
		}
	case ">":
		if com == 1 {
			return true
		}
	case "<=":
		if com == 0 || com == 2 {
			return true
		}
	case ">=":
		if com == 0 || com == 1 {
			return true
		}
	}
	return false
}
