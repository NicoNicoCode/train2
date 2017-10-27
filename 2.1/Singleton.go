package manager
import(
	"sync"
	"fmt"
)

var m *Manager
var lock *sync.Mutex = &sync.Mutex()

func GetInstance() *Manager{
	if m == nil{
		lock.Lock()
		defer lock.Unlock()
		if m == nil{
			m = &Manager{}
		}
	}
	
	return m
}

type Manager struct{}

func (p Manager) Manager() {
	fmt.Println("manager...")
}