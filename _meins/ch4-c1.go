// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Membership struct {
	Type             string
	MessageCharLimit int
}

type UserInfo struct {
	email string
	username string
}

type User struct {
	Name string
	Membership
	UserInfo
}



type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}

type truck struct {
  car
  bedSize int
}

func newUser(name string, membershipType string) User {
	l := 100
	if membershipType == "premium" {
		l = 1000
	}
	u := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: l,
		},
	}
	return u
}

// #region CH4: C2
func (u User) SendMessage(message string, messageLength int) (messageRet string, status bool) {
	status = messageLength <= u.MessageCharLimit
	messageRet = message
	if !status {
		messageRet = ""
	}
	return messageRet, status
}
// #endregion
func main() {
	m18:=truck{
		car: car{
			brand:"Steyr",
			model:"M18",
			doors:4,
			mileage:0,
			frontWheel:wheel{
				radius:35,
				material:"gum",
			},
		},
		bedSize:0,
	}
	fmt.Printf("Truck %v\n",m18)
	us := User{
		Name: "n",
		Membership: Membership{
			Type:             "standard",
			MessageCharLimit: 100,
		},
		UserInfo: UserInfo{
			username: "un",
			email: "un@email.local",
		},
	}
	up := newUser("p", "premium")
	fmt.Printf("%v\n", us.Name)

	m, stat := us.SendMessage("bla", 10)
	fmt.Printf("%v, %v\n", m, stat)
	m, stat = up.SendMessage("blabla", 1001)
	fmt.Printf("%v, %v\n", m, stat)	
}
