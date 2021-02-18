package schema

// funservicemodel "funservice/model"

// var deviceActType = graphql.NewObject(graphql.ObjectConfig{
// 	Name:        "DeviceAct",
// 	Description: "deviceact",
// 	Fields: graphql.Fields{
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 	},
// })

// var deviceType = graphql.NewObject(graphql.ObjectConfig{
// 	Name:        "Device",
// 	Description: "Device Service Model",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.ID,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				account := p.Source.(model.Device)
// 				return "account-" + strconv.FormatUint(uint64(account.ID), 10), nil
// 			},
// 		},
// 		"username": &graphql.Field{
// 			Type:        graphql.String,
// 			Description: "用户账号",
// 		},
// 		"referid": &graphql.Field{
// 			Type: graphql.ID,
// 		},
// 		"refertype": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"status": &graphql.Field{
// 			Type:        StatusEnumType,
// 			Description: "账号状态",
// 		},
// 		"remark": &graphql.Field{
// 			Type:        graphql.String,
// 			Description: "备注",
// 		},
// 		"updatedAt": &graphql.Field{
// 			Type: graphql.DateTime,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				obj := p.Source.(model.Device)
// 				return obj.UpdatedAt, nil
// 			},
// 		},
// 		"deletedAt": &graphql.Field{
// 			Type: graphql.DateTime,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				obj := p.Source.(model.Device)
// 				return obj.DeletedAt, nil
// 			},
// 		},
// 		"createdAt": &graphql.Field{
// 			Type: graphql.DateTime,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				obj := p.Source.(model.Device)
// 				return obj.CreatedAt, nil
// 			},
// 		},
// 		"action": &graphql.Field{
// 			Type: accountActType,
// 			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 				if params.Info.Operation.GetOperation() != "mutation" {
// 					return nil, errors.New("call by update mutation")
// 				}

// 				v := params.Source.(model.Device)
// 				return v, nil
// 			},
// 		},
// 	},
// })

// func init() {

// }

// func InitAccount() {
// 	union, err := plugin.AutoField("employee,admin")
// 	if err != nil {
// 		panic(errors.New("not have object type"))
// 	}
// 	accountType.AddFieldConfig("user", union)

// 	accountActType.AddFieldConfig("updatePassWord", &UpdatePassWord)
// }

// var createAccount = graphql.Field{
// 	Name:        "createAccount",
// 	Description: "Create Device",
// 	Type:        accountType,
// 	Args: graphql.FieldConfigArgument{
// 		"username": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"password": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"status": &graphql.ArgumentConfig{
// 			Type: StatusEnumType,
// 		},
// 		"refertype": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"referid": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 		"roles": &graphql.ArgumentConfig{
// 			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
// 		},
// 		"remark": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		username := p.Args["username"].(string)
// 		password := p.Args["password"].(string)

// 		var status model.STATUS_TYPE
// 		if p.Args["status"] != nil {
// 			status = p.Args["status"].(model.STATUS_TYPE)
// 		}

// 		var refertype string
// 		if p.Args["refertype"] != nil {
// 			refertype = p.Args["refertype"].(string)
// 		}

// 		var referid uint
// 		if p.Args["referid"] != nil {
// 			referid = uint(p.Args["referid"].(int))
// 		}

// 		var remark string
// 		if p.Args["remark"] != nil {
// 			remark = p.Args["remark"].(string)
// 		}

// 		obj := &model.Device{}
// 		err := obj.Create(username, password, status, refertype, referid, remark)
// 		return *obj, err
// 	},
// }

