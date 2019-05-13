package servmodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

type Keeper struct {
	Name       string
	Acct       string
	Secret     string
	CreateTime time.Time
	UpdateTime time.Time
}

func KeeperWrapper(form url.Values) Keeper {
	name := form.Get("name")
	acct := form.Get("acct")
	secret := form.Get("secret")
	ctime, _ := time.Parse(time.RFC3339Nano, form.Get("ctime"))
	utime, _ := time.Parse(time.RFC3339Nano, form.Get("utime"))
	return Keeper{Name: name, Acct: acct, Secret: secret, CreateTime: ctime, UpdateTime: utime}
}

func KeeperFilterWrapper(form url.Values) bson.M {
	var r = make(map[string]interface{})

	name := form.Get("name")
	if name != "" {
		r["name"] = name
	}
	acct := form.Get("acct")
	if acct != "" {
		r["acct"] = acct
	}
	secret := form.Get("secret")
	if secret != "" {
		r["secret"] = secret
	}

	return bson.M(r)
}

type News struct {
	KeeperId primitive.ObjectID
	Title    string
	Content  bson.Raw
	DateTime time.Time
}

func NewsWrapper(form url.Values) News {
	keeperId, _ := primitive.ObjectIDFromHex(form.Get("keeperId"))
	title := form.Get("title")
	content := form.Get("content")
	dateTime, _ := time.Parse(time.RFC3339Nano, form.Get("dateTime"))
	return News{KeeperId: keeperId, Title: title, Content: bson.Raw(content), DateTime: dateTime}
}
