/*
@Time : 2019/10/19 16:18
@Author : Hermes
@File : bench_test
@Description:
*/
package mathEngine

import "testing"

/*性能测试报告
个人配置如下i5-6400
BenchmarkAddShort-4       	 1683066	       686 ns/op
BenchmarkAddLong-4        	  332305	      3595 ns/op
BenchmarkComplexShort-4   	  534854	      2230 ns/op
BenchmarkComplexLong-4    	  273314	      4399 ns/op
性能可以说非常优秀
*/

func BenchmarkAddShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseAndExec("111+222")
	}
}
func BenchmarkAddLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseAndExec("111+222+3+4+5+66+89+556+564+46")
	}
}
func BenchmarkComplexShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseAndExec("11+22*3/(4+5)")
	}
}
func BenchmarkComplexLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseAndExec("111+(222/3+4)+5*(66+89)+556/(564*46)")
	}
}
