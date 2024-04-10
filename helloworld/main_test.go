package helloworld 

import "testing"

func TestHello(t *testing.T) {
    t.Run("Saying hello to people in English", func(t *testing.T){
        got := Hello("Chris", "English")
        want := "Hello, Chris"
        assertMessage(t, got, want)
    })

    t.Run("Saying hello to people in Spanish", func(t *testing.T){
        got := Hello("Chris", "Spanish")
        want := "Hola, Chris"
        assertMessage(t, got, want)
    })
    t.Run("Saying hello to people in French", func(t *testing.T){
        got := Hello("Chris", "French")
        want := "Bonjour, Chris"
        assertMessage(t, got, want)
    })
}

func assertMessage(t testing.TB, got, want string){
    t.Helper()
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
