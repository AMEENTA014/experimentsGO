package main 
import("fmt")
type PhysicalMemory struct{
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
 value := mem.Read(10)
 fmt.Println("value at address 10 " , value )

 }
