package fileInfo

func (s *FilesService) GetFullPath(fileSystemPath *string) *string {
	return s.uploader.GetFullPath(fileSystemPath)
}
