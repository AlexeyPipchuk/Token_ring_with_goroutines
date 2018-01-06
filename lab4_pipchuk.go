package main
import (
	"fmt"
	"math/rand"
	"time"
)

type Token struct {
	data string
	recipient int 
	ttl int 
}

type Node struct {
  data string
  position int
}


func show(){
  fmt.Print("\n")
  for i := range nodes {   
    if (nodes[i].data=="") {
        fmt.Print(i,"'th node: the message did not reach or is missing\n")
        continue;
    }
        fmt.Print(i,"'th node has this data: ",nodes[i].data,"\n")
    }
}

func drop (chanel chan Token, i int){
  token:= <- chanel
  for j := range nodes{
    if (token.recipient == nodes[j].position) {
    nodes[j].data = token.data
    }else{
      token.ttl--
      if(token.ttl<=0){
        break
      }
    }
  }
  defer close(chanel)
}

func set(i int, N int) Token {
    var token Token
    token.recipient = rand.Intn(N)
    token.ttl = (rand.Intn(N)+(N/2)) // to minimize chance timeout
    token.data = fmt.Sprintf("data for node ¹%d",token.recipient)
    return token
}

var nodes []Node;
func main() {
	var N int
	rand.Seed(time.Now().UnixNano())
  fmt.Scanf("%d", &N)
  chans := make([]chan Token, N)
  nodes = make([]Node, N)
  for i := range nodes {                
        nodes[i].position = i
    }
  for i := range chans {                
      chans[i] = make(chan Token,1) 
      a := set(i,N)
      chans[i]<- a
      go drop(chans[i],i)
  }
  show()
  time.Sleep(1 * time.Second)
}