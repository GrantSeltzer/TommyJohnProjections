package bbsavant

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/agrison/go-tablib"
)

// Spreadsheet blah
type Spreadsheet struct {
	filenames []string
	dataset   *tablib.Dataset
}

// New creates a new Spreadsheet object
func New() *Spreadsheet {
	fn := []string{}
	ds := NewDataset()
	return &Spreadsheet{fn, ds}
}

// NewDataset creates a new dataset
func NewDataset() *tablib.Dataset {
	return tablib.NewDataset(ColumnNames)
}

// SetDataset is the mutator for Spreadsheet.dataset
func (s *Spreadsheet) SetDataset(ds *tablib.Dataset) {
	s.dataset = ds
}

// SetFilenames is the mutator for Spreadsheet.filenames
func (s *Spreadsheet) SetFilenames(fn []string) {
	s.filenames = fn
}

// ReadFilesIntoDataset read the files set in Spreadsheet.filenames into Spreadhseet.dataset
func (s *Spreadsheet) ReadFilesIntoDataset() error {
	for _, filename := range s.filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		s.dataset.Append([]interface{}{data})
	}
	fmt.Println(s.dataset)
	return nil
}

// ReadFile takes a csv file and returns a 3D array of it
func ReadFile(filename string) ([][]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	strreader := strings.NewReader(string(data))
	csvreader := csv.NewReader(strreader)
	result, err := csvreader.ReadAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
