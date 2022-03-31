package infrastructure

import "fmt"

const (
	Doctor UserType = iota
	Patient
	ClinicAdmin
)

func (u *UserType) String() (string, error) {
	switch *u {
	case Doctor:
		return "Doctor", nil
	case Patient:
		return "Patient", nil
	case ClinicAdmin:
		return "ClinicAdmin", nil
	default:
		return "", fmt.Errorf("invalid usertype")
	}
}

func (u *UserType) FromString(val string) error {
	switch val {
	case "Doctor":
		*u = Doctor
	case "Patient":
		*u = Patient
	case "ClinicAdmin":
		*u = ClinicAdmin
	default:
		return fmt.Errorf("invalid usertype")
	}
	return nil
}

func (u *UserType) MarshalJSON() ([]byte, error) {
	x, y := u.String()
	return []byte(x), y
}
