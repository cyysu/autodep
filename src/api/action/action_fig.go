package action
import (
    "github.com/codeskyblue/go-sh"
    "fmt"
    "api/common"
)

func FigCreate(request *common.RequestData) string {
    session := sh.NewSession()
    session.ShowCMD = true
    fmt.Println("action.FigCreate:")
/*
    err := session.Call("ssh",request.ServerIP, "touch ","tt.aa")
    if err != nil {
            fmt.Println("exec remote shell faild error:", err)
    }
*/
    return "go-sh"
}
