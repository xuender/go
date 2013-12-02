package main
import (
  "fmt"
)
type House struct{
  位置 int // 1-5
  国籍 int // 0 未知
  饮料 int // 0 未知 
  香烟 int // 0 未知
  宠物 int // 0 未知
}
func 问题(){
  fmt.Println("爱因斯坦的问题")
  fmt.Println("1、在一条街上，有5座房子，喷了5种颜色。 ")
  fmt.Println("2、每个房里住着不同国籍的人 ")
  fmt.Println("3、每个人喝不同的饮料，抽不同品牌的香烟，养不同的宠物 ")
  fmt.Println("问题是：谁养鱼？ ")
  fmt.Println("提示： ")
  fmt.Println("1、英国人住红色房子 ")
  fmt.Println("2、瑞典人养狗 ")
  fmt.Println("3、丹麦人喝茶 ")
  fmt.Println("4、绿色房子在白色房子左面 ")
  fmt.Println("5、绿色房子主人喝咖啡 ")
  fmt.Println("6、抽Pall Mall 香烟的人养鸟 ")
  fmt.Println("7、黄色房子主人抽Dunhill 香烟 ")
  fmt.Println("8、住在中间房子的人喝牛奶 ")
  fmt.Println("9、 挪威人住第一间房 ")
  fmt.Println("10、抽Blends香烟的人住在养猫的人隔壁 ")
  fmt.Println("11、养马的人住抽Dunhill 香烟的人隔壁 ")
  fmt.Println("12、抽Blue Master的人喝啤酒 ")
  fmt.Println("13、德国人抽Prince香烟 ")
  fmt.Println("14、挪威人住蓝色房子隔壁 ")
  fmt.Println("15、抽Blends香烟的人有一个喝水的邻居")
}
func main() {
  问题()
}
