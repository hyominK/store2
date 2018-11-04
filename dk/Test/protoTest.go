package main

import (
	"context"
	ed "dk/protoTest/protobuf"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func newEduData(name string, id string, email string, list []*ed.EduData_Subject) ed.EduData {
	data := ed.EduData{
		Name:     name,
		Id:       id,
		Email:    email,
		Subjects: list,
	}
	return data
}

const (
	address = "localhost:8000"
	//defaultName = "world"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ed.NewSetDataClient(conn)

	//객체
	p1 := ed.EduData{
		Name:  "kim",
		Id:    "1234a342kjdfskmfdslkjfdslkj",
		Email: "d@d.d",
		Subjects: []*ed.EduData_Subject{
			{SubjectName: "TOEIC", SubjectRank: "990", Type: ed.EduData_ENGLISH},
			{SubjectName: "GISA", SubjectRank: "PASS", Type: ed.EduData_COMPUTER},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r1, err := c.SetData(ctx, &p1)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r1.Name)
	log.Printf(r1.Id)
	log.Printf(r1.Email)
	for i := 0; i < 2; i++ {
		fmt.Println("subject", i, ": ", p1.GetSubjects()[i].GetSubjectName())

	}
	//객체 생성
	ddd := []*ed.EduData_Subject{
		{SubjectName: "test1", SubjectRank: "990", Type: ed.EduData_ENGLISH},
		{SubjectName: "test2", SubjectRank: "PASS", Type: ed.EduData_COMPUTER},
	}
	p2 := newEduData("호", "123", "asdf", ddd)
	fmt.Println(p2)
	r2, err := c.SetData(ctx, &p2)
	log.Printf(r2.Name)
	log.Printf(r2.Id)
	log.Printf(r2.Email)
	for i := 0; i < 2; i++ {
		fmt.Println("subject", i, ": ", p2.GetSubjects()[i].GetSubjectName())

	}
	// p3 := []*ed.EduData{&p1, &p2}
	// r3, err := c.SetHistory(ctx, p3)
	// //값추출
	// fmt.Println("Name: ", p.GetName())
	// //nested
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("subject", i, ": ", p.GetSubjects()[i].GetSubjectName())
	// }

}

// // Set up a connection to the server.
// conn, err := grpc.Dial(address, grpc.WithInsecure())
// if err != nil {
// 	log.Fatalf("did not connect: %v", err)
// }
// defer conn.Close()
// c := pb.NewGreeterClient(conn)

// // Contact the server and print out its response.
// name := defaultName
// if len(os.Args) > 1 {
// 	name = os.Args[1]
// }
// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// defer cancel()

// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
// if err != nil {
// 	log.Fatalf("could not greet: %v", err)
// }
// log.Printf("Greeting: %s", r.Message)
