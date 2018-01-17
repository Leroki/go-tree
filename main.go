package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func tabs(rep int, aaa []bool) string {
	var ret string
	if rep != 0 {
		for i := 0; i < rep; i++ {
			if aaa[i] {
				ret += "│" + "   "
				//ret += "│" + "\t"
			} else {
				ret += "    "
				//ret += "\t"
			}
		}
	} else {
		ret += "│" + "   "
		//ret += "│" + "\t"
	}
	return ret
}

func Size(size int64) string {
	if size != 0 {
		return fmt.Sprintf(" (%db)", size)
	} else {
		return " (empty)"
	}
}

func myReadDirWithFiles(path string, out *os.File, vloj int, aaa []bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	rng := len(files) - 1
	aaa = append(aaa, true)
	for i, file := range files {
		if file.IsDir() {
			if i == rng {
				aaa[vloj] = false
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "└───" + file.Name() + "\n")
					myReadDirWithFiles(path+"/"+file.Name(), out, vloj+1, aaa)
				} else {
					out.WriteString("└───" + file.Name() + "\n")
					myReadDirWithFiles(path+"/"+file.Name(), out, vloj+1, aaa)
				}
			} else {
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "├───" + file.Name() + "\n")
					myReadDirWithFiles(path+"/"+file.Name(), out, vloj+1, aaa)
				} else {
					out.WriteString("├───" + file.Name() + "\n")
					myReadDirWithFiles(path+"/"+file.Name(), out, vloj+1, aaa)
				}
			}
		} else {
			if i == rng {
				aaa[vloj] = false
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "└───" + file.Name() + Size(file.Size()) + "\n")
				} else {
					out.WriteString("└───" + file.Name() + Size(file.Size()) + "\n")
				}
			} else {
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "├───" + file.Name() + Size(file.Size()) + "\n")
				} else {
					out.WriteString("├───" + file.Name() + Size(file.Size()) + "\n")
				}
			}
		}
	}
}

func myReadDir(path string, out *os.File, vloj int, aaa []bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	rng := -1
	for _, v := range files {
		if v.IsDir() {
			rng++
		}
	}
	i := 0
	aaa = append(aaa, true)
	for _, file := range files {
		if file.IsDir() {
			if i == rng {
				aaa[vloj] = false
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "└───" + file.Name() + "\n")
					myReadDir(path+"/"+file.Name(), out, vloj+1, aaa)
				} else {
					out.WriteString("└───" + file.Name() + "\n")
					myReadDir(path+"/"+file.Name(), out, vloj+1, aaa)
				}
			} else {
				if vloj != 0 {
					out.WriteString(tabs(vloj, aaa) + "├───" + file.Name() + "\n")
					myReadDir(path+"/"+file.Name(), out, vloj+1, aaa)
				} else {
					out.WriteString("├───" + file.Name() + "\n")
					myReadDir(path+"/"+file.Name(), out, vloj+1, aaa)
				}
			}
			i++
		}
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	var str string
	var aaa = make([]bool, 0, 10)
	if printFiles {
		myReadDirWithFiles(path, out, 0, aaa)
	} else {
		myReadDir(path, out, 0, aaa)
	}
	out.WriteString(str)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
