func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    defer r.Body.Close() //This line might cause error if r.Body is nil
    // ... other code ...
}