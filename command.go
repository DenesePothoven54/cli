package cli

// ... existing code ...

func (c *Command) inheritFlags(parent *Command) {
	// Existing inheritance for standard flags
	for _, f := range parent.Flags {
		if f.Persistent {
			c.Flags = append(c.Flags, f)
		}
	}

	// New inheritance for MutuallyExclusiveFlags
	for _, group := range parent.MutuallyExclusiveFlags {
		newGroup := MutuallyExclusiveFlags{
			Name:  group.Name,
			Flags: []Flag{},
		}
		for _, f := range group.Flags {
			if f.Persistent {
				newGroup.Flags = append(newGroup.Flags, f)
				c.Flags = append(c.Flags, f)
			}
		}
		if len(newGroup.Flags) > 0 {
			c.MutuallyExclusiveFlags = append(c.MutuallyExclusiveFlags, newGroup)
		}
	}
}

// ... existing code ...