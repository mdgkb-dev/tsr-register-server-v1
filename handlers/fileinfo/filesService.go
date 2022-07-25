package fileinfo

func (s *FilesService) GetFullPath(fileSystemPath *string) *string {
	return s.helper.Uploader.GetFullPath(fileSystemPath)
}
