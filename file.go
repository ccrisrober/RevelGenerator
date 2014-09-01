package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

func UpcaseInitial(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

func main() {
	/*r := bufio.NewReader(os.Stdin)
	s, _ := Readln(r)
	fmt.Println("Dime nombre controlador: ")
	for len(s) == 0 {
	    s, _ = Readln(r)
	}*/
	s := "MiControlador CRUD"
	actions := strings.Split(s, " ")

	file, err := os.Create(actions[0] + ".go")
	if err != nil {
		return
	}
	defer file.Close()

	controller := actions[0] //+ "Controller"
	fmt.Println("Controlador: ", controller)
	file.WriteString("package controllers\r\n")
	file.WriteString("import \"github.com/revel/revel\"\r\n")
	file.WriteString("type " + controller + " struct {\r\n")
	file.WriteString("\t*revel.Controller\r\n\r\n")
	file.WriteString("}\r\n\r\n")

	if actions[1] == "CRUD" {
		file.WriteString("func (c " + controller + ") Index() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Create() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Insert() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Show() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Edit() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Update() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		file.WriteString("func (c " + controller + ") Delete() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
	} else {
		size := len(actions)
		for i:= 1; i < size; i++ {
			actions[i] = strings.ToLower(actions[i])
			actions[i] = UpcaseInitial(actions[i])
			fmt.Println(actions[i])
			file.WriteString("func (c " + controller + ") " + actions[i] + "() revel.Result {\r\n")
			file.WriteString("\treturn c.Render()\r\n")
			file.WriteString("}\r\n\r\n")
		}
	}
	}