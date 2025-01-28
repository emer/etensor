// Code generated by "core generate"; DO NOT EDIT.

package databrowser

import (
	"io/fs"

	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/etensor/tensor/databrowser.Browser", IDName: "browser", Doc: "Browser is a data browser, for browsing data either on an os filesystem\nor as a datafs virtual data filesystem.", Methods: []types.Method{{Name: "UpdateFiles", Doc: "UpdateFiles Updates the file picker with current files in DataRoot,", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "FS", Doc: "Filesystem, if browsing an FS"}, {Name: "DataRoot", Doc: "DataRoot is the path to the root of the data to browse"}, {Name: "toolbar"}}})

// NewBrowser returns a new [Browser] with the given optional parent:
// Browser is a data browser, for browsing data either on an os filesystem
// or as a datafs virtual data filesystem.
func NewBrowser(parent ...tree.Node) *Browser { return tree.New[Browser](parent...) }

// SetFS sets the [Browser.FS]:
// Filesystem, if browsing an FS
func (t *Browser) SetFS(v fs.FS) *Browser { t.FS = v; return t }

// SetDataRoot sets the [Browser.DataRoot]:
// DataRoot is the path to the root of the data to browse
func (t *Browser) SetDataRoot(v string) *Browser { t.DataRoot = v; return t }

var _ = types.AddType(&types.Type{Name: "github.com/emer/etensor/tensor/databrowser.FileNode", IDName: "file-node", Doc: "FileNode is databrowser version of FileNode for FileTree", Methods: []types.Method{{Name: "EditFiles", Doc: "EditFiles calls EditFile on selected files", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "PlotFiles", Doc: "PlotFiles calls PlotFile on selected files", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "DiffDirs", Doc: "DiffDirs displays a browser with differences between two selected directories", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}}, Embeds: []types.Field{{Name: "Node"}}})

// NewFileNode returns a new [FileNode] with the given optional parent:
// FileNode is databrowser version of FileNode for FileTree
func NewFileNode(parent ...tree.Node) *FileNode { return tree.New[FileNode](parent...) }
