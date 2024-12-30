package dataBase

import (
	"context"
	"log"
	"strconv"
)

type DBScouterTemplate struct {
	matches map[string]interface{}
}

func (b *DB) CreateUser(ctx context.Context, userName string) error {
	f, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("users").
		Doc("usersList").Get(ctx)
	ret := make(map[string]interface{})
	if f.Data() != nil {
		h := f.Data()["users"].([]interface{})
		n := make([]string, len(h))
		for i, v := range h {
			if v.(string) == userName {
				return err
			}
			n[i] = v.(string)
		}
		n = append(n, userName)
		println(f)
		println(h)
		ret["users"] = n
	} else {
		ret["users"] = []string{userName}
	}
	i, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("users").
		Doc("usersList").Set(ctx, ret)
	println(i)
	return err
}
func (b *DB) SubmitData(ctx context.Context, data map[string]int, matchNum string, teamNum string) error {
	old, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("schedule").
		Doc("matches").Get(ctx)
	f := old.Data()
	f[matchNum].(map[string]interface{})[teamNum].(map[string]interface{})["results"] = data
	f[matchNum].(map[string]interface{})[teamNum].(map[string]interface{})["gotData"] = true
	sum := 0
	for _, v := range data {
		sum += v
	}
	f[matchNum].(map[string]interface{})[teamNum].(map[string]interface{})["results"].(map[string]int)["sum"] = sum
	_, err = b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("schedule").
		Doc("matches").Set(ctx, f)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}
