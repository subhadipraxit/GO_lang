package strutil


func Reverse( s string) string{


	str := ""
  for i:=len(s);i !=0 ; i--{

		str = str+string(s[i-1])
	}
	return str
}
