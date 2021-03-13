package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Old 原始结构
type Old struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Address string    `json:"address"`
	Sex     *int      `json:"sex"`
	Ignore  string    `json:"ignore"`
	Dec     string    `json:"-"`
	Time    time.Time `json:"time"`
}

// New 该收
type New struct {
	Name      string    `json:"name,omitempty"`
	Age       int       `json:"age,omitempty"`
	Address   string    `json:"address"`
	Sex       *int      `json:"sex,omitempty"`
	Ignore    string    `json:"-"`
	Dec       string    `json:"-,"`
	Time      time.Time `json:"-"`
	Timestamp int64     `json:"timestamp"`
}

// UnmarshalJSON 用于反序列化JSON
func (n *New) UnmarshalJSON(data []byte) error {
	type Alise New
	n1 := new(Alise)
	if err := json.Unmarshal(data, n1); err != nil {
		log.Println(err.Error())
		return err
	}
	n1.Time = time.Unix(n1.Timestamp, 0)
	*n = *(*New)(n1)
	return nil
}

// MarshalJSON 实现json.Marshaler
func (n *New) MarshalJSON() ([]byte, error) {
	type Alise New
	n.Timestamp = n.Time.Unix()
	n1 := (*Alise)(n)
	return json.Marshal(n1)
}

func main() {
	now := time.Now()
	mOld := new(Old)
	mOld.Time = now
	bOldData, mOldErr := json.Marshal(mOld)
	if mOldErr != nil {
		log.Fatalln(mOldErr.Error())
	}
	fmt.Println("old ==> ", string(bOldData))
	mNew := new(New)
	mNew.Time = now
	bNewData, mNewErr := json.Marshal(mNew)
	if mNewErr != nil {
		log.Fatalln(mOldErr.Error())
	}
	fmt.Println("new ==> ", string(bNewData))
	mUnNew := new(New)
	json.Unmarshal([]byte(`{"address":"erdads","-":"","timestamp":1615315849}`), mUnNew)
	fmt.Printf("%+v \n", mUnNew)
}
