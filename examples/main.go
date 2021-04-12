package main

import (
	"fmt"
	"github.com/Genekoh/SchoolUtil"
)

func main() {
	student1 := schoolutil.NewStudent("Gene", "Koh")
	student2 := schoolutil.NewStudent("Emi", "Chua")
	student3 := schoolutil.NewStudent("Danny","Skrrt")
	student4 := schoolutil.NewStudent("Dan","TDM")
	student5 := schoolutil.NewStudent("Daddy","dad")
	student6 := schoolutil.NewStudent("Emilia","Zero")

	var class1 = schoolutil.Class{
		Students: []schoolutil.Student{student1, student2, student3, student4, student5, student6},
	}

	fmt.Println(class1.RandomStudent())
}
