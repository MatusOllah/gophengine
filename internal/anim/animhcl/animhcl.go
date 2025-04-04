// Package animhcl allows loading from external HCL files (see `assets/images` for examples).
package animhcl

import (
	"cmp"
	"fmt"
	_ "image/png"
	"io/fs"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"time"

	"github.com/MatusOllah/gophengine/internal/anim"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/gocty"
)

type animation struct {
	Name          string   `hcl:"name,label"`
	Frames        []string `hcl:"frames"`
	FrameDuration string   `hcl:"frame_duration"`
}

type animController struct {
	Name        string      `hcl:"name,label"`
	DefaultAnim string      `hcl:"default_anim,optional"`
	Animations  []animation `hcl:"animation,block"`
}

func refineNonNull(b *cty.RefinementBuilder) *cty.RefinementBuilder {
	return b.NotNull()
}

func LoadAnimsFromFS(fsys fs.FS, _path string, name string) (*anim.AnimController, error) {
	src, err := fs.ReadFile(fsys, _path)
	if err != nil {
		return nil, err
	}

	evalctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"PATH":      cty.StringVal(path.Dir(_path)),
			"DUR_24FPS": cty.StringVal(anim.Dur24FPS.String()),
		},
		Functions: map[string]function.Function{
			"fromPrefix": function.New(&function.Spec{
				Description: "",
				Params: []function.Parameter{
					{
						Name: "path",
						Type: cty.String,
					},
					{
						Name: "prefix",
						Type: cty.String,
					},
				},
				Type:         function.StaticReturnType(cty.List(cty.String)),
				RefineResult: refineNonNull,
				Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
					frames, err := getFramesByPrefixFromFS(fsys, args[0].AsString(), args[1].AsString())
					if err != nil {
						return cty.ListValEmpty(cty.String), err
					}

					var list []cty.Value
					for _, f := range frames {
						list = append(list, cty.StringVal(f))
					}

					return cty.ListVal(list), nil
				},
			}),
			"fromIndices": function.New(&function.Spec{
				Description: "",
				Params: []function.Parameter{
					{
						Name: "path",
						Type: cty.String,
					},
					{
						Name: "prefix",
						Type: cty.String,
					},
					{
						Name: "indices",
						Type: cty.List(cty.Number),
					},
				},
				Type:         function.StaticReturnType(cty.List(cty.String)),
				RefineResult: refineNonNull,
				Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
					var indices []int
					if err := gocty.FromCtyValue(args[2], &indices); err != nil {
						return cty.ListValEmpty(cty.String), err
					}

					frames, err := getFramesByIndicesFromFS(fsys, args[0].AsString(), args[1].AsString(), indices)
					if err != nil {
						return cty.ListValEmpty(cty.String), err
					}

					var list []cty.Value
					for _, f := range frames {
						list = append(list, cty.StringVal(f))
					}

					return cty.ListVal(list), nil
				},
			}),
		},
	}

	var v struct {
		AnimControllers []animController `hcl:"controller,block"`
	}

	err = hclsimple.Decode(path.Base(_path), src, evalctx, &v)
	if err != nil {
		return nil, err
	}

	ac := anim.NewAnimController()
	for _, _ac := range v.AnimControllers {
		if _ac.Name != name {
			continue
		}

		for _, a := range _ac.Animations {
			dur, err := time.ParseDuration(a.FrameDuration)
			if err != nil {
				return nil, err
			}

			var images []*ebiten.Image
			for _, frame := range a.Frames {
				img, _, err := ebitenutil.NewImageFromFileSystem(fsys, path.Dir(_path)+"/"+frame)
				if err != nil {
					return nil, err
				}
				images = append(images, img)
			}

			ac.SetAnim(a.Name, anim.NewAnimation(images, dur))
		}

		ac.Play(_ac.DefaultAnim)
	}

	return ac, nil
}

func getFramesByPrefixFromFS(fsys fs.FS, path string, prefix string) ([]string, error) {
	var finalFiles []string
	err := fs.WalkDir(fsys, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			name := d.Name()

			if name[:len(name)-8] == prefix {
				finalFiles = append(finalFiles, name)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	slices.SortFunc(finalFiles, func(name1, name2 string) int {
		index1, _ := strconv.Atoi(name1[len(name1)-4:])
		index2, _ := strconv.Atoi(name2[len(name2)-4:])
		return index1 - index2
	})

	return finalFiles, nil
}

func getFramesByIndicesFromFS(fsys fs.FS, path string, prefix string, indices []int) ([]string, error) {
	var files []string
	err := fs.WalkDir(fsys, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			name := d.Name()

			if name[:len(name)-8] == prefix {
				files = append(files, name)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var fileIndices []int

	for _, index := range indices {
		wantedFile := prefix + fmt.Sprintf("%04d", index) + ".png"
		fileIndex, ok := slices.BinarySearchFunc(files, wantedFile, func(s1, s2 string) int {
			return cmp.Compare(filepath.Base(s1), filepath.Base(s2))
		})
		if !ok {
			return nil, fmt.Errorf("frame not found: %s", wantedFile)
		}
		fileIndices = append(fileIndices, fileIndex)
	}

	var finalFiles []string
	for _, i := range fileIndices {
		finalFiles = append(finalFiles, files[i])
	}

	return finalFiles, nil
}
