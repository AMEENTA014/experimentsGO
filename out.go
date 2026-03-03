p
ckage main   // package main is essential for executable   [2]    importing fmt and keeping the order of packaging and importing then declaration
import("fmt")
type PhysicalMemory struct{  // defining a data type which is considered as RAM , planning to use another function to allow space allocation
    data []byte
}
func NewPhysicalMemory(size int32 )   *PhysicalMemory {
            return &PhysicalMemory{
              data: make([]byte , size),  // its used for byte size
	    }
}
func (m * PhysicalMemory)  Read(addr uint32) byte {
        return m.data[addr]
}

func (m * PhysicalMemory)  Write(addr uint32 , value byte)  {
        m.data[addr] =  value 
}
func main(){
 mem := NewPhysicalMemory(1024)
 mem.Write(10 , 42)
 value := mem.Read(10)  // reading from the 10 location
 fmt.Println("value at address 10 " , value )

 }
