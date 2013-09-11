package model

import (
	"os"
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"morr.cc/nutsh.git/parser"
)

type Tutorial struct {
	Name    string
	Target  string
	Version int
	Basedir string
	Lessons map[string]*Lesson
}

type Lesson struct {
	Root parser.Node
	Done bool
}

func Init(dir string) Tutorial {
	info, _ := ioutil.ReadFile(dir + "/info.yaml")
	var tut Tutorial
	goyaml.Unmarshal(info, &tut)
	tut.Basedir = dir
	tut.Lessons = make(map[string]*Lesson)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if len(file.Name()) > 7 && file.Name()[len(file.Name())-6:len(file.Name())] == ".nutsh" {
			content, _ := ioutil.ReadFile(dir + "/" + file.Name())
			rootnode := parser.Parse(string(content))
			tut.Lessons[file.Name()[0:len(file.Name())-6]] = &Lesson{rootnode, false}
		}
	}

	var done_lessons []string

	s, err := ioutil.ReadFile(dir+"/progress.yaml")
	if err == nil {
		goyaml.Unmarshal(s, &done_lessons)
	}
	for _, l := range done_lessons {
		tut.Lessons[l].Done = true
	}

	return tut
}

func (t Tutorial) SelectLesson() (*Lesson, bool) {
	fmt.Printf("== %s ==\n", t.Name)
	fmt.Println("0: (Beenden)")
	i := 1
	lessons := make([]*Lesson, 0)
	for name, l := range t.Lessons {
		if l.Done {
			fmt.Print("[35m")
		}
		fmt.Printf("%d: ", i)
		fmt.Print(name)
		if l.Done {
			fmt.Print("[0m")
		}
		fmt.Println()
		lessons = append(lessons, l)
		i += 1
	}

	sel := 0
tryagain:
	fmt.Print("Bitte wählen Sie eine Lektion: ")
	_, err := fmt.Scanf("%d", &sel)
	if err != nil {
		goto tryagain
	}


	if sel < 0 || sel > len(lessons) {
		goto tryagain
	}

	if sel == 0 {
		return &Lesson{}, false
	}

	return lessons[sel-1], true
}

func (t Tutorial) SaveProgress() {
	done_lessons := make([]string, 0)
	for name, l := range t.Lessons {
		if l.Done {
			done_lessons = append(done_lessons, name)
		}
	}
	s, _ := goyaml.Marshal(done_lessons)
	println(string(s))
	f, _ := os.Create(t.Basedir+"/progress.yaml")
	f.Write(s)
	f.Close()
}
