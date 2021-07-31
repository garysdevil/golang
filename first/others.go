// 求n次方 math.Pow(num,power)
func Exponent(num,power int)int{
    var result int
    if num==0{
        return 0
    }
    if power==0{
        return 1
    }
    for i:=0; i<power-1; i++{
        result = num*num
    }
    return result
}