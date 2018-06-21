package sync

import (
	"bytes"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
	pathpkg "path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pkg/errors"

	"github.com/havoc-io/mutagen/pkg/filesystem"
)

// testEntryDecomposer provides the implementation for testDecomposeEntry.
type testEntryDecomposer struct {
	// creation records whether or not the decomposition is for a creation
	// transition (as opposed to a removal transition).
	creation bool
	// transitions is the accumulated list of decomposed transitions.
	transitions []*Change
}

// decompose decomposes an entry recursively and records associated transitions.
func (d *testEntryDecomposer) decompose(path string, entry *Entry) {
	// If the entry is non-existent, then there are no transitions.
	if entry == nil {
		return
	}

	// Create a shallow copy of the entry.
	shallowEntry := entry.CopyShallow()

	// If this is a creation decomposition, add this entry before processing any
	// contents.
	if d.creation {
		d.transitions = append(d.transitions, &Change{Path: path, New: shallowEntry})
	}

	// If this is a directory, handle its contents.
	if entry.Kind == EntryKind_Directory {
		for name, entry := range entry.Contents {
			d.decompose(pathpkg.Join(path, name), entry)
		}
	}

	// If this is a removal decomposition, add this entry after processing any
	// contents.
	if !d.creation {
		d.transitions = append(d.transitions, &Change{Path: path, Old: shallowEntry})
	}
}

// testDecomposeEntry decomposes an entry into a sequence of transitions that
// can more granularly test Transition. Instead of calling Transition with a
// single change for creation/removal, testDecomposeEntry allows one to
// decompose an Entry into a single transition per node. It can perform
// decomposition for both creation and removal transitions.
func testDecomposeEntry(path string, entry *Entry, creation bool) []*Change {
	// Create a decomposer for creations.
	decomposer := &testEntryDecomposer{creation: creation}

	// Have it perform decomposition.
	decomposer.decompose(path, entry)

	// Return the relevant transitions.
	return decomposer.transitions
}

// testProvider is an implementation of the Provider interfaces for tests. It
// loads file data from a content map.
type testProvider struct {
	// servingRoot is the temporary directory where "staged" files are served
	// from.
	servingRoot string
	// contentMap is a map from path to file content.
	contentMap map[string][]byte
	// hasher is the hasher to use when verifying content.
	hasher hash.Hash
}

// newTestProvider creates a new instance of testProvider with the specified
// content map.
func newTestProvider(contentMap map[string][]byte, hasher hash.Hash) (*testProvider, error) {
	// Create a temporary directory for serving files.
	servingRoot, err := ioutil.TempDir("", "mutagen_provide_root")
	if err != nil {
		return nil, errors.Wrap(err, "unable to create serving directory")
	}

	// Create the test provider.
	return &testProvider{
		servingRoot: servingRoot,
		contentMap:  contentMap,
		hasher:      hasher,
	}, nil
}

// Provide implements the Provider interface for testProvider.
func (p *testProvider) Provide(path string, entry *Entry, baseMode os.FileMode) (string, error) {
	// Ensure the entry is a file type.
	if entry.Kind != EntryKind_File {
		return "", errors.New("invalid entry kind provision requested")
	}

	// Grab the content for this path.
	content, ok := p.contentMap[path]
	if !ok {
		return "", errors.New("unable to find content for path")
	}

	// Ensure it matches the requested hash.
	p.hasher.Reset()
	p.hasher.Write(content)
	if !bytes.Equal(entry.Digest, p.hasher.Sum(nil)) {
		return "", errors.New("requested entry digest does not match expected")
	}

	// Create a temporary file in the serving root.
	temporaryFile, err := ioutil.TempFile(p.servingRoot, "mutagen_provide")
	if err != nil {
		return "", errors.Wrap(err, "unable to create temporary file")
	}

	// Write content.
	_, err = temporaryFile.Write(content)
	temporaryFile.Close()
	if err != nil {
		os.Remove(temporaryFile.Name())
		return "", errors.Wrap(err, "unable to write file contents")
	}

	// Compute the file mode.
	mode := baseMode
	if mode == 0 {
		mode = newFileBaseMode
	}
	if entry.Executable {
		mode = MarkExecutableForReaders(mode)
	} else {
		mode = StripExecutableBits(mode)
	}

	// Set the file mode.
	if err := os.Chmod(temporaryFile.Name(), mode); err != nil {
		os.Remove(temporaryFile.Name())
		return "", errors.Wrap(err, "unable to set file mode")
	}

	// Success.
	return temporaryFile.Name(), nil
}

