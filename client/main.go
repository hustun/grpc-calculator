package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpc_simple "hasanustun/grpc-simple/calculator"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func calculate(expr string, c grpc_simple.ChatServiceClient) string {
	var stack []string

	expr = reverse(expr)

	for _, char := range expr {
		if strings.Contains("+-*/", string(char)) {
			n := len(stack) - 1
			top := stack[n]
			stack = stack[:n]

			n = len(stack) - 1
			top2 := stack[n]
			stack = stack[:n]

			topFloat, _ := strconv.ParseFloat(top, 32)
			topFloat2, _ := strconv.ParseFloat(top2, 32)
			topFloat32 := float32(topFloat)
			topFloat32_2 := float32(topFloat2)

			// fmt.Printf("%f %f\n", topFloat32, topFloat32_2)

			switch char {
			case '+':
				response, err := c.Add(context.Background(), &grpc_simple.CalcRequest{Num1: topFloat32, Num2: topFloat32_2})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				stack = append(stack, fmt.Sprintf("%f", response.Res))
			case '-':
				response, err := c.Subtract(context.Background(), &grpc_simple.CalcRequest{Num1: topFloat32, Num2: topFloat32_2})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				stack = append(stack, fmt.Sprintf("%f", response.Res))
			case '*':
				response, err := c.Multiply(context.Background(), &grpc_simple.CalcRequest{Num1: topFloat32, Num2: topFloat32_2})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				stack = append(stack, fmt.Sprintf("%f", response.Res))
			case '/':
				response, err := c.Divide(context.Background(), &grpc_simple.CalcRequest{Num1: topFloat32, Num2: topFloat32_2})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				stack = append(stack, fmt.Sprintf("%f", response.Res))
			}
		} else {
			stack = append(stack, string(char))
		}
	}

	return stack[0]
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := grpc_simple.NewChatServiceClient(conn)

	// response, err := c.Add(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %f", response.Res)

	// responseSub, err := c.Subtract(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %f", responseSub.Res)

	// responseMul, err := c.Multiply(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %f", responseMul.Res)

	// responseDiv, err := c.Divide(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %f", responseDiv.Res)

	fmt.Println(calculate("-+7*45+20", c))
	fmt.Println(calculate("-+8/632", c))
	fmt.Println(calculate("+1*2/3-45", c))
	fmt.Println(calculate("+1*2/3-55", c))
	fmt.Println(calculate("-1/40", c))
}
