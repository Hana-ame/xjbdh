package dwz

import (
	"fmt"
	"math/rand"
)

// import "crypto/rand"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func T() {

	Init()

}

// not tested
func GenID(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		ai := rand.Intn(len(alphabet))
		b[i] = alphabet[ai]
	}
	return string(b)
}

func GetLink(id string) (*Link, error) { // works well
	l := &Link{
		Id: id,
	}
	err := l.Read(db)
	if err != nil {
		return nil, err
	}

	// TODO：update ClickCnt

	return l, nil
}

func GetUrl(id string) (string, error) {
	l, err := GetLink(id)
	if err != nil {
		return "", err
	}
	return l.Url, nil
}

func AddUrl(title, url, description string) error { // works well
	// var id string
	// var err error
	id := GenID(4)
	for _, err := GetUrl(id); err == nil; {
		fmt.Println(err)
		id = GenID(4)
		_, err = GetUrl(id)
	}

	// create a link
	link := &Link{
		Id:          id,
		Title:       title,
		Url:         url,
		Description: description,
		ClickCnt:    0,
	}
	if err := link.Create(db); err != nil {
		return err
	}

	return nil
}

// not tested
func DelUrl(id string) error {
	link := &Link{
		Id:          id,
		Url:         "",
		Description: "",
		ClickCnt:    0,
	}
	return link.Delete(db)
}

// not tested
func AddTag(tag, id string) error {
	t := &Tag{tag, id, 0}

	err := t.Read(db)
	if err == nil { // record found
		t.Cnt += 1
		t.Update(db)
	}

	err = t.Create(db)
	if err != nil {
		return err
	}

	return nil
}

// not tested
// must be tested
// https://gorm.io/zh_CN/docs/query.html#%E6%A3%80%E7%B4%A2%E5%85%A8%E9%83%A8%E5%AF%B9%E8%B1%A1
// https://gorm.io/zh_CN/docs/query.html#%E6%9D%A1%E4%BB%B6
func ReadLinksByTag(tag string) ([]*Link, error) {
	var rst []*Link
	tx := db.Model(&Link{}).Joins("INNER JOIN `tags` ON tags.id = links.id").Where("tags.tag=?", tag).Find(&rst)
	return rst, tx.Error
}

// tx := db.Raw("select links.id, links.title, links.description, links.url, links, ").Find(rst)

// if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
// 	Joins("JOIN movies on artist_movies.movie_id=movies.id").Where("movies.title=?", "Nayagan").
// 	Group("artists.id").Find(&artists).Error; err != nil {
// 		log.Fatal(err)
// }

// for _, ar := range artists {
// 	fmt.Println(ar.Name)
// }
