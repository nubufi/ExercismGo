package erratum

// Use opens a resource and frobs input
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	// Open resource
	for r, err = o(); err != nil; r, err = o() {
		if _, ok := err.(TransientError); ok {
			continue
		}
		return err
	}
	defer r.Close()

	defer func() {
		if e := recover(); e != nil {
			if f, ok := e.(FrobError); ok {
				r.Defrob(f.defrobTag)
			}
			err = e.(error)
		}
	}()

	r.Frob(input)
	return err
}
