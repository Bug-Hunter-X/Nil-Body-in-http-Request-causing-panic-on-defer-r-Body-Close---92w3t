func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    if r.Body != nil {
        defer r.Body.Close()
    }
    // ... other code ...
} 