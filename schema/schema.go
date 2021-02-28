package schema

import (
	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/photoShopPlugin/model"
	"github.com/graphql-go/graphql"
)

var proJSchema *funplugin.ObjectSchema
var pageSchema *funplugin.ObjectSchema
var load = false

func Init() {
	// InitAccount()
	//if employee, ok := plugin.GetObject("employee"); !ok {
	//	panic(errors.New("not have object type"))
	//} else {
	//	photoShop := graphql.NewObject(graphql.ObjectConfig{
	//		Name:        "employee_photoShop",
	//		Description: "employee_photoShop plugin",
	//		Fields: graphql.Fields{
	//			"photoShops": applyformSchema.Query["photoShops"],
	//		},
	//	})
	//
	//	employee.AddFieldConfig("photoShop", &graphql.Field{
	//		Name:        "queryemployee_storeroom",
	//		Description: "queryemployee_storeroom",
	//		Type:        photoShop,
	//		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	//			return params.Source, nil
	//		},
	//	})
	//}
}

func marge(oc *funplugin.ObjectSchema) {
	for k, v := range oc.Query {
		queryFields[k] = v
	}
	for k, v := range oc.Mutation {
		mutationFields[k] = v
	}
}

var queryFields = graphql.Fields{
	// "account":  &queryAccount,
	// "accounts": &queryAccountList,
	// "authority":  &queryAuthority,
	// "authoritys": &queryAuthorityList,
}

var mutationFields = graphql.Fields{
	// "createAccount": &createAccount,
	// "updateAccount": &updateAccount,
}

// NewSchema 用于插件主程序调用
func NewPlugSchema(pls funplugin.PluginManger) funplugin.Schema {
	if load != true {
		proJSchema, _ = pls.NewSchemaBuilder(model.ProJ{})
		marge(proJSchema)

		pageSchema, _ = pls.NewSchemaBuilder(model.Page{})
		marge(pageSchema)
		load = true
	}

	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			"proJ": proJSchema.GraphQLType,
			"page": pageSchema.GraphQLType,
			// "account": accountType,
			// "role":        roleSchema.GraphQLType,
			// "roleaccount": roleAccountSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
	}
}
