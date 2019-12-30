/*
mathEngine

--------

数学表达式计算引擎

使用 Go 实现的数学表达式解析计算引擎，无任何依赖，相对比较完整的完成了数学表达式解析执行，包括词法分析、语法分析、构建AST、运行。

能够处理的表达式样例：

1+127-21+(3-4)*6/2.5
(88+(1+8)*6)/2+99
123_345_456 * 1.5 - 2 ^ 4
-4 * 6 + 2e2 - 1.6e-3
sin(pi/2)+cos(45-45*1)+tan(pi/4)
99+abs(-1)-ceil(88.8)+floor(88.8)
max(min(2^3, 3^2), 10*1.5-7)

直接调用解析执行函数 :
func main() {
  s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
  // call top level function
  r, err := mathEngine.ParseAndExec(s)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%s = %v", s, r)
}

已实现
 加 +
 减 -
 乘 *
 除 /
 取余 %
 整数次方 ^
 科学计数法 e.g. 1.2e7、 1.2e-7
 括号 ()
 混合运算 e.g. 1+2*6/4+(456-8*9.2)-(2+4^5)*2e3+1.2e-2
 友好的长数字 e.g. 123_456_789
 三角函数 e.g. sin, cos, tan, cot, sec, csc
 常量 pi
 辅助函数 e.g. abs, ceil, floor, sqrt, cbrt
 友好的错误消息 e.g.
*/
package mathEngine
