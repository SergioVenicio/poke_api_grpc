package entities

type Details struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type MoveDetails struct {
	Details
}
type MoveLearnMethod struct {
	Details
}
type VersionGroup struct {
	Details
}
type VersionGroupDetails struct {
	LevelLearnedAt  int64           `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
}

type Move struct {
	MoveDetails         MoveDetails           `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type AbilityDetail struct {
	Details
}
type Ability struct {
	AbilityDetail AbilityDetail `json:"ability"`
	IsHidden      bool          `json:"is_hidden"`
	Slot          int64         `json:"slot"`
}

type Form struct {
	Details
}

type GameIndexVersion struct {
	Details
}
type GameIndex struct {
	GameIndex int64            `json:"game_index"`
	Version   GameIndexVersion `json:"version"`
}

type ItemDetail struct {
	Details
}
type VersionDetail struct {
	Details
}
type HeldDetail struct {
	Rarity  int64         `json:"rarity"`
	Version VersionDetail `json:"version"`
}
type HeldItem struct {
	Item       ItemDetail   `json:"item"`
	HeldDetail []HeldDetail `json:"version_details"`
}
type Pokemon struct {
	Name           string      `json:"name"`
	Id             int64       `json:"id"`
	BaseExperience int64       `json:"base_experience"`
	Height         int64       `json:"height"`
	Moves          []Move      `json:"moves"`
	Abilities      []Ability   `json:"abilities"`
	Forms          []Form      `json:"forms"`
	HeldItems      []HeldItem  `json:"held_items"`
	GameIndices    []GameIndex `json:"game_indices"`
}
