package dao

import (
	"context"
	"lucar/shared/id"
	mgo "lucar/shared/mongo"
	"lucar/shared/mongo/objid"
	mongotesting "lucar/shared/mongo/testing"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func InitMongo() (*mongo.Database, context.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	// 设置账户密码
// 	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
// 		Username: "admin",
// 		Password: "123456",
// 	}).ApplyURI("mongodb://127.0.0.1:27017"))
// 	// clientOptions := options.Client().ApplyURI("mongodb://admin:123456@127.0.0.1:27017")
// 	// client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, nil
// 	}
// 	return client.Database("lucar"), ctx
// }

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 设置账户密码
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://127.0.0.1:27017"))
	// clientOptions := options.Client().ApplyURI("mongodb://admin:123456@127.0.0.1:27017")
	// client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		t.Fatalf("can not connect mongodb: %v\n", err)
	}
	m := NewMongo(client.Database("lucar"))
	_, err = m.col.InsertMany(ctx, []interface{}{
		bson.M{
			mgo.IDField: objid.MustFromID(id.AccountID("5f7c245ab0361e00ffb9fd6f")),
			openIDField: "openid_1",
		},
		bson.M{
			mgo.IDField: objid.MustFromID(id.AccountID("5f7c245ab0361e00ffb9fd70")),
			openIDField: "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}

	m.NewObjID = func() primitive.ObjectID {
		objID := objid.MustFromID(id.AccountID("67ccf51a69f3d4554b28a87c"))
		return objID
	}
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "5f7c245ab0361e00ffb9fd6f",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "5f7c245ab0361e00ffb9fd70",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "67ccf51a69f3d4554b28a87c",
		},
	}

	for _, cc := range cases {
		// 子测试
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("failed resolve account id for %q: %v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id: want: %q, got:%q", cc.want, id)
			}
		})
	}

	id, err := m.ResolveAccountID(ctx, "678")
	if err != nil {
		t.Errorf("failed resolve account id for 678: %v", err)
	} else {
		want := "67ccf51a69f3d4554b28a87c"
		if id != want {
			t.Errorf("resolve account id: want: %q, got:%q", want, id)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
