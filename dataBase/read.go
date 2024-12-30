package dataBase

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
)

func (b *DB) GetUserSchedule(ctx context.Context) map[string]int {
	dsnap, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("users").
		Doc("usersSchedule").Get(ctx)
	if err != nil {
	}
	m := dsnap.Data()
	temp := make(map[string]interface{})
	for str, i := range m {
		if str == b.UserName {
			temp = i.(map[string]interface{})
			break
		}
	}
	ret := make(map[string]int)
	for key, value := range temp {

		ret[key] = int(value.(int64))
	}
	b.Schedule = ret
	return ret

}

func (b *DB) GetScheduleResults(ctx context.Context) (map[string]interface{}, error) {
	dsnap, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("schedule").
		Doc("matches").Get(ctx)

	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	if err != nil {
		fmt.Println(err)
	}
	return m, err
}

func (b *DB) GetUsers(ctx context.Context) ([]string, error) {
	dsnap, err := b.Db.Collection(strconv.Itoa(b.Year)).Doc(b.Code).Collection("users").
		Doc("usersList").Get(ctx)

	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	fmt.Printf("Document type: %#v\n", reflect.TypeOf(m))
	if m == nil {
		return []string{}, err
	}
	ret := make([]string, len(m["users"].([]interface{})))
	for i, v := range m["users"].([]interface{}) {
		ret[i] = v.(string)
	}
	return ret, err
}
