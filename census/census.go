// Package census simulates a system used to collect census data.
package census

// import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	r := Resident{Name: name, Age: age, Address:address}
    new_r := &r
    return new_r
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    known_key := "street"
	if r.Name != "" && r.Address != nil {
		val, ok := r.Address[known_key]
        if val != "" && ok {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
    r.Age = 0
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    tally := 0
	for _, resident := range residents {
        ok := resident.HasRequiredInfo()
        if ok {
            tally += 1
        }
    }

    return tally
}