// finalize removes the temporary serving directory underlying the testProvider.
func (p *testProvider) finalize() error {
	return os.RemoveAll(p.servingRoot)
}

// testTransitionCreate creates test content on disk using Transition based on
// the specified entry and content map. It can optionally decompose the entry
// into individual node creations to stress-test Transition.
func testTransitionCreate(temporaryDirectory string, entry *Entry, contentMap map[string][]byte, decompose bool) (string, string, error) {
	// Create temporary directory to act as the parent of our root.
	parent, err := ioutil.TempDir(temporaryDirectory, "mutagen_simulated")
	if err != nil {
		return "", "", errors.Wrap(err, "unable to create temporary root parent")
	}

	// Determine whether or not we need to recompose Unicode when transitioning
	// inside this directory.
	recomposeUnicode, err := filesystem.DecomposesUnicode(parent)
	if err != nil {
		os.RemoveAll(parent)
		return "", "", errors.Wrap(err, "unable to determine Unicode decomposition behavior")
	}

	// Compute the path to the root.
	root := filepath.Join(parent, "root")

	// Set up the creation transitions.
	var transitions []*Change
	if decompose {
		transitions = testDecomposeEntry("", entry, true)
	} else {
		transitions = []*Change{{New: entry}}
	}

	// Create a provider and ensure its cleanup.
	provider, err := newTestProvider(contentMap, newTestHasher())
	if err != nil {
		os.RemoveAll(parent)
		return "", "", errors.Wrap(err, "unable to create test provider")
	}
	defer provider.finalize()

	// Perform the creation transition.
	if entries, problems := Transition(root, transitions, nil, SymlinkMode_SymlinkPortable, recomposeUnicode, provider); len(problems) != 0 {
		os.RemoveAll(parent)
		return "", "", errors.New("problems occurred during creation transition")
	} else if len(entries) != len(transitions) {
		os.RemoveAll(parent)
		return "", "", errors.New("unexpected number of entries returned from creation transition")
	} else {
		for e, entry := range entries {
			if !entry.Equal(transitions[e].New) {
				os.RemoveAll(parent)
				return "", "", errors.New("created entry does not match expected")
			}
		}
	}

	// Success.
	return root, parent, nil
}

// testTransitionRemove removes test content from disk using Transition based on
// the specified entry. It can optionally decompose the entry into individual
// node removals to stress-test Transition.
func testTransitionRemove(root string, expected *Entry, cache *Cache, symlinkMode SymlinkMode, decompose bool) error {
	// Set up the removal transitions.
	var transitions []*Change
	if decompose {
		transitions = testDecomposeEntry("", expected, false)
	} else {
		transitions = []*Change{{Old: expected}}
	}

	// If we're expecting to remove a directory, then determine the necessary
	// Unicode recomposition behavior.
	var recomposeUnicode bool
	if expected != nil && expected.Kind == EntryKind_Directory {
		if r, err := filesystem.DecomposesUnicode(root); err != nil {
			return errors.Wrap(err, "unable to determine Unicode decomposition behavior")
		} else {
			recomposeUnicode = r
		}
	}

	// Perform the removal transition.
	if entries, problems := Transition(root, transitions, cache, symlinkMode, recomposeUnicode, nil); len(problems) != 0 {
		return errors.New("problems occurred during removal transition")
	} else if len(entries) != len(transitions) {
		return errors.New("unexpected number of entries returned from removal transition")
	} else {
		for _, entry := range entries {
			if entry != nil {
				return errors.New("post-removal entry non-nil")
			}
		}
	}

	// Success.
	return nil
}

