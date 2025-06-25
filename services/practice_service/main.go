package main

import (
	"log"
	"practice_service/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "John Doe",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-9876-5432"},
		Project: map[string]*pb.Company_Project{
			"ProjectX": &pb.Company_Project{},
		},
		Profile: &pb.Employee_Text{
			Text: "my name is John Doe",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can't serialize", err)
	// }

	// if err := ioutil.WriteFile("employee.bin", binData, 0666); err != nil {
	// 	log.Fatalln("Can't write file", err)
	// }

	// in, err := ioutil.ReadFile("employee.bin")
	// if err != nil {
	// 	log.Fatalln("Can't read file", err)
	// }

	// readEmployee := &pb.Employee{}
	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Can't deserialize", err)
	// }

	// fmt.Println(readEmployee)

	out, err := protojson.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize to JSON", err)
	}
	// fmt.Printf("%s\n", out)

	readEmployee := &pb.Employee{}
	if err := protojson.Unmarshal(out, readEmployee); err != nil {
		log.Fatalln("Can't deserialize from JSON", err)
	}
	log.Printf("Employee: %s\n", readEmployee)
}
