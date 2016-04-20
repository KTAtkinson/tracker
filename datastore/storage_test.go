package trackerds

import (
    "testing"
    
    "appengine/datastore"
    "appengine/aetest"
)


func TestCreateUser(t *testing.T) {
    ctx, err := aetest.NewContext(nil)
    if err != nil {
        t.Fatal(err)
    }
    defer ctx.Close()

    email := "testing@testing.go"
    newKey, err := CreateUser(ctx, email)
    if err != nil {
        t.Errorf("Failed to create a new user: %v", err)
    }

    u := User{}
    datastore.Get(ctx, newKey, &u)

    if u.Email != email {
        t.Errorf("Expected email to be %s, found %v.", email, u.Email)
    }
}

func TestGetUser(t *testing.T) {
    ctx, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
    if err != nil {
        t.Fatal(err)
    }
    defer ctx.Close()

    email := "testing@testing.go"
    _, err = CreateUser(ctx, email)
    u, e := GetUser(ctx, email)
    if e != nil {
        t.Fatalf("Error while retieving users: %#v", e)
     }

    if u.Email != email {
        t.Error("Expected email to be %s, found %#v.", email, u.Email)
    }
}

func TestRetrieveNoUser(t *testing.T) {
    ctx, err := aetest.NewContext(&aetest.Options{StronglyConsistentDatastore: true})
    if err != nil {
        t.Fatal(err)
    }
    defer ctx.Close()

    u, _ := GetUser(ctx, "nothing@testing.com")

    if u != nil {
        t.Errorf("Expected user to be nil found %#v.", u)
    }
}