type testContentModifier func(string, *Entry) (*Entry, error)

func testTransitionCycle(temporaryDirectory string, entry *Entry, contentMap map[string][]byte, decompose bool, modifier testContentModifier) error {
	// Create test content on disk and defer its removal. This will exercise
	// the creation portion of Transition.
	root, parent, err := testTransitionCreate(temporaryDirectory, entry, contentMap, decompose)
	if err != nil {
		return errors.Wrap(err, "unable to create test content")
	}
	defer os.RemoveAll(parent)

	// Compute the expected entry.
	expected := entry

	// If a modifier has been specified, allow it to modify the disk contents
	// and expected result.
	if modifier != nil {
		if e, err := modifier(root, expected.Copy()); err != nil {
			return errors.Wrap(err, "modifier failed")
		} else {
			expected = e
		}
	}

	// Perform a scan.
	snapshot, preservesExecutability, _, cache, err := Scan(root, newTestHasher(), nil, nil, SymlinkMode_SymlinkPortable)
	if !preservesExecutability {
		snapshot = PropagateExecutability(nil, expected, snapshot)
	}
	if err != nil {
		return errors.Wrap(err, "unable to perform scan")
	} else if cache == nil {
		return errors.New("nil cache returned")
	} else if modifier == nil && !snapshot.Equal(expected) {
		return errors.New("snapshot not equal to expected")
	}

	// Remove the test content. This will exercise the removal portion of
	// Transition.
	if err := testTransitionRemove(root, expected, cache, SymlinkMode_SymlinkPortable, decompose); err != nil {
		return errors.Wrap(err, "unable to remove test content")
	}

	// Success.
	return nil
}

func testTransitionCycleWithPermutations(entry *Entry, contentMap map[string][]byte, modifier testContentModifier, expectFailure bool) error {
	// Loop over decomposition cases.
	for _, decompose := range []bool{false, true} {
		// Compute the composition case name.
		caseName := "composed"
		if decompose {
			caseName = "decomposed"
		}

		// Loop over temporary directories.
		for _, temporaryDirectory := range testingTemporaryDirectories() {
			err := testTransitionCycle(temporaryDirectory.path, entry, contentMap, decompose, modifier)
			if expectFailure && err == nil {
				return errors.Errorf("transition cycle succeeded unexpectedly in %s case for %s temporary directory", caseName, temporaryDirectory.name)
			} else if !expectFailure && err != nil {
				return errors.Wrap(err, fmt.Sprintf("transition cycle failed in %s case for %s temporary directory", caseName, temporaryDirectory.name))
			}
		}
	}

	// Success.
	return nil
}

