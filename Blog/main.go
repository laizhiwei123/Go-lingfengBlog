package main

import (
	"Blog/models"
	"Blog/routers"
)

func main() {
	models.Init()
	routers.Router()
}

/*

   <form  action="/login" method="post" >
       <label>UserName:</label>
       <input name="UserName">
       <br>
       <br>
       <label>PassWord:</label>
       <input name="PassWord">
       <br>
       <br>
       <div style="position: fixed; top: 400px; left: 600px"><input  type="checkbox" name="checkbox"value="on">RememberPassword </div>
       <input  class="Centered1" type="submit" value="Login">
   </form>
*/