// var queryAccount = graphql.Field{
// 	Name:        "queryAccount",
// 	Description: "Get Device Service by id",
// 	Type:        accountType,
// 	Args: graphql.FieldConfigArgument{
// 		"id": &graphql.ArgumentConfig{
// 			Type: graphql.ID,
// 		},
// 		"refertype": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"referid": &graphql.ArgumentConfig{
// 			Type: graphql.ID,
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		if p.Args["id"] != nil {
// 			id, err := funplugin.ID2id(p.Args["id"])
// 			if err != nil {
// 				return nil, err
// 			}
// 			obj := &model.Device{}
// 			err = obj.QueryById(id)
// 			return *obj, err
// 		} else if p.Args["refertype"] != nil && p.Args["referid"] != nil {
// 			id, err := funplugin.ID2id(p.Args["referid"])
// 			if err != nil {
// 				return nil, err
// 			}
// 			refer := p.Args["refertype"].(string)
// 			obj := &model.Device{}
// 			err = obj.QueryByRefer(refer, id)
// 			return *obj, err
// 		}
// 		return nil, errors.New("Params Error")
// 	},
// }

// var accountConnection = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "AccountConnection",
// 	Fields: graphql.Fields{
// 		"totalCount": &graphql.Field{
// 			Type: graphql.NewNonNull(graphql.Int),
// 		},
// 		"edges": &graphql.Field{
// 			Type: graphql.NewList(accountType),
// 		},
// 	},
// })

// var queryAccountList = graphql.Field{
// 	Name:        "queryAccountList",
// 	Description: "Get Device List Service by id",
// 	Type:        accountConnection,
// 	Args: graphql.FieldConfigArgument{
// 		"first": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 		"skip": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 		"role": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		first := 10
// 		if p.Args["first"] != nil {
// 			first = p.Args["first"].(int)
// 		}

// 		skip := 0
// 		if p.Args["skip"] != nil {
// 			skip = p.Args["skip"].(int)
// 		}

// 		role := ""
// 		if p.Args["role"] != nil {
// 			role = p.Args["role"].(string)
// 		}

// 		obj := &model.Device{}
// 		return obj.QueryByList(first, skip, role)
// 	},
// }

// var updateAccount = graphql.Field{
// 	Name:        "updateAccount",
// 	Description: "update Device",
// 	Type:        accountType,
// 	Args: graphql.FieldConfigArgument{
// 		"id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.ID),
// 		},
// 		"username": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"password": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"status": &graphql.ArgumentConfig{
// 			Type: StatusEnumType,
// 		},
// 		"refertype": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"referid": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 		"roles": &graphql.ArgumentConfig{
// 			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
// 		},
// 		"remark": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		id, err := funplugin.ID2id(p.Args["id"])
// 		if err != nil {
// 			return nil, err
// 		}

// 		// @todo check session
// 		// session := p.Context.Value("session").(*model.Session)
// 		// if id != session.ID {
// 		// 	return nil, errors.New("id error")
// 		// }

// 		obj := &model.Device{}
// 		obj.QueryById(id)

// 		if p.Args["roles"] != nil {
// 			list, ok := p.Args["roles"].([]interface{})
// 			if !ok {
// 				return nil, errors.New("error roles type")
// 			}

// 			var roles []string
// 			roles = make([]string, len(list))
// 			for i, r := range list {
// 				roles[i] = r.(string)
// 			}

// 			// authority := &model.Authority{}
// 			// id := strconv.FormatUint(uint64(obj.ID), 10)
// 			// err := authority.SetGroupingPolicy(id, roles)
// 			// if err != nil {
// 			// 	return nil, err
// 			// }
// 		}

// 		var username string
// 		if p.Args["username"] != nil {
// 			username = p.Args["username"].(string)
// 		}
// 		var password string
// 		if p.Args["password"] != nil {
// 			password = p.Args["password"].(string)
// 		}

// 		var status model.STATUS_TYPE
// 		if p.Args["status"] != nil {
// 			status = p.Args["status"].(model.STATUS_TYPE)
// 		}

// 		var refertype string
// 		if p.Args["refertype"] != nil {
// 			refertype = p.Args["refertype"].(string)
// 		}

// 		var referid uint
// 		if p.Args["referid"] != nil {
// 			referid = uint(p.Args["referid"].(int))
// 		}

// 		var remark string
// 		if p.Args["remark"] != nil {
// 			remark = p.Args["remark"].(string)
// 		}

// 		err = obj.Update(id, username, password, status, refertype, referid, remark)
// 		return *obj, err
// 	},
// }

// var UpdatePassWord = graphql.Field{
// 	Name:        "updatePassWord",
// 	Description: "update PassWord",
// 	Type:        accountType,
// 	Args: graphql.FieldConfigArgument{
// 		"oldpassword": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 		"newpassword": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		obj := p.Source.(model.Device)

// 		oldpassword := p.Args["oldpassword"].(string)
// 		newpassword := p.Args["newpassword"].(string)

// 		err := obj.UpdatePassWord(oldpassword, newpassword)
// 		return obj, err
// 	},
// }
