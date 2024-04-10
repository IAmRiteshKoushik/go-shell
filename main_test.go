package main

import "testing"

func TestHello(t *testing.T){
    // Writing sub-tests
    // Test case 1
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        // if got != want {
        //     t.Errorf("got %q want %q", got, want)
        // }
        assertCorrectMessage(t, got, want)
    })
    // Test case 2
    t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T){
        got := Hello("")
        want := "Hello, World"

        // if got != want {
        //     t.Errorf("got %q want %q", got, want)
        // }
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper() // Tells the test suite that this method is a helper

    // When a test case fails, then the test suite reports back the  
    // line number of the line what is failing. Now, if we do not 
    // specify t.Helper() then the lines that follow will be reported
    // for error instead of the lines from the main() file as this will 
    // be treated as any ordinary function and not a testing-helper func

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
