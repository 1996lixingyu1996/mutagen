	// the path given as the first argument with the digest specified by the
	// second argument.
	Provide(path string, digest []byte) (string, error)
func (t *transitioner) ensureExpectedFile(path string, expected *Entry) (os.FileMode, int, int, error) {
		return 0, 0, 0, errors.New("unable to find cache information for path")
		return 0, 0, 0, errors.Wrap(err, "unable to grab file statistics")
		return 0, 0, 0, errors.Wrap(err, "unable to convert cached modification time format")
		return 0, 0, 0, errors.New("modification detected")
	}

	// Extract ownership.
	uid, gid, err := getOwnership(info)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "unable to compute file ownership")
	return mode, uid, gid, nil
	if _, _, _, err := t.ensureExpectedFile(path, target); err != nil {
	mode, uid, gid, err := t.ensureExpectedFile(path, oldEntry)
	// Compute the new file mode based on the new entry's executability.
	if newEntry.Executable {
		mode = MarkExecutableForReaders(mode)
	} else {
		mode = StripExecutableBits(mode)
	}
	// If both files have the same contents (differing only in permissions),
	// then we won't have staged the file, so we just change the permissions on
	// the existing file.
	if bytes.Equal(oldEntry.Digest, newEntry.Digest) {
	stagedPath, err := t.provider.Provide(path, newEntry.Digest)
	// Set the mode for the staged file.
	if err := os.Chmod(stagedPath, mode); err != nil {
		return errors.Wrap(err, "unable to set staged file mode")
	}

	// Set the ownership for the staged file.
	if err := setOwnership(stagedPath, uid, gid); err != nil {
		return errors.Wrap(err, "unable to set staged file ownership")
	}

	stagedPath, err := t.provider.Provide(path, target.Digest)
	// Compute the new file mode based on the new entry's executability.
	mode := newFileBaseMode
	if target.Executable {
		mode = MarkExecutableForReaders(mode)
	} else {
		mode = StripExecutableBits(mode)
	}

	// Set the mode for the staged file.
	if err := os.Chmod(stagedPath, mode); err != nil {
		return errors.Wrap(err, "unable to set staged file mode")
	}
