package main
import (
  "fmt"
  "github.com/ntns/goitertools/itertools"
)
const (
  颜色 = iota
  国籍
  宠物
  饮料
  香烟
)
const (
  德国 = iota
  英国
  瑞典
  丹麦
  挪威
)

var 国籍Lookup = map[int]string {
  德国: "德国",
  英国: "英国",
  瑞典: "瑞典",
  丹麦: "丹麦",
  挪威: "挪威",
}
const (
  蓝色 = iota
  红色
  绿色
  白色
  黄色
)
var 颜色Lookup = map[int]string {
  蓝色: "蓝色",
  红色: "红色",
  绿色: "绿色",
  白色: "白色",
  黄色: "黄色",
}
const (
  鱼 = iota
  狗
  鸟
  马
  猫
)

const (
  水 = iota
  茶
  咖啡
  牛奶
  啤酒
)
var 饮料Lookup = map[int]string {
  水: "水",
  茶: "茶",
  咖啡: "咖啡",
  牛奶: "牛奶",
  啤酒: "啤酒",
}
const (
  Prince = iota
  PallMall
  Dunhill
  BlueMaster
  Blends
)
/*
初始化
*/
func 初始化() [][]int{
  /*
  i 代表位置
  f 代表 0 颜色 1 国籍 2 饮料 3 香烟 4 宠物
  取值是 0 代表未知
  */
  var house [][]int
  house = make([][]int, 5)
  for i := 0; i < 5; i++{
    house[i] = make([]int, 5)
    for f := 0; f < 5; f++{
      house[i][f] = 0
    }
  }
  return house
}
// 1、英国人住红色房子
func check1(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][国籍] == 英国 && house[i][颜色] == 红色{
      return true
    }
  }
  return false
}
//  fmt.Println("2、瑞典人养狗 ")
func check2(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][国籍] == 瑞典 && house[i][宠物] == 狗{
      return true
    }
  }
  return false
}
//  fmt.Println("3、丹麦人喝茶 ")
func check3(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][国籍] == 丹麦 && house[i][饮料] == 茶{
      return true
    }
  }
  return false
}
//  fmt.Println("4、绿色房子在白色房子左面 ")
func check4(house [][]int) bool{
  var g, w = 0, 0
  for i := 0; i < 5; i++{
    if house[i][颜色] == 绿色{
      g = i
    }
    if house[i][颜色] == 白色{
      w = i
    }
  }
  //TODO 左侧有错，需要检查
  return g + 1 == w
}
//  fmt.Println("5、绿色房子主人喝咖啡 ")
func check5(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][颜色] == 绿色 && house[i][饮料] == 咖啡{
      return true
    }
  }
  return false
}
//  fmt.Println("6、抽Pall Mall 香烟的人养鸟 ")
func check6(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][香烟] == PallMall && house[i][宠物] == 鸟{
      return true
    }
  }
  return false
}
//  fmt.Println("7、黄色房子主人抽Dunhill 香烟 ")
func check7(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][香烟] == Dunhill && house[i][颜色] == 黄色{
      return true
    }
  }
  return false
}
//  fmt.Println("8、住在中间房子的人喝牛奶 ")
func check8(house [][]int) bool{
  return house[3][饮料] == 牛奶
}
//  fmt.Println("9、 挪威人住第一间房 ")
func check9(house [][]int) bool{
  return house[0][国籍] == 挪威
}
//  fmt.Println("10、抽Blends香烟的人住在养猫的人隔壁 ")
func check10(house [][]int) bool{
  var f = 0
  for i := 0; i < 5; i++{
    if house[i][香烟] == Blends{
      f = i
    }
  }
  var l = false
  if f > 0{
    l = house[f - 1][宠物] == 猫
  }
  var r = false
  if f < 4{
    r = house[f + 1][宠物] == 猫
  }
  return l || r
}
//  fmt.Println("11、养马的人住抽Dunhill 香烟的人隔壁 ")
func check11(house [][]int) bool{
  var f = 0
  for i := 0; i < 5; i++{
    if house[i][香烟] == Dunhill{
      f = i
    }
  }
  var l = false
  if f > 0{
    l = house[f - 1][宠物] == 马
  }
  var r = false
  if f < 4{
    r = house[f + 1][宠物] == 马
  }
  return l || r
}
//  fmt.Println("12、抽Blue Master的人喝啤酒 ")
func check12(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][香烟] == BlueMaster && house[i][饮料] == 啤酒{
      return true
    }
  }
  return false
}
//  fmt.Println("13、德国人抽Prince香烟 ")
func check13(house [][]int) bool{
  for i := 0; i < 5; i++{
    if house[i][香烟] == Prince && house[i][国籍] == 德国{
      return true
    }
  }
  return false
}
//  fmt.Println("14、挪威人住蓝色房子隔壁 ")
func check14(house [][]int) bool{
  return house[1][颜色] == 蓝色
}
//  fmt.Println("15、抽Blends香烟的人有一个喝水的邻居")
func check15(house [][]int) bool{
  var f = 0
  for i := 0; i < 5; i++{
    if house[i][香烟] == Blends{
      f = i
    }
  }
  var l = false
  if f > 0{
    l = house[f - 1][饮料] == 水
  }
  var r = false
  if f < 4{
    r = house[f + 1][饮料] == 水
  }
  return l || r
}
func check(house [][]int) bool{
  /*
  for i := 0; i < 5; i++{
    fmt.Println(i)
    for f := 0; f < 5; f++{
      fmt.Print(house[i][f])
    }
  }
  */
  if !check1(house){
    return false
  }
  if !check2(house){
    return false
  }
  if !check3(house){
    return false
  }
  if !check4(house){
    return false
  }
  if !check5(house){
    return false
  }
  if !check6(house){
    return false
  }
  if !check7(house){
    return false
  }
  if !check8(house){
    return false
  }
  if !check9(house){
    return false
  }
  if !check10(house){
    return false
  }
  if !check11(house){
    return false
  }
  if !check12(house){
    return false
  }
  if !check13(house){
    return false
  }
  if !check14(house){
    return false
  }
  if !check15(house){
    return false
  }
  fmt.Println("计算出结果..")
  for i := 0; i < 5; i++{
    if house[i][宠物] == 鱼{
      out(i, house[i])
    }
  }
  return true
}
func out(i int, house []int){
  fmt.Printf(
    "第%d个%s房子里面喝%s抽%d的%s人养鱼",
    i+1,
    颜色Lookup[house[颜色]],
    饮料Lookup[house[饮料]],
    house[香烟],
    国籍Lookup[house[国籍]],
  )
}
func run(){
  var house = 初始化()
  all := itertools.Permutations([]int{0, 1, 2, 3, 4}, 5)
  for _, g := range all{
    //fmt.Println("9、 挪威人住第一间房 ")
    if g[0] != 挪威{
      continue
    }
    for _, l := range all{
      //fmt.Println("8、住在中间房子的人喝牛奶 ")
      if l[3] != 牛奶{
        continue
      }
      for _, y := range all{
        //fmt.Println("14、挪威人住蓝色房子隔壁 ")
        if y[1] != 蓝色{
          continue
        }
        for _, c := range all{
          for _, x := range all{
            for i := 0; i < 5; i++{
              house[i][颜色] = y[i]
              house[i][国籍] = g[i]
              house[i][宠物] = c[i]
              house[i][香烟] = x[i]
              house[i][饮料] = l[i]
            }
            if check(house){
              fmt.Println("成功")
              return
            }
          }
        }
      }
    }
  }
  fmt.Println("计算失败")
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
  run()
}
