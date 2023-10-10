package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func getIntHash(val int) string {
	h := sha1.New()
	h.Write([]byte(strconv.Itoa(val)))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func getHash(val string) string {
	h := sha1.New()
	h.Write([]byte(val))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func buildTree(hashes []Node) Node {
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}
	secondary := []Node{}
	for i := 0; i < len(hashes); i = i + 2 {
		temp_hash := getHash(hashes[i].data + hashes[i+1].data)
		temp_node := Node{temp_hash, &hashes[i], &hashes[i+1]}
		secondary = append(secondary, temp_node)
	}

	if len(secondary) == 1 {
		return secondary[0]
	} else {
		return buildTree(secondary)
	}
}

func callBuildTree(t_hashes []Node) {
	merkle_root := buildTree(t_hashes)
	fmt.Println("---------------------------------------------")
	fmt.Println("\nMerkle Root Hash Value: ", merkle_root.data)
	fmt.Println("---------------------------------------------")
}

func main() {

	var transaction_count int
	fmt.Println("Enter the number of transactions: ")
	fmt.Scanln(&transaction_count)

	var transactions = make([]int, transaction_count)
	for i := 0; i < int(transaction_count); i++ {
		fmt.Scanf("%d", &transactions[i])
	}
	t_hashes := []Node{}
	for _, val := range transactions {
		temp_node := Node{getIntHash(val), nil, nil}
		t_hashes = append(t_hashes, temp_node)
	}

	callBuildTree(t_hashes)

	var choice int

	for {
		fmt.Println("\n1.Insert Transaction\n2.Delete Transaction\n3.Verify Transaction")
		fmt.Println("\nEnter the choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var temp_data int
			fmt.Print("Enter the transaction data: ")
			fmt.Scanf("%d", &temp_data)
			transactions = append(transactions, temp_data)
			temp_node := Node{getIntHash(temp_data), nil, nil}
			t_hashes = append(t_hashes, temp_node)
			fmt.Println("\nMerkle Root Updated")
			callBuildTree(t_hashes)
		case 2:
			var temp_val int
			fmt.Println("Enter the trasaction data: ")
			fmt.Scanf("%d", &temp_val)
			// deleteTransaction(&transactions)
		case 3:
			var value int
			fmt.Println("Enter the transaction data: ")
			fmt.Scanf("%d", &value)
			flag := 1
			for _, val := range transactions {
				if value == val {
					fmt.Println("Transaction is valid")
					flag = 0
				}
			}
			if flag == 1 {
				fmt.Println("Transaction is not valid")
			}
		}
	}

}
