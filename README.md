# Go: Handling Nil Body in http.Request

This repository demonstrates a common error in Go's `net/http` package when dealing with request bodies.  Improper handling of potentially `nil` `http.Request.Body` can lead to panics, particularly when using `defer r.Body.Close()`.  The solution shows a safe way to handle this edge case, preventing unexpected crashes.

## Problem

The `defer r.Body.Close()` statement is commonly used to ensure that the request body is closed after processing.  However, under certain circumstances, such as malformed or empty requests, `r.Body` can be `nil`. Calling `Close()` on a `nil` value results in a panic.

## Solution

The recommended approach is to check for `nil` before closing the body.  This prevents the panic and handles the scenario gracefully.
