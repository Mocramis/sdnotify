// +build !linux

package sdnotify


func Reloading() error {
	return nil
}

// SdNotify sends a specified string to the systemd notification socket.
func SdNotify(state string) error {
	// do nothing
	return nil
}
