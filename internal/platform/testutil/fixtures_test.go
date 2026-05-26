/*
Copyright © 2026 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

// Inline: pure function, scalar assertions

func writeFixtureFile(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadTestdataFS(t *testing.T) {
	dir := t.TempDir()
	writeFixtureFile(t, filepath.Join(dir, "a.txt"), []byte("hello"))
	writeFixtureFile(t, filepath.Join(dir, "sub", "b.json"), []byte(`{"ok":true}`))

	mfs := LoadTestdataFS(t, dir, "/")

	data, err := mfs.ReadFile("/a.txt")
	if err != nil {
		t.Fatalf("expected /a.txt in MapFS: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("got %q, want %q", string(data), "hello")
	}

	data, err = mfs.ReadFile("/sub/b.json")
	if err != nil {
		t.Fatalf("expected /sub/b.json in MapFS: %v", err)
	}
	if string(data) != `{"ok":true}` {
		t.Errorf("got %q, want %q", string(data), `{"ok":true}`)
	}
}

func TestLoadTestdataFS_CustomRoot(t *testing.T) {
	dir := t.TempDir()
	writeFixtureFile(t, filepath.Join(dir, "f.txt"), []byte("content"))

	mfs := LoadTestdataFS(t, dir, "/root")

	data, err := mfs.ReadFile("/root/f.txt")
	if err != nil {
		t.Fatalf("expected /root/f.txt in MapFS: %v", err)
	}
	if string(data) != "content" {
		t.Errorf("got %q, want %q", string(data), "content")
	}
}

func TestReadFixture(t *testing.T) {
	dir := t.TempDir()
	writeFixtureFile(t, filepath.Join(dir, "test.txt"), []byte("fixture"))

	mfs := LoadTestdataFS(t, dir, "/")
	got := ReadFixture(t, mfs, "/test.txt")

	if string(got) != "fixture" {
		t.Errorf("got %q, want %q", string(got), "fixture")
	}
}

func TestCheckGolden_WithFS(t *testing.T) {
	dir := t.TempDir()
	writeFixtureFile(t, filepath.Join(dir, "goldens", "output.txt"), []byte("expected"))

	mfs := LoadTestdataFS(t, dir, "/")

	CheckGolden(t, "output", []byte("expected"), GoldenOptions{
		Dir:       "goldens",
		Extension: ".txt",
		FS:        mfs,
	})
}

func TestCheckGolden_WithFS_ReadsMissingFile(t *testing.T) {
	dir := t.TempDir()
	writeFixtureFile(t, filepath.Join(dir, "goldens", "exists.txt"), []byte("content"))

	mfs := LoadTestdataFS(t, dir, "/")

	// Verify ReadFile on a missing path returns an error (CheckGolden would t.Fatalf)
	_, err := mfs.ReadFile("goldens/nonexistent.txt")
	if err == nil {
		t.Error("expected error reading nonexistent file from MapFS")
	}
}
