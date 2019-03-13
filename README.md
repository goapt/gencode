# gencode
Generate unique code

```
import "github.com/ilibs/gencode"

func main(){
    manager := gencode.New("", "20180919")
    code := manager.Get()
    fmt.Println(code)
    c, err := manager.Verify(code)
    fmt.Println(c, err)	
}
```