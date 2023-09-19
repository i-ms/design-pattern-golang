package observer

type ProfileList struct {
	profiles []Profile
}

func NewProfileList() *ProfileList {
	return &ProfileList{
		profiles: make([]Profile, 0),
	}
}

func (f *ProfileList) Add(p Profile) {
	f.profiles = append(f.profiles, p)
}

func (f *ProfileList) Remove(p Profile) {
	for i, v := range f.profiles {
		if v == p {
			f.profiles = append(f.profiles[:i], f.profiles[i+1:]...)
		}
	}
}
