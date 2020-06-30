package main

import "fmt"

func average(salary []int) float64 {
	l := len(salary)
	if l <= 2 {
		return 0
	}
	var max int
	var min int
	max = -1
	min = 100000009
	sum := 0
	for i := 0; i < l; i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		sum += salary[i]
	}
	return float64(sum-max-min) / float64(l-2)
}

func kthFactor(n int, k int) int {
	k--
	if k <= 0 {
		return 1
	}
	for i := 2; i <= n; i++ {
		flag := false
		if n%i == 0 {
			flag = true
			n = n / i
		}
		if flag {
			k--
		}
		if k == 0 {
			return i
		}
	}
	return -1
}

func longestSubarray(nums []int) int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	left[0] = 0
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}
	}
	right[l-1] = 0
	for i := l - 2; i >= 0; i-- {
		if nums[i+1] > 0 {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 0
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		if left[i]+right[i] > ans {
			ans = left[i] + right[i]
		}
	}
	return ans
}

func isPathCrossing(path string) bool {

	m:=make(map[string]bool)
	l := len(path)
	x,y:=0,0
	begin:=fmt.Sprintf("%v_%v",x,y)
	m[begin]=true
	for i := 0; i < l; i++ {
		if path[i]=='N' {
			y++
			str := fmt.Sprintf("%v_%v",x,y)
			_,ok := m[str]
			if ok{
				return true
			} else {
				m[str] = false
			}
		}
		if path[i]=='E' {
			x++
			str := fmt.Sprintf("%v_%v",x,y)
			_,ok := m[str]
			if ok{
				return true
			} else {
				m[str] = false
			}

		}
		if path[i]=='W' {
			x--
			str := fmt.Sprintf("%v_%v",x,y)
			_,ok := m[str]
			if ok{
				return true
			} else {
				m[str] = false
			}

		}
		if path[i]=='S' {
			y--
			str := fmt.Sprintf("%v_%v",x,y)
			_,ok := m[str]
			if ok{
				return true
			} else {
				m[str] = false
			}

		}
	}
	return false
}
func canArrange(arr []int, k int) bool {
	l := len(arr)
	mod := make([]int, l)
	for i := 0; i < l; i++ {
		mod[i] = arr[i]%k
	}
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		a := mod[i]
		b:= k-a
		b2 := 0-a
		b3 := -k-a
		count,ok := m[b]
		count2,ok2 := m[b2]
		count3,ok3 := m[b3]
		if ok  {
			if count==1{
				delete(m, b)
			}else {
				m[b] = count-1
			}
		} else if ok2 {
			if count2==1{
				delete(m,b2)
			}else {
				m[b2]=count2-1
			}
		}else if ok3 {
			if count3==1{
				delete(m,b3)
			}else {
				m[b3]=count3-1
			}
		} else {
			v,ok := m[a]
			if ok{
				m[a] = v+1
			}else {
				m[a] = 1
			}
		}
	}
	if len(m)==0 {
		return true
	}
	return false
}
/**
class Solution {
    int a[15],f[32768],o[32768];
public:
    int minNumberOfSemesters(int n, vector<vector<int>>& dependencies, int k) {
        memset(a,0,sizeof(a));
        int i,j,l;
        for(auto e:dependencies)a[e[1]-1]|=1<<e[0]-1;
        for(i=1;i<1<<n;i++)o[i]=o[i>>1]+(i&1);
        memset(f,127,sizeof(f));
        for(i=f[0]=0;i<1<<n;i++)if(f[i]<=n)
        {
            for(j=l=0;j<n;j++)if(!(i>>j&1)&&(a[j]&i)==a[j])l|=1<<j;
            for(j=l;j;j=j-1&l)if(o[j]<=k)f[i|j]=min(f[i|j],f[i]+1);
        }
        return f[(1<<n)-1];
    }
};
*/
// 状态压缩dp：整数代表集合 bit代表元素
func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	a:= make([]int, 15)
	f:= make([]int,32768)
	o:= make([]int,32768)
	for i := 0; i < len(dependencies); i++ {
		// a[节点] = 前依赖节点  ？超过一个
		a[dependencies[i][1]-1] = 1<<dependencies[i][0]-1
	}
	for i := 1; i < 1<<n; i++ {
		o[i] = o[i>>1] + (i&1)
	}
	for i := 0; i < len(f); i++ {
		f[i] = 127
	}
	f[0] = 0
	for i := 0; i < 1<<n; i++ {
		if f[i] < n {
			l:=0
			for j:=0;j<n;j++{
				if !(i>>j&1)&&(a[j]&i)==a[j] {
					l|=1<<j
				}
			}
			for j:=l;j>0;j=j-1&l{
				if o[j]<=k {
					f[i|j]=min(f[i|j],f[i]+1)
				}
			}
		}
	}
	return f[(1<<n)-1]
}
func minNumberOfSemesters2(n int, dependencies [][]int, k int) int {
	flag := make([]bool, n+1)
	before := make(map[int]map[int]bool)
	ld := len(dependencies)
	for i := 0; i < ld; i++ {
		value, ok := before[dependencies[i][1]]
		if ok {
			value[dependencies[i][0]] = true
		} else {
			tm := make(map[int]bool)
			tm[dependencies[i][0]] = true
			before[dependencies[i][1]] = tm
		}
	}
	ans := 0
	count := 0
	// 遍历每一层
	for count < n {
		temp := k
		// 遍历每一个节点 可选则选
		for i := 1; i <= n; i++ {
			value := before[i]
			// 找到一个
			if len(value) == 0 && flag[i] == false {
				temp--
				flag[i] = true
				for num := 1; num <= n; num++ {
					v := before[num]
					delete(v, num)
				}
				count++
			}
			if temp == 0 {
				break
			}
		}
		ans++
	}
	return ans
}
func main() {
	fmt.Print(canArrange([]int{-4,-7,5,2,9,1,10,4,-8,-3},3))
	//fmt.Print(isPathCrossing("NESWW"))
	//fmt.Println(longestSubarray([]int{1,1,0,1}))
	//fmt.Println(minNumberOfSemesters(4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2))
}
