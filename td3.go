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


var tabfilename = []string{
	"NC_000964.fna",
	"NC_009725.fna",
	"NC_009848.fna",
	"NC_014174.fna",
	"NC_012472.fna",
	"NC_003997.fna",
	"NC_015634.fna",
	"NC_014639.fna",
	"NC_006322.fna",
	"NC_000913.fna",
	"NC_011770.fna",
	"NC_009428.fna",
	"NC_016114.fna",
	"NC_012803.fna",
	"NC_002662.fna",
}

  var mat_sym [15][15]float64


func CompressString(data_s string)string{
    str:="data.gz"  
    file, _:= os.Create(str)

    w, _ := gzip.NewWriterLevel(file, gzip.BestCompression)
    w.Write([]byte(data_s))
    w.Close()
    return str
}




func ProcessFile(s string)string{
	
	//fmt.Println("Debut")
	file, _ := os.Open(s)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var lineNumber int
	var shortTreatment bool
	//scanner.Split(bufio.ScanLines)
	next_line_is_dna := false
	for scanner.Scan() {
		lineNumber++
         
		if shortTreatment && (lineNumber > 4000) {
			break
		}
		//var val_str string
		fmt.Println("mileu",lineNumber)
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


func distance(a string,b string )float64{
	
	//fmt.Println("calul !!")	
	za_b:=(float64)(getsize(CompressString(a+b)))
	//fmt.Println("calcul 1")
	za:=(float64)(getsize(CompressString(a)))
	//fmt.Println("calcul 2")
	zb:=(float64)(getsize(CompressString(b)))
	//fmt.Println("calcul 3")
	za_a:=(float64)(getsize(CompressString(a+a)))
	zb_b:=(float64)(getsize(CompressString(b+b)))
	
	var d float64=(float64)((za_b/(za+zb))-(za_a/(4*za))-(zb_b/(4*zb)))
	
	//~ fmt.Println(d)
		//~ fmt.Println(za)
		//~ fmt.Println(zb)
		//~ fmt.Println(za_b)
		//~ fmt.Println(za_a)
		//~ fmt.Println(zb_b)
	//~ 
	
	return d
}





func matrix_sym(){
	
   for i := 0; i < 15; i++ {
	   fmt.Println("out ! ")
        for j := i+1; j < 15; j++ {
			fmt.Println("in !!!")
		    mat_sym[i][j]=distance(ProcessFile(tabfilename[i]),ProcessFile(tabfilename[j]))	
		    
		    fmt.Println(mat_sym[i][j])
		    fmt.Println("in 2!!!")
	    }
	}
	
	
	for i := 1; i < 15; i++ {
        for j := 0; j < i; j++ {
			mat_sym[i][j]=mat_sym[j][i]
		}
	}
	
	fmt.Println(" Matrix not empty !! ")
	
	file, err := os.Create("matrix.dat")
    if err != nil {
       fmt.Println(err)
       return
    }

	
	for i := 0; i < 15; i++ {
        for j := 0; j < 15; j++ {
	        
	       fmt.Fprintf(file, "%2f ",mat_sym[i][j])
	        
	    }
	    fmt.Println("")
    }
}




func main() {
	
	//file1 := `NC_000964.fna`
	//file2 := `NC_000913.fna`
    
    
	fmt.Println(" ...in progress ! ");
	
	matrix_sym()
    //ioutil.WriteFile("test.data",tobytes(mat_sym) , 0644)
	
	fmt.Println(" Matrix ready !! ")
	
	
	//var str1 string=ProcessFile(file1)
	//var str2 string=ProcessFile(file2)
	
	//var val float64=distance(str1,str2)
	
	////fileComp:=CompressString(str)
	
	//fmt.Println(val)
	
    // Done.
   fmt.Println("End  !")
}
