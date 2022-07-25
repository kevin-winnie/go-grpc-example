package service

import (
	_"encoding/gob"
	"encoding/json"
	_"fmt"
	"io"
	_"io"
	_"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// UrlStore - 长链接置换短链接
type UrlStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save1 chan Record
}

type Record struct {
	Key, Url string
}

func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{
		urls: make(map[string]string),
		save1: make(chan Record,1000),
	}
	//file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Fatal("url store :", err)
	//}
	//s.file = file

	if err := s.load(filename); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
	s.Get(s.Get("94ws0"))
	go s.saveLoop(filename)
	return s
}

func (s *UrlStore)saveLoop(filename string)  {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("url store :", err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	for {
		r := <-s.save1
		if err := e.Encode(r);err != nil{
			log.Println("URLStore:", err)
		}
	}
}

// Get -
func (s *UrlStore) Get(key string) string {
	//读锁
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

// Set -
func (s *UrlStore) Set(key, url string) bool {
	//写锁
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.urls[key]
	if ok {
		return false
	}
	s.urls[key] = url
	return true
}

// Count -
func (s *UrlStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// PutKey -
func (s *UrlStore) PutKey(url string) string {
	for {
		key := genKey(5)
		if s.Set(key, url) {
			s.save1 <- Record{
				key,url,
			}
			return key
		}

	}
}

func genKey(count int) string {
	str:="0123456789abcdefghigklmnopqrstuvwxyz"
	strList:=[]byte(str)

	result:=[]byte{}
	i:=0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i< count{
		new:=strList[r.Intn(len(strList))]
		result=append(result,new)
		i=i+1
	}
	return string(result)
}

//func (s *UrlStore) save(key, url string) error {
//	e := gob.NewEncoder(s.file)
//	return e.Encode(Record{
//		key, url,
//	})
//}

func (s *UrlStore) load(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("url store :", err)
	}
	d := json.NewDecoder(file)
	for err == nil {
		var r Record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.Url)
		}
	}

	if err == io.EOF {
		return nil
	}
	return err
	//exist,err := PathExists(filename)
	//if exist {
	//	data ,err := ioutil.ReadFile(filename)
	//	if err != nil{
	//		fmt.Println("readfile",err)
	//	}
	//	var config Record
	//	err = json.Unmarshal(data,&config)
	//	fmt.Println(config.Url)
	//	fmt.Println(config.Key)
	//	s.Set(config.Key,config.Url)
	//	return err
	//}
	//
	//return err

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
