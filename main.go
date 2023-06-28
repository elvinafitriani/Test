package main

import (
	"fmt"
	"net/http"
	"test/controller"
	"test/model"
)

func sum(a int, b int) int {
	sum := a + b
	return sum
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	//soal 1
	var angka1, angka2 int
	fmt.Print("Input angka pertama : ")
	fmt.Scan(&angka1)
	fmt.Print("Input angka kedua : ")
	fmt.Scan(&angka2)
	fmt.Println("Jumlah : ", sum(angka1, angka2))

	//soal 2
	http.HandleFunc("/hello/", controller.Salam)

	//soal 4
	fmt.Printf("/n/nSoal 4")
	var n int
	result := []model.User{}
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&n)

	users := make([]model.User, n)

	for i := 0; i < n; i++ {
		fmt.Println("Masukkan data user ke-", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&users[i].ID)
		fmt.Print("Name: ")
		fmt.Scan(&users[i].Name)
		fmt.Print("Email: ")
		fmt.Scan(&users[i].Email)
		fmt.Print("Umur: ")
		fmt.Scan(&users[i].Age)

		if users[i].Age > 18 {
			result = append(result, users[i])
		}
	}

	fmt.Println("Users who are above 18 years old:")
	for _, result := range result {
		fmt.Println("ID:", result.ID)
		fmt.Println("Name:", result.Name)
		fmt.Println("Email:", result.Email)
		fmt.Println("Age:", result.Age)
		fmt.Println()
	}

	//soal 5
	http.HandleFunc("/getProduct", controller.GetProduct)

	//soal 6
	http.HandleFunc("/product", controller.PostProduct)

	//soal 7
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	var target int
	fmt.Print("Angka yang dicari: ")
	fmt.Scan(&target)
	index := binarySearch(nums, target)
	if index != -1 {
		fmt.Printf("Angka %d ditemukan di indeks %d\n", target, index)
	} else {
		fmt.Println(index)
	}

	fmt.Println("Running Service")
	if err := http.ListenAndServe(":5000", nil); err != nil { //return listenandserve bertipe error
		fmt.Println("Error Starting Service")
	}

}
