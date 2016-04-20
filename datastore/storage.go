package trackerds

import (
    "appengine"
    "appengine/datastore"
)


type User struct {
    Email string
}


func CreateUser(ctx appengine.Context, email string) (key *datastore.Key, err error) {
    u := &User{
        Email: email,
        }
    incompleteKey := datastore.NewIncompleteKey(ctx, "User", nil)
    err = datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
        key, err = datastore.Put(ctx, incompleteKey, u)
        return err
        }, nil)

    return key, err
}

func GetUser(ctx appengine.Context, email string) (u *User, e error) {
    users := []User{}
    q := datastore.NewQuery("User").Filter("Email=", email).Limit(1)
    _, err := q.GetAll(ctx, &users)
    if err != nil {
        return nil, err
    }

    if len(users) > 0 {
        return &users[0], e
    }

    return u, e
}
