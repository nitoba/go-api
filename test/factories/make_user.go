package factories

import (
	"github.com/jaswdr/faker"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	pkg "github.com/nitoba/go-api/pkg/entity"
)

func MakeUser(props ...map[string]interface{}) *entity.User {
	fake := faker.New()
	id := pkg.NewID()
	name := fake.Lorem().Word()
	email := fake.Internet().Email()

	if len(props) > 0 && props[0]["id"] != nil {
		id, _ = pkg.ParseID(props[0]["id"].(string))
	}

	if len(props) > 0 && props[0]["name"] != nil {
		name = props[0]["name"].(string)
	}

	if len(props) > 0 && props[0]["email"] != nil {
		email = props[0]["email"].(string)
	}

	return &entity.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: fake.Internet().Password(),
	}
}
