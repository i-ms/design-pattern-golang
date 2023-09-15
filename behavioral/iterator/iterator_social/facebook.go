package iteratorsocial

type Facebook struct{}

func (f *Facebook) CreateFriendsIterator(profileID int) ProfileIterator {
	return &FacebookIterator{
		network:     f,
		profileID:   profileID,
		profileType: "Friend",
	}
}

func (f *Facebook) CreateCoworkersIterator(profileID int) ProfileIterator {
	return &FacebookIterator{
		network:     f,
		profileID:   profileID,
		profileType: "Coworker",
	}
}

type FacebookIterator struct {
	network      *Facebook
	profileID    int
	profileType  string
	currentIndex int
	profiles     []Profile
}

func (fi *FacebookIterator) Next() Profile {
	profile := fi.profiles[fi.currentIndex]
	fi.currentIndex++
	return profile
}

func (fi *FacebookIterator) HasNext() bool {
	fi.lazyInit()
	return fi.currentIndex < len(fi.profiles)
}

func (fi *FacebookIterator) lazyInit() {
	if fi.profiles == nil {
		// Simulate fetfhing profiles from the social network.
		// In a real application, you would make a network request here.
		fi.profiles = []Profile{
			{ID: 1, Name: "John Doe", Email: "john@fb.com"},
			{ID: 2, Name: "Scarlette Johanson", Email: "scarlet@fb.com"},
			{ID: 3, Name: "Mark Zuckerberg", Email: "mark@fb.com"},
		}
	}
}
