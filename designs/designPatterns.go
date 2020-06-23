/*this is a simple way of designing
1. Storage
2. State Machine
3. HTTP API
4. Cron (jobs)
5. ctrl-C handlers
*/

package main

/*state Machine - usually can look like a struct
possible reference to the storage layer, or other
bits of state, or other bits of state
*/
type stateMachine struct {
	state string
	fooc  chan fooReq   // (set)something to do with the struct
	barc  chan barReq   // (get)something to do with the struct
	quitc chan struct{} //quit chan
}

//method on the struct called a "loop", or a "run"
func (sm *stateMachine) loop() {
	for {
		select {
		case r := <-sm.fooc:
			sm.handleFoo(r) //receive requests on
		case r := <-sm.barc:
			sm.handleBar(r) //receive requests on
		case <-sm.quitc:
			return
		}
	}
}

//Public API methods
func (sm *stateMachine) foo() int { //only returns int
	c := make(chan int)
	sm.foo <- fooReq{c}
	return <-c //pass along into the "loop"
}

//MORE ELEGANT DESIGN
type stateMachine2 struct{
		state int
		actionc chan  //there is a single internal channel (THIS IS A CHANNEL OF EMPTY FUNCTIONS)
		quitc chan struct{},
}

//Method to stateMachine2  (here you will select on the action channel and just run the function)
func (sm *stateMachine2) loop(){
		for{
				select{
				case f:=<-sm.actionc:
						f()
						case<-sm.quitc:
						return
				}
		}
}
/* This loop doesn't care what happens, it's just running and restricted to only the serielization point 
this changes the public API functions
*/

//function closure
func (sm *stateMachine2) foo()int{
		c := make (chan int)
		sm.actionc <- func(){
				if sm.state =="A"{
						sm.state ="B"
				}
				c <- 123
		}
		return <-c
}