func TestTransitionNilRoot(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testNilEntry, nil, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionFile1Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testFile1Entry, testFile1ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionFile2Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testFile2Entry, testFile2ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionFile3Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testFile3Entry, testFile3ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionDirectory1Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionDirectory2Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testDirectory2Entry, testDirectory2ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionDirectory3Root(t *testing.T) {
	if err := testTransitionCycleWithPermutations(testDirectory3Entry, testDirectory3ContentMap, nil, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionSwapFile(t *testing.T) {
	// Create a modifier function that will modify the case of a subpath and
	// attempt an additional create transition.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		// Perform a scan to grab Unicode recomposition behavior and a cache.
		_, _, recomposeUnicode, cache, err := Scan(root, newTestHasher(), nil, nil, SymlinkMode_SymlinkPortable)
		if err != nil {
			return nil, errors.Wrap(err, "unable to perform scan")
		} else if cache == nil {
			return nil, errors.New("nil cache returned")
		}

		// Attempt to switch the content of a file.
		transitions := []*Change{{
			Path: "file",
			Old:  testFile1Entry,
			New:  testFile2Entry,
		}}

		// Set up a custom content map for this.
		contentMap := map[string][]byte{
			"file": testFile2Contents,
		}

		// Create a provider and ensure its cleanup.
		provider, err := newTestProvider(contentMap, newTestHasher())
		if err != nil {
			return nil, errors.Wrap(err, "unable to create creation provider")
		}
		defer provider.finalize()

		// Perform the swap transition, ensure that it succeeds, and update the
		// expected contents.
		if entries, problems := Transition(root, transitions, cache, SymlinkMode_SymlinkPortable, recomposeUnicode, provider); len(problems) != 0 {
			return nil, errors.New("file swap transition failed")
		} else if len(entries) != 1 {
			return nil, errors.New("unexpected number of entries returned from swap transition")
		} else if !entries[0].Equal(testFile2Entry) {
			return nil, errors.New("file swap transition returned incorrect new file")
		} else {
			expected.Contents["file"] = entries[0]
		}

		// Success.
		return expected, nil
	}

	// Ensure that the whole cycle succeeds.
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, modifier, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionSwapFileOnlyExecutableChange(t *testing.T) {
	// Create a modifier function that will modify the case of a subpath and
	// attempt an additional create transition.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		// Perform a scan to grab Unicode recomposition behavior and a cache.
		_, _, recomposeUnicode, cache, err := Scan(root, newTestHasher(), nil, nil, SymlinkMode_SymlinkPortable)
		if err != nil {
			return nil, errors.Wrap(err, "unable to perform scan")
		} else if cache == nil {
			return nil, errors.New("nil cache returned")
		}

		// Create a copy of the current entry and mark it as executable.
		executableEntry := testFile1Entry.Copy()
		executableEntry.Executable = true

		// Attempt to switch the content of a file.
		transitions := []*Change{{
			Path: "file",
			Old:  testFile1Entry,
			New:  executableEntry,
		}}

		// Perform the swap transition with a nil provider (since it shouldn't
		// be used), ensure that it succeeds, and update the expected contents.
		if entries, problems := Transition(root, transitions, cache, SymlinkMode_SymlinkPortable, recomposeUnicode, nil); len(problems) != 0 {
			return nil, errors.New("file swap transition failed")
		} else if len(entries) != 1 {
			return nil, errors.New("unexpected number of entries returned from swap transition")
		} else if !entries[0].Equal(executableEntry) {
			return nil, errors.New("file swap transition returned incorrect new file")
		} else {
			expected.Contents["file"] = entries[0]
		}

		// Success.
		return expected, nil
	}

	// Ensure that the whole cycle succeeds.
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, modifier, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionCaseConflict(t *testing.T) {
	// Determine whether or not we expect case conflicts.
	// HACK: We actually ought to be determining this based on the filesystem
	// being used, but it's a sufficient test mechanism for now.
	expectCaseConflict := runtime.GOOS == "windows" || runtime.GOOS == "darwin"

	// Check for case conflicts.
	if err := testTransitionCycleWithPermutations(testDirectoryWithCaseConflict, testDirectoryWithCaseConflictContentMap, nil, expectCaseConflict); err != nil {
		t.Error("case conflict behavior not as expected:", err)
	}
}

func TestTransitionFailRemoveModifiedSubcontent(t *testing.T) {
	// Create a modifier function that will modify subcontent.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		if err := ioutil.WriteFile(filepath.Join(root, "file"), testFile3Contents, 0600); err != nil {
			return nil, errors.Wrap(err, "unable to modify file content")
		}
		return expected, nil
	}

	// Test that the removal fails.
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, modifier, true); err != nil {
		t.Error(err)
	}
}

func TestTransitionFailRemoveModifiedRootFile(t *testing.T) {
	// Create a modifier function that will modify the root.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		if err := ioutil.WriteFile(root, testFile3Contents, 0600); err != nil {
			return nil, errors.Wrap(err, "unable to modify file content")
		}
		return expected, nil
	}

	// Test that the removal fails.
	if err := testTransitionCycleWithPermutations(testFile1Entry, testFile1ContentMap, modifier, true); err != nil {
		t.Error(err)
	}
}

