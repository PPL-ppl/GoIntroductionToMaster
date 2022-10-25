package main

//创建 map
//在 map 中搜索值
//向 map 添加新值
//更新 map 中的值
//从 map 中删除值
//了解更多错误相关的知识
//如何创建常量类型的错误
//对错误进行封装
import (
	"testing"
)

func Test_maps(t *testing.T) {
	t.Run("Search-存在", func(t *testing.T) {
		mp := MyMap{"test": "test"}
		got, err := mp.Search("test")
		want := "test"
		assertStrings(t, got, want)

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
	t.Run("Search-不存在", func(t *testing.T) {
		mp := MyMap{"test": "test"}
		_, err := mp.Search("test2")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
	t.Run("Add-存在", func(t *testing.T) {
		var mp = MyMap{}
		err := mp.Add("test", "test")
		assertErrors(t, err, nil)
		assertDefinition(t, mp, "test", "test")
	})
	t.Run("Add-不存在", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := MyMap{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
func assertDefinition(t *testing.T, dictionary MyMap, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}
func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

type MyMap map[string]string

const ErrorNotFount = MyError("could not find the word you were looking for")
const ErrWordExists = MyError("cannot add word because it already exists")
const ErrWordDoesNotExist = MyError("cannot update word because it does not exist")

// 封装错误
type MyError string

// 实现了Error接口
func (error MyError) Error() string {
	return string(error)
}

// Search
func (mp MyMap) Search(key string) (string, error) {
	definition, ok := mp[key]
	if !ok {
		return "", ErrorNotFount
	}
	return definition, nil
}

// Add方法
func (mp MyMap) Add(key, value string) error {
	//存在 报错 不存在添加
	_, err := mp.Search(key)
	switch err {
	case ErrorNotFount:
		mp[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
func (d MyMap) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFount:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d MyMap) Delete(word string) {
	delete(d, word)
}
