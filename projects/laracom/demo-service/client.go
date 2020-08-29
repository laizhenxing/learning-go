package main

//func client() {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("连接到grpc服务器失败: %v", err)
//	}
//	defer conn.Close()
//
//	client := demo.NewDemoServiceClient(conn)
//	req := &demo.DemoRequest{
//		Name: "兴小狸",
//	}
//	rsp, err := client.SayHello(context.Background(), req)
//	if err != nil {
//		log.Fatalf("调用grpc服务接口失败：%v", err)
//	}
//	log.Printf("%s", rsp.Text)
//}
