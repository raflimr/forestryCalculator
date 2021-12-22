# Program Perhitungan Kehutanan

- [Installing](#installing)
- [Example](#example)
- [Details](#details)
- [License](#license)

## Installing
    
1. Install [Golang](https://go.dev/doc/install)
2. Install [Visual Studio Code](https://code.visualstudio.com/)
3. Install extension Golang di Visual Studio Code (VsCode)


## Example
1. Buat file di Visual Studio Code misal main.go
2. Lalu buat go modules misal nama modulus ambilLibrary 
    (diketik di terminal - untuk new terminal bisa menggunakan shotcut ctrl + shift + `)    
    ```
    go mod init ambilLibrary
    ```
3. Install Library
    (diketik di terminal - untuk new terminal bisa menggunakan shotcut ctrl + shift + `)
    ```
    go get github.com/rafllimr/forestryCalculator
    ```
4. Isi file yang telah dibuat(main.go)
    ```
    package main

    import (
        "fmt"
        "github.com/rafllimr/forestryCalculator"
    )

    func main(){
        var hitungDecimalDegree = forestryCalculator.DMS(108, 10, 20)
	    fmt.Println(hitungDMS)
        
    	var cariZonaUtm = forestryCalculator.FindUTM(108, 12, 25, "LS", "BT")
	    fmt.Println(cariZonaUtm)
    }
    ```
5. Jalankan program
    ```
    go run main.go
    ```
## Details
See Go Package documentation for more details: https://pkg.go.dev/github.com/rafllimr/forestryCalculator

## License
Licensed under MIT License.
See the LICENSE file.