package main 
import 
( 
    "fmt"
    "os"
) 
func main {
fmt.Println(banner.inline("R3SP1R1D0N"))
var path  string 
var nonpath string 
fmt print ('Вам известен путь до config.yaml (Y/n)') 
fmt.Fscan; (os.Stdin, &path ) 
if fmt.Fscan ==Y
fmt print ('Введите путь до config.yaml') 
else fmt print ('жалко')  
f, _ := filepath.Glob("config.yaml") 
file, err := os.Open ("config.yaml")
if err ! = nill { 
    fmt.Println(err)
    os.Exit(1)
    defer file.close()  
    data :=make ([]byte, 64)
    for{ 
        n,err := file.Read
    }
}

