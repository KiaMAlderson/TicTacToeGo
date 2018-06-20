package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

var p = fmt.Print
var turn string
var pieces = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Even(number int) bool {
    return number%2 == 0
}

func drawBoard() {
  var hor = "-------------\n"
  for i := 0; i < 3; i++ {
    p(hor)
    p("| ", pieces[(i*3)], " | ", pieces[(i*3)+1], " | ", pieces[(i*3)+2], " |\n")
  }
  p(hor)
}

func makeMove(turn string) {
  //Accepts square to place piece from user, then updates pieces
  var squareint int64
  reader := bufio.NewReader(os.Stdin)
  p("Please enter a value from 1 to 9 that is available: ")
  square, _ := reader.ReadString('\n')
  square = strings.Trim(square, "\n")
  squareint, _ = strconv.ParseInt(square, 0, 64)

  for (squareint <= 0 || squareint > 9) {
    p("Please enter a value from 1 to 9 that is available: ")
    square, _ := reader.ReadString('\n')
    square = strings.Trim(square, "\n")
    squareint, _ = strconv.ParseInt(square, 0, 64)
  }

  for (pieces[squareint-1] == "X") || (pieces[squareint-1] == "O") {
    p("Please enter a value from 1 to 9 that is available: ")
    square, _ := reader.ReadString('\n')
    square = strings.Trim(square, "\n")
    squareint, _ = strconv.ParseInt(square, 0, 64)
  }
  pieces[squareint-1] = turn
}

func checkStatus() bool {
  //Checks for winning rows of 3

  flag := false

  for i := 0; i < 3; i++ {
    //Check horizontals
    //012
    //345
    //678
    if pieces[(i*3)] == pieces[(i*3)+1] && pieces[(i*3)] == pieces[(i*3)+2] {
      flag = true
    }
  }

  for i := 0; i < 3; i++ {
    //Check verticals
    //036
    //147
    //258
    if pieces[(i)] == pieces[(i)+3] && pieces[(i)] == pieces[(i)+6] {
      flag = true
    }
  }

  if pieces[0] == pieces[4] && pieces[0] == pieces[9] {
    //Check L(top)-R diagonal
    flag = true
  }

  if pieces[6] == pieces[4] && pieces[6] == pieces[2] {
    //Check L(bottom)-R diagonal
    flag = true
  }
  return flag
}


func main() {
  p("-----------------------------------------------------------\n")
  p("Welcome to TicTacToe!\n")
  reader := bufio.NewReader(os.Stdin)
  p("x Player 1 name : ")
  name1, _ := reader.ReadString('\n')
  p("O Player 2 name : ")
  name2, _ := reader.ReadString('\n')

  name1 = strings.Trim(name1, "\n")
  name2 = strings.Trim(name2, "\n")

  for i := 0; i < 9; i++ {
    if Even(i){
      p("\n" + name1 + "'s turn!\n")
      turn = "X"
    } else {
      p("\n" + name2 + "'s turn!\n")
      turn = "O"
    }
    drawBoard()
    makeMove(turn)

    if checkStatus() {
      p(turn, " wins!\n")
      break
    }

  }

  p("Thanks for playing!\n")

}
