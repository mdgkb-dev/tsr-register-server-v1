package meta

func (s *Service) GetCount(table *string) (*int, error) {
	return s.repository.getCount(table)
}
