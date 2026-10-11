package main
import ("encoding/json"; "fmt")
func main() {
    // Entirely synthetic responses, never obtained from a real account.
    var sample any
    err := json.Unmarshal([]byte(`{"data":{"data":[{"id":"sample-a","note":"Fictional event"}],"last_page":"3"}}`), &sample)
    if err != nil { panic(err) }
    fmt.Printf("Events: %d; last page: %d\n", len(findOccurrenceArray(sample)), findLastPage(sample))
}
