# Object Pool Pattern
***
The object pool creational design pattern is used to prepare and keep multiple instances according to the demand expectation.

## Implementation
***
    package pool
    
    type Pool chan *Object
    
    func New(total int) *Pool {
    p := make(Pool, total)
    
        for i := 0; i < total; i++ {
            p <- new(Object)
        }
    
        return &p
    }


## Usage
***
Given below is a simple lifecycle example on an object pool.

    p := pool.New(2)
    
    select {
    case obj := <-p:
    obj.Do( /*...*/ )
    
        p <- obj
    default:
    // No more objects left — retry later or fail
    return
    }

## Rule of Thumb
***
- Object pool pattern is usefull in cases where object initialization is more expensive than the object maintenance.
- If there are spikes in demand to a steady demand, the maintenance overhead might overweight the benefits of an object pool.
- It has positive effects on performance due to objects being initialized beforehand.