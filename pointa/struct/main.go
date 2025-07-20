package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

// メモリ位置で引数を受けているため、値がコピーされない。
func changeVertex2(v *Vertex) {
	(*v).X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println("v2",v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("v6:%T %v\n", v6, v6)

	// v6と一緒の結果
	v7 := &Vertex{}
	fmt.Printf("v7:%T %v\n", v7, v7)

	/*
		    s := make([]int, 0)
			s := []int{}
			fmt.Println(s)
	*/

	v8 := Vertex{1, 2, "test"}
	changeVertex(v8)
	fmt.Println("changeVertex",v8)

	v9 := &Vertex{1, 2, "test"}
	changeVertex2(v9)
	fmt.Println("changeVertex2",v9)
	

	fmt.Println("\n================================================")
	practice()
}


func practice() {
  var i int = 100 // i = 100
  var j int = 200 // j = 200

  var p1 *int // p1 は int を指すポインタ (最初はnil)
  var p2 *int // p2 も int を指すポインタ (最初はnil)

  p1 = &i // p1 に i のアドレスを代入
          // (例: p1 = 0x1000) <- i のアドレス
          // *p1 は 100 を指す

  p2 = &j // p2 に j のアドレスを代入
          // (例: p2 = 0x2000) <- j のアドレス
          // *p2 は 200 を指す

	fmt.Println("p1, p2", p1, p2)
  fmt.Println("*p1, *p2", *p1, *p2)
  // 出力: *p1, *p2 100 200
  // 現在の状態:
  // i: 100 (アドレス 0x1000)
  // j: 200 (アドレス 0x2000)
  // p1: 0x1000 (i のアドレスを指している)
  // p2: 0x2000 (j のアドレスを指している)

  i = *p1 + *p2
  // *p1 は 100 (p1がiを指しているため)
  // *p2 は 200 (p2がjを指しているため)
  // i = 100 + 200 = 300
  // i の値が 100 から 300 に変更される。//////////////////////////////
  // 現在の状態:
  // i: 300 (アドレス 0x1000) <- 値が変わった！ //////////////////////////////
  // j: 200 (アドレス 0x2000)
  // p1: 0x1000 (i のアドレスを指している)
  // p2: 0x2000 (j のアドレスを指している)

  fmt.Println("i", i)
  // 出力: i 300

  p2 = p1
  // ここがポイントです！
  // p1 の値は i のアドレス (0x1000)。
  // p2 に p1 の値 (i のアドレス) が代入されます。
  // つまり、p2 も i のアドレス 0x1000 を指すようになります。///////////////////////
  // p2 は、もう j (0x2000) を指していません。//////////////////////////////////
  // 現在の状態:
  // i: 300 (アドレス 0x1000)
  // j: 200 (アドレス 0x2000)
  // p1: 0x1000 (i のアドレスを指している)
  // p2: 0x1000 (i のアドレスを指している) <- p2が指す先が変わった！

  fmt.Println("*p2", *p2)
  // p2 は現在 0x1000 (i のアドレス) を指しています。
  // *p2 は、0x1000 番地にある値を取得します。
  // 0x1000 番地には i の現在の値である 300 が格納されています。
  // よって、*p2 は 300 となります。
  // 出力: *p2 300

  j = *p2 + i
  // *p2 は 300 (p2がiを指しているため)
  // i は 300
  // j = 300 + 300 = 600
  // j の値が 200 から 600 に変更される。
  // 現在の状態:
  // i: 300 (アドレス 0x1000)
  // j: 600 (アドレス 0x2000) <- 値が変わった！
  // p1: 0x1000 (i のアドレスを指している)
  // p2: 0x1000 (i のアドレスを指している)

  fmt.Println("演習2:", j)
  // 出力: 演習2: 600
}

// Go言語において、ポインタ変数の前に * (アスタリスク) をつけると、それは「そのポインタが指し示しているメモリ番地にある値」を参照するという意味になります。
// この操作を逆参照 (dereference) と呼びます。
