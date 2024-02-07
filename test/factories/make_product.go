package factories

import (
	"github.com/jaswdr/faker"
	"github.com/nitoba/go-api/internal/domain/enterprise/entity"
	pkg "github.com/nitoba/go-api/pkg/entity"
)

func MakeProduct(props ...map[string]interface{}) *entity.Product {
	fake := faker.New()
	id := pkg.NewID()
	name := fake.Lorem().Word()
	price := fake.Float64(1, 0, 10000.0)

	if len(props) > 0 && props[0]["id"] != nil {
		id, _ = pkg.ParseID(props[0]["id"].(string))
	}

	if len(props) > 0 && props[0]["name"] != nil {
		name = props[0]["name"].(string)
	}

	if len(props) > 0 && props[0]["price"] != nil {
		price = props[0]["price"].(float64)
	}

	return &entity.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
