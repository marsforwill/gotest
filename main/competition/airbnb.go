package main

import "fmt"

func validate_xml(xml string) string {
	// Write your code here
	var tags []string
	for i := 0; i < len(xml); i++ {
		if xml[i] == '>' {
			return "parse error"
		}
		if xml[i] == '<' {
			// start tag
			if i+1 < len(xml) && xml[i+1] != '/' {
				var j int
				for j = i + 1; j < len(xml) && xml[j] != '>'; j++ {
					if xml[j] == '<' {
						return "parse error"
					}
				}
				if j == len(xml) || j == i+1 {
					return "parse error"
				}
				startTag := xml[i+1 : j]
				tags = append(tags, startTag)
				i = j
			} else if i+1 < len(xml) && xml[i+1] == '/' {
				// end tag
				var j int
				for j = i + 2; j < len(xml) && xml[j] != '>'; j++ {
					if xml[j] == '<' {
						return "parse error"
					}
				}
				if j == len(xml) || j == i+2 {
					return "parse error"
				}
				endTag := xml[i+2 : j]
				if len(tags) > 0 && tags[len(tags)-1] == endTag {
					tags = tags[:len(tags)-1]
				} else {
					return "encountered closing tag without matching open tag for </" + endTag + ">"
				}
				i = j
			} else {
				return "parse error"
			}
		}
	}
	if len(tags) > 0 {
		return "missing closing tag for <" + tags[0] + ">"
	}
	return "valid"
}

//39. 组合总和
// 深度搜索回溯 可加剪枝优化
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var comb []int
	var dfsComb func(target int, idx int)
	dfsComb = func(target int, idx int) {
		// 递归出口
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//不选当前数
		dfsComb(target, idx+1)
		//选当前数
		if target >= candidates[idx] {
			comb = append(comb, candidates[idx])
			dfsComb(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfsComb(target, 0)
	return ans
}

func main() {
	fmt.Println(validate_xml("<a></a>"))
	//	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
