package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type FileSorter struct {
	files []fs.FileInfo
	by    func(f1, f2 *fs.FileInfo) bool
}

func (f *FileSorter) Len() int {
	return len(f.files)
}

func (f *FileSorter) Less(i, j int) bool {
	return f.by(&f.files[i], &f.files[j])
}

func (f *FileSorter) Swap(i, j int) {
	f.files[i], f.files[j] = f.files[j], f.files[i]
}

type By func(f1, f2 *fs.FileInfo) bool

func (by By) Sort(files []fs.FileInfo) {
	ps := &FileSorter{
		files: files,
		by:    by,
	}
	sort.Sort(ps)
}

// Сортирует слайс структур fs.FileInfo по полю Name
func sortFilesByName(files *[]fs.FileInfo) {
	name := func(f1, f2 *fs.FileInfo) bool {
		return (*f1).Name() < (*f2).Name()
	}

	By(name).Sort(*files)
}

// Функция вывода дерева каталогов и файлов. Вывод дерева осуществляется с помощью рекурсии.
//  out - куда осуществяется вывод;
//  path - путь к директории, для которой выводятся файлы и каталоги;
//  printFiles - флаг, определяющий, будут ли выводиться файлы;
//  deep - глубина текущего каталога. Если уровень рекурсии начальный, то глубина - это пустая строка.
func printDirTree(out io.Writer, path string, printFiles bool, deep string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if !printFiles {
		deleteAllFiles(&files)
	}

	sortFilesByName(&files)

	for i, file := range files {
		switch {
		case file.IsDir() && i == len(files)-1:
			fmt.Fprintln(out, deep+"└───"+file.Name())
			printDirTree(out, filepath.Join(path, file.Name()), printFiles, deep+"\t")
		case file.IsDir() && i != len(files)-1:
			fmt.Fprintln(out, deep+"├───"+file.Name())
			printDirTree(out, filepath.Join(path, file.Name()), printFiles, deep+"│\t")
		case !file.IsDir() && i == len(files)-1:
			fmt.Fprintln(out, deep+"└───"+file.Name()+getStringFileSize(file.Size()))
		case !file.IsDir() && i != len(files)-1:
			fmt.Fprintln(out, deep+"├───"+file.Name()+getStringFileSize(file.Size()))
		}
	}

	return nil
}

// Возвращает строку с размером файла, если размер не нулевой. Иначе возвращает информацию о том, что файл пустой.
// Нужна для красивого отображения размера файла.
func getStringFileSize(size int64) string {
	if size > 0 {
		return fmt.Sprintf(" (%vb)", size)
	} else {
		return " (empty)"
	}
}

// Удаляет все файлы из files - списка файлов и директорий.
func deleteAllFiles(files *[]fs.FileInfo) {
	for i := 0; i < len(*files); i++ {
		if !(*files)[i].IsDir() {
			*files = append((*files)[:i], (*files)[i+1:]...)
			i--
		}
	}
}

func main() {
	// Куда осуществляется вывод каталогов и файлов (в данном случае в консоль)
	out := os.Stdout

	// Проверка на то, указаны ли аргументы
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	// Путь к директории, дерево которой нужно вывести
	path := os.Args[1]
	// Флаг, указывающий печатать ли файлы при выводе
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := printDirTree(out, path, printFiles, "")
	if err != nil {
		panic(err.Error())
	}
}
