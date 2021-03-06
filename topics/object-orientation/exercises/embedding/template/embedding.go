// Create a user type, and an admin type that embeds a user. Create a notifier
// interface, and make your user type satisfy that interface. Write a function
// that accepts a value of the interface type, and ensure it works correctly
// when passed a value of your admin type.
//
// Template available at: http://play.golang.org/p/5qrrcfHdiZ
package main

// Add your imports here.
import (
	"fmt"
)

// Define a `notifier` interface.
type notifier interface {
	notify()
}

// Create a `user` type, with fields for name and email address.  Ensure your
// type satisfies the Notifier interface.
type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s at %s", u.name, u.email)
}

// Create an `admin` type which embeds a user, and has a security level.
type admin struct {
	user
	securityLevel int
}

// Write a function that accepts a value of your interface and calls the method
// associated with that interface.
func notify(notifier notifier) {
	notifier.notify()
}

func main() {
	// Create an admin user.
	admin := admin{user{"Mike", "Mike@gmail.com"}, 5}

	// Send the admin a notification via the function you created.
	notify(admin)
}
