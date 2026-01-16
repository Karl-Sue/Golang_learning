// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident {
        Name: name,
        Age: age,
        Address: address,
    }
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address == nil || len(r.Address) == 0 {
        return false
    } 
    val, ok := r.Address["street"]
    
    if !ok {
        return false
    } else if val == "" {
        return false
    }
    
    return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    count := 0
	for _, val := range residents {
        if val.HasRequiredInfo() {
            count++
        }
    }
    return count 
}
