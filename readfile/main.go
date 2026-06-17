package main

// import (
//     "encoding/csv"
//     "fmt"
//     "os"
// )

// func main() {
//     path := "products.csv"
    
//     // Открываем файл для чтения и записи, создаем если не существует
//     file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
//     if err != nil {
//         fmt.Println("Error opening file:", err)
//         return
//     }
//     defer file.Close()

//     writer := csv.NewWriter(file)
    
//     data := [][]string{
//         {"Product", "Price", "Quantity"},
//         {"Apple", "1.00", "10"},
//         {"Banana", "2.00", "20"},
//         {"Cherry", "3.00", "30"},
//     }
    
//     err = writer.WriteAll(data)
//     if err != nil {
//         fmt.Println("Error writing to file:", err)
//         return
//     }
    
//     fmt.Println("Data has been written to the file.")
// }