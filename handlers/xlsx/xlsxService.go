package xlsx

func (s *ServiceXLSX) GetFile() ([]byte, error) {
	return s.xlsxHelper.CreateFile()
}
