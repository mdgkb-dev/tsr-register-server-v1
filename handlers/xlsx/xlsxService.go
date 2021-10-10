package xlsx

func (s *XlsxService) GetFile() ([]byte, error) {
	return s.xlsxHelper.CreateFile()
}
