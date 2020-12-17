package grpc

// var conn *grpc.ClientConn

// func init() {
// 	var err error
// 	conn, err = grpc.Dial(":8001", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %s", err)
// 	}
// }

// func TestUser_Store(t *testing.T) {
// 	client := NewUserServiceClient(conn)
// 	ctx := context.Background()
// 	resp, err := client.Store(ctx, &ReqCreateUser{
// 		Username: "thedevsaddam",
// 		Password: "12344",
// 		Type:     "admin",
// 		Profile: &Profile{
// 			Name: "Saddam H",
// 			Age:  30,
// 			Bio:  "This is bio...",
// 		},
// 	})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// 	spew.Dump(resp)
// }

// func TestUser_FetchUsers(t *testing.T) {
// 	client := NewUserServiceClient(conn)
// 	ctx := context.Background()
// 	resp, err := client.FetchUsers(ctx, &ReqFetchUsers{})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// 	spew.Dump(resp)
// }
