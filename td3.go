package main

import (
    "bufio"
    "compress/gzip"
    "fmt"
    //"io/ioutil"
    "os"
    "log"
    //"strings"
)

var(
 file_inf os.FileInfo
 DNAdata  string
 data     string
 err      error
)

func CompressString(data_s string)string{
    str:="data.gz"  
    file, _:= os.Create(str)

    w, _ := gzip.NewWriterLevel(file, gzip.BestCompression)
    w.Write([]byte(data_s))
    w.Close()
    return str
}




func ProcessFile(s string)string{
	file, _ := os.Open(s)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var lineNumber int
	var shortTreatment bool
	//scanner.Split(bufio.ScanLines)
	next_line_is_dna := false
	for scanner.Scan() {
		lineNumber++

		if shortTreatment && (lineNumber > 40000) {
			break
		}
		//var val_str string
		val_str := scanner.Text()
		if len(val_str) > 0 {
			if val_str[0] == '>' {
				next_line_is_dna = true
				continue
			} else if next_line_is_dna {
				DNAdata = DNAdata + val_str
				continue
			}
		}
	}
	return DNAdata
}


func ProcessLine(s string, m map[string]int, k int) {
	if len(s) > (k) {
		s_minus_k := s[:len(s)-k]
		for i, _ := range s_minus_k {
			m[s[i:i+k]]++
			//how to speed up?
		}
	}
}

//get size of compressed file in bytes
 
func getsize(s string)int64{
	file_inf, err = os.Stat(s)
    if err != nil {
        log.Fatal(err)
    }
    return (file_inf.Size())
	
}


func distance(a string,b string )float32{
		
	za_b:=getsize(CompressString(a+b))
	za:=getsize(CompressString(a))
	zb:=getsize(CompressString(b))
	za_a:=getsize(CompressString(a+a))
	zb_b:=getsize(CompressString(b+b))
	
	var d float32
	d=(float32)((za_b/(za+zb))-(za_a/(4*za))-(zb_b/(4*zb)))
	
	fmt.Println(d)
	
	return d
}



func main() {
	
	file1 := `NC_000964.fna`
	file2 := `NC_000964.fna`
    
    
	_, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "the file %s doesn't exist! \n", file)
		os.Exit(1)
	}
	
	
	fmt.Println(" ...in progress ! ");
	
	str:=ProcessFile(file)
	
	fileComp:=CompressString(str)
	
	fmt.Println(getsize(fileComp))
    // Done.
    fmt.Println("File compressed !")
}
