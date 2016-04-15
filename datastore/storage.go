package tracker

import (
    "appengine"
    "appengine/datastore"
)


type User struct {
    Email string
}


func createUser(ctx appengine.Context, email string) (*datastore.Key, error) {
    u := &User{
        Email: email,
        }
    incompleteKey := datastore.NewIncompleteKey(ctx, "User", nil)
    key, err := datastore.Put(ctx, incompleteKey, u)
    if err != nil {
        return key, err
    }

    return key, nil
}

func getUser(ctx appengine.Context, email string) (u *User, e error) {
    users := []User{}
    q := datastore.NewQuery("User").Filter("Email=", email).Limit(1)
    _, err := q.GetAll(ctx, &users)
    if err != nil {
        return nil, err
    }
    return &users[0], nil
}
