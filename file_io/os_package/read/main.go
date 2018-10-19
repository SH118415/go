package main
import ( "fmt"
        "os"
)
func main(){
        file,err:=os.Open("myfile")
        if err!=nil {
                panic(err)
        }
        //The data read from the file will be of the size of the byte array
	data:=make([]byte,5)

	//I want to seek the file pointer to the 6th byte of the data
        file.Seek(6,0) 
        //Seek has two arguments, first one is the position of file pointer to read data
        //From where in the file to read from, 0 - beginning, 1-at current position, 2-from end

	//read the data from file to the byte array 'data'
        file.Read(data) 

        //fmt.Printf("The file data is %s\n",string(data))
        fmt.Println(string(data))
        file.Seek(0,0)
        newdata:=make([]byte,5)
        file.Read(newdata)
        fmt.Println(string(newdata))
        file.Close()
}
