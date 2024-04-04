package users

import (
	"fmt"
	"time"

	"github.com/antero/go/modelos"
)

func AltaUsuario(){
	u :=  new(modelos.User)
	u.AddUser(10, "Eduar", time.Now(), true)
	fmt.Println(u)
}