type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserRepository struct {
    users map[string]User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        users: make(map[string]User),
    }
}

func (r *UserRepository) Save(user User) {
    r.users[user.ID] = user
}

func (r *UserRepository) FindById(id string) (User, bool) {
    user, exists := r.users[id]
    return user, exists
}

func (r *UserRepository) Update(user User) bool {
    _, exists := r.users[user.ID]
    if exists {
        r.users[user.ID] = user
    }
    return exists
}

func (r *UserRepository) Delete(id string) bool {
    _, exists := r.users[id]
    if exists {
        delete(r.users, id)
    }
    return exists
}