func TestTransitionFailCreateInvalidPathCase(t *testing.T) {
	// Create a modifier function that will modify the case of a subpath and
	// attempt an additional create transition.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		// Perform a scan to grab Unicode recomposition behavior and a cache.
		_, _, recomposeUnicode, cache, err := Scan(root, newTestHasher(), nil, nil, SymlinkMode_SymlinkPortable)
		if err != nil {
			return nil, errors.Wrap(err, "unable to perform scan")
		} else if cache == nil {
			return nil, errors.New("nil cache returned")
		}

		// Change the directory case.
		if err := os.Rename(filepath.Join(root, "directory"), filepath.Join(root, "directory-temp")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to temporary name")
		}
		if err := os.Rename(filepath.Join(root, "directory-temp"), filepath.Join(root, "DiRecTory")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to different case name")
		}

		// Attempt to create content inside the directory.
		transitions := []*Change{{Path: "directory/new", New: testFile1Entry}}

		// Set up a custom content map for this.
		contentMap := map[string][]byte{
			"directory/new": testFile1Contents,
		}

		// Create a provider and ensure its cleanup.
		provider, err := newTestProvider(contentMap, newTestHasher())
		if err != nil {
			return nil, errors.Wrap(err, "unable to create creation provider")
		}
		defer provider.finalize()

		// Perform the create transition and ensure that it fails.
		if entries, problems := Transition(root, transitions, cache, SymlinkMode_SymlinkPortable, recomposeUnicode, provider); len(problems) == 0 {
			return nil, errors.New("transition succeeded unexpectedly")
		} else if len(entries) != 1 {
			return nil, errors.New("unexpected number of entries returned from creation transition")
		} else if entries[0] != nil {
			return nil, errors.New("failed creation transition returned non-nil entry")
		}

		// Change the directory case back to normal.
		if err := os.Rename(filepath.Join(root, "DiRecTory"), filepath.Join(root, "directory-temp")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to temporary name")
		}
		if err := os.Rename(filepath.Join(root, "directory-temp"), filepath.Join(root, "directory")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to original name")
		}

		// Success.
		return expected, nil
	}

	// Ensure that the whole cycle succeeds (since our create will have failed
	// and we will have returned the directory to normal).
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, modifier, false); err != nil {
		t.Error(err)
	}
}

func TestTransitionFailRemoveInvalidPathCase(t *testing.T) {
	// Create a modifier function that will modify the case of a subpath.
	modifier := func(root string, expected *Entry) (*Entry, error) {
		if err := os.Rename(filepath.Join(root, "directory"), filepath.Join(root, "directory-temp")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to temporary name")
		}
		if err := os.Rename(filepath.Join(root, "directory-temp"), filepath.Join(root, "DiRecTory")); err != nil {
			return nil, errors.Wrap(err, "unable to rename directory to different case name")
		}
		return expected, nil
	}

	// Test that the removal fails.
	if err := testTransitionCycleWithPermutations(testDirectory1Entry, testDirectory1ContentMap, modifier, true); err != nil {
		t.Error(err)
	}
}

func TestTransitionFailOnParentPathIsFile(t *testing.T) {
	// Create a temporary file and defer its removal.
	var parent string
	if file, err := ioutil.TempFile("", "mutagen_simulated"); err != nil {
		t.Fatal("unable to create temporary file:", err)
	} else if err = file.Close(); err != nil {
		t.Fatal("unable to close temporary file:", err)
	} else {
		parent = file.Name()
	}
	defer os.Remove(parent)

	// Compute the path to the root.
	root := filepath.Join(parent, "root")

	// Set up the creation transitions.
	transitions := []*Change{{New: testDirectory1Entry}}

	// Create a provider and ensure its cleanup.
	provider, err := newTestProvider(testDirectory1ContentMap, newTestHasher())
	if err != nil {
		t.Fatal("unable to create test provider:", err)
	}
	defer provider.finalize()

	// Perform the creation transition and ensure that it encounters a problem.
	if entries, problems := Transition(root, transitions, nil, SymlinkMode_SymlinkPortable, false, provider); len(problems) != 1 {
		t.Error("transition succeeded unexpectedly")
	} else if len(entries) != 1 {
		t.Error("transition returned invalid number of entries")
	} else if entries[0] != nil {
		t.Error("failed creation transition returned non-nil entry")
	}
}