package main

import "fmt"

type Author struct {
Name string
VIP bool
Icon string
Signature string
Focus int
}

func main() {
a:=Author{
Name: "高数叔的百宝箱",
VIP:true,
Icon: "gaoshu",
Signature: "hanshujixian",
Focus: 33800,
}
 fmt.Println(a)
}