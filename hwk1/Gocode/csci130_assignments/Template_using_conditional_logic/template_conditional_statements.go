
package main

import (
"fmt"
"log"
"os"
"text/template"
)

type student struct{
   name string
   person_age int
   major string
        standing string
        study_level string
}

func main(){
        //temp := template.new("template.gohtml")

        s:=student{
           name : "will"
           person_age : "20"
                major : "csci"
                standing : "Senior"
        }

        if s.standing == "Freshmen"{
                s.study_level = "mild"
        }else if s.standing == "Sophomore"{
                s.study_level == "medium"

        }else if s.standing == "junior"{
                s.study_level = "get ready to stay up all night"

        }else if s.standing == "senior"{
                s.study_level = "eh almost done who studies!"
        }

        tpl, err := template.ParseFiles("template.gohtml")
        if err != nil{
                log.Fatalln(err)
        }

        err = tpl.Execute(os.Stdout, s)
        if err != nil{
                log.Fatalln(err)
        }

}
