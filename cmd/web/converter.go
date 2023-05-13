package web

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Converter(Res Result) (string, error) {
	res, err := PasteArt(Res.Text, Res.Banner)
	if err != nil {
		return "", fmt.Errorf("Err")
	}
	return res, nil
}

func CheckFile(path string, check string) bool {
	var checkhash string
	if check == "standard" {
		checkhash = "ac85e83127e49ec42487f272d9b9db8b"
	} else if check == "thinkertoy" {
		checkhash = "86d9947457f6a41a18cb98427e314ff8"
	} else if check == "shadow" {
		checkhash = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	}
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash == checkhash
}

func PasteArt(word, banner string) (string, error) {
	path := "banners/" + banner
	if CheckFile(path+".txt", banner) {
		txt := strings.Split(word, "\r\n")
		// reading banner.txt
		content, err := ioutil.ReadFile("banners/" + banner + ".txt")
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("error")
		}
		// array of graphic of characters from banner.txt
		arr := []string{}
		st := ""
		for i := 0; i < len(content); i++ {
			st += string(content[i])
		}
		st = strings.ReplaceAll(st, "\r", "")
		t := strings.Split(st, "\n")
		k, j := 0, 8
		l := len(t)
		for l != 0 {
			for i := k; i < j; i++ {
				if t[i] != "" {
					arr = append(arr, t[i])
				}
			}
			k += 8
			j += 8
			l -= 8
		}
		// save ascii symbols to string ss
		s := ""
		ss := ""
		for v := 0; v < len(txt); v++ {
			if txt[v] == "" && v != 0 {
				fmt.Println()
			} else if txt[v] == "" && v == 0 {
			} else {
				for k := 0; k < 8; k++ {
					for j := 0; j < len(txt[v]); j++ {
						n := int(txt[v][j]) - 32
						for i := n*8 + k; i < n*8+1+k; i++ {
							s += arr[i]
						}
					}
					ss += fmt.Sprintln(s)
					s = ""
				}
			}
		}
		return ss, nil
	} else {
		return "", fmt.Errorf("err")
	}
}

func Check(s []string) bool {
	for _, x := range s {
		for _, l := range x {
			if rune(l) < 32 || rune(l) > 126 {
				return false
			}
		}
	}
	return true
}
