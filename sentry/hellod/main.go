package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

func inner_handle() {
	panic("Don't panic!")
}

func handle() {
	inner_handle()
}

func main() {
	sentry_token := os.Getenv("SENTRY_TOKEN")
	if sentry_token == "" {
		fmt.Println("SENTRY_TOKEN environment variable not set")
		os.Exit(1)
	}

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	dsn := fmt.Sprintf("https://%s@o4506002231394304.ingest.sentry.io/4506002233163776", sentry_token)
	if err := sentry.Init(sentry.ClientOptions{
		// The DSN to use. If the DSN is not set, the client is effectively
		// disabled.
		Dsn: dsn,
		// ********************* Error capturing options *********************
		// The sample rate for event submission in the range [0.0, 1.0]. By default,
		// all events are sent. Thus, as a historical special case, the sample rate
		// 0.0 is treated as if it was 1.0. To drop all events, set the DSN to the
		// empty string.
		SampleRate: 1.0,
		// Configures whether SDK should generate and attach stack traces to pure
		// capture message calls.
		AttachStacktrace: true,

		// ********************* Performance options *********************
		// Enable performance tracing
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,

		// ********************* Profiling options *********************
		// The sampling rate for profiling is relative to TracesSampleRate:
		ProfilesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	// Once it's done, you can set up routes and attach the handler as one of your middleware
	// http.Handle("/", sentryHandler.Handle(&handler{}))
	http.HandleFunc("/foo", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		panic("foo panic")
	}))

	http.HandleFunc("/bar", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		handle()
	}))

	http.HandleFunc("/fast", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		t := time.Duration(100+rand.Intn(100)) * time.Millisecond
		time.Sleep(t)
	}))

	http.HandleFunc("/slow", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		t := time.Duration(100+rand.Intn(1000)) * time.Millisecond
		time.Sleep(t)
	}))

	http.HandleFunc("/prof", sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		n, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil || n <= 0 {
			http.NotFound(rw, r)
			return
		}
		f1 := fib1(n)
		f2 := fib2(n)
		fmt.Fprintf(rw, "fib1: %d, fib2: %d", f1, f2)
	}))

	fmt.Println("Listening and serving HTTP on :3000")

	// And run it
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

// Fibonacci functions for profiling example
func fib1(n int) int {
	if n < 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if n < 3 {
		return 1
	}
	f1 := 1
	f2 := 1
	res := 0
	for n > 2 {
		res = f1 + f2
		f1 = f2
		f2 = res
		n -= 1
	}
	return res
}
