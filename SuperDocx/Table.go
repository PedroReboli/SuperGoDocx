package docx

//ReplaceAll Search for old string and repalce in entire table
func (t *table) ReplaceAll(old string, new string) {
	for _, r := range t.Rows {
		for _, c := range r.Cols {
			for _, p := range c.Paragraphs {
				for p.Replace(old, new) == true {

				}
			}
		}
	}
}
