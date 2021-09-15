package main

import (
	"Utils"
	"bufio"
	_ "encoding/binary"
	_ "encoding/gob"
	"fmt"
	bitset "github.com/bits-and-blooms/bitset"
	"io"
	"math"
	"os"
	"strconv"
	_ "strconv"
)
type Node struct {
	m_low int
	m_high int
	m_content string
	m_type int
}
func bitCompute (bstr string,size int) bitset.BitSet{
	bint, _ :=strconv.Atoi(bstr)
	bturn := bint << size
	buf:=[]uint64{uint64(bturn)}
	newb := bitset.From(buf)
	return *newb
}
var delta = Utils.Stack{}

func backtrackingPath(basic string,s string)  {
	m := len(basic)
	n := len(s)
	var dp = make([][]int , m+1)
	for i :=0 ; i < m+1 ; i++{
		dp[i] = make([]int , n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if basic[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = int (math.Min(math.Min(float64(dp[i-1][j]+1),float64(dp[i][j-1]+1)),float64(dp[i-1][j-1]+1)))
			}
		}
	}
	fmt.Printf("distance: %d\n",dp[m][n])

	for n >= 0 || m >= 0{
		if n!=0 && dp[m][n-1]+1 == dp[m][n] {
			var a Node
			a.m_content = string(s[n-1])
			a.m_low = m-1
			a.m_high = m
			a.m_type = 0
			if delta.Size()==0 {
				delta.Push(a)
			}else {
				b := delta.Top().(Node)
				if b.m_type==a.m_type && b.m_low==a.m_low{
					a.m_content=a.m_content + b.m_content
					delta.Pop()
				}
				delta.Push(a)
			}
			fmt.Printf("insert %s at %d\n",string(s[n-1]),m-1)
			n = n - 1
			continue
		}else if m!=0 && dp[m-1][n]+1 == dp[m][n] {
			fmt.Printf("delete %s at %d\n",string(basic[m-1]),m-1)
			var a Node
			a.m_content = "-"
			a.m_low = m-1
			a.m_high = m-1
			a.m_type = 1
			if delta.Size()==0 {
				delta.Push(a)
			}else {
				b := delta.Top().(Node)
				if b.m_type==a.m_type && b.m_low==a.m_low+1 {
					a.m_high = b.m_high
					delta.Pop()
				}
				delta.Push(a)
			}
			m = m - 1
			continue
		}else if m==0&&n==0 {

		}else if dp[m-1][n-1]+1 == dp[m][n] {
			fmt.Printf("replace %s to %s at %d\n",string(basic[m-1]),string(s[n-1]),m-1)
			var a Node
			a.m_content = string(s[n-1])
			a.m_low = m-1
			a.m_high = m-1
			a.m_type = 2
			if delta.Size() == 0 {
				delta.Push(a)
			}else {
				b := delta.Top().(Node)
				if b.m_type==a.m_type && b.m_low==a.m_low+1 {
					a.m_high = b.m_high
					a.m_content = a.m_content + b.m_content
					delta.Pop()
				}
				delta.Push(a)
			}
			n = n-1
			m = m-1
			continue
		}
		n = n-1
		m = m-1
	}
}

func main()  {
	delta.InitStack()

	//获取参考串文件第一行
	file1,_ := os.Open("D:\\GO_CODE\\go1\\resources\\chr.txt")
	defer file1.Close()
	reader1 := bufio.NewReader(file1)
	basic,_ := reader1.ReadString('\n')
	//fmt.Println(basic)

	//获取原始文件行数
	size := 0
	file2,_ := os.Open("D:\\GO_CODE\\go1\\resources\\query.txt")
	defer file2.Close()
	reader2 := bufio.NewReader(file2)
	for  {
		_,_,err := reader2.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Printf("%s\n",s)
		//backtrackingPath(basic, string(s))
		size++
	}
	//fmt.Printf("%d\n",size)

	file3,err := os.Create("index.txt")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file3.Close()
	//writer := bufio.NewWriter(file3)
	size_log :=bitset.New(24)
	size_log.Set(uint(size))
	file3.WriteString(size_log.String())

	file4,_ := os.Open("D:\\GO_CODE\\go1\\resources\\query.txt")
	defer file4.Close()

	reader4 := bufio.NewReader(file4)
	for  {
		s,_,err := reader4.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Printf("%s\n",s)
		backtrackingPath(basic, string(s))

		size2 := delta.Size()
		size_delta :=bitset.New(16)
		size_delta.Set(uint(size2))
		file3.WriteString(size_delta.String())

		for delta.Size() != 0 {
			d := delta.Top().(Node)
			//fmt.Println(d)
			b := bitset.New(24)
			b.Set(uint(d.m_low))
			bstr := b.String()
			basic_b := bitCompute(bstr,11)
			hb :=bitset.New(24)
			hb.Set(uint(d.m_high))
			basic_bb := basic_b.SymmetricDifference(hb)
			basic_bbb := bitCompute(basic_bb.String(),2)
			tb :=bitset.New(24)
			tb.Set(uint(d.m_type))
			final_b := basic_bbb.SymmetricDifference(tb)
			file3.WriteString(final_b.String())
			content_size := len(d.m_content)
			bit_content := bitset.New(16)
			bit_content.Set(uint(content_size))
			file3.WriteString(bit_content.String())
			file3.WriteString(d.m_content+"\n")
			delta.Pop()
		}
	}
}
