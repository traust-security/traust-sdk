package web

// htmx swaps error responses too, so a refusal shows where its request was made.
const htmxConfig = `{"responseHandling":[{"code":"204","swap":false},{"code":"...","swap":true}]}`
