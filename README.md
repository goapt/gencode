# gencode
Generate unique code

同一秒内生成10W个不重复的，无规则的16位纯数字编码，并且当一秒内生成超过10W个编码，程序会阻塞1秒之后再次生产编码


```
import "github.com/goapt/gencode"

func main(){
    manager := gencode.New("", "20180919",false)
    code := manager.Get()
    fmt.Println(code)
    c, err := manager.Verify(code)
    fmt.Println(c, err)	
}
```