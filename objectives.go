package okr2go

var objectives Objectives

type Objectives []*Objective

type Objective struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	KeyResults  []*KeyResult `json:"keyResults"`
}

func LoadObjectives() (err error) {
	objectives, err = ParseMarkdown("example.md")

	if err != nil {
		return err
	}

	return nil
}

func SaveObjectives() (err error) {
	err = WriteMarkdown("example.md", objectives)

	if err != nil {
		return err
	}

	return nil
}

func (o Objectives) FindObjective(objectiveID int) *Objective {
	if objectiveID < 0 || objectiveID > len(o) {
		return nil
	}

	return o[objectiveID]
}

func (o *Objective) FindKeyResult(resultID string) *KeyResult {
	for _, result := range o.KeyResults {
		if result.ID == resultID {
			return result
		}
	}

	return nil
}

type KeyResult struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Current      int64    `json:"current"`
	Target       int64    `json:"target"`
	Contributors []string `json:"contributors"`
	Comments     []string `json:"comments"`
}
