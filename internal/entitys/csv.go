package entitys

import "math"

// Csv é a struct principal que representa um arquivo csv
type Csv struct {
	Name string
	Size float64
}

// NewCsv retorna uma nova entidade CSV
func NewCsv(name string, size int64) *Csv {
	return &Csv{
		Name: name,
		Size: float64(size),
	}
}

// ConvertSizeToMB convert o tamanho do arquivo para megabytes com duas casas decimais
func (c *Csv) ConvertSizeToMB() {
	mb := c.Size / 1024
	c.Size = math.Floor(mb*100) / 100
}
