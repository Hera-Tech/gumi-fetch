package types

type MALMediaTypes string

const (
	MALMediaTypesUnknown MALMediaTypes = "unknown"
	MALMediaTypesTV      MALMediaTypes = "tv"
	MALMediaTypesOVA     MALMediaTypes = "ova"
	MALMediaTypesMovie   MALMediaTypes = "movie"
	MALMediaTypesSpecial MALMediaTypes = "special"
	MALMediaTypesONA     MALMediaTypes = "ona"
	MALMediaTypesMusic   MALMediaTypes = "music"
)

type MALStatus string

const (
	MALStatusFinishedAiring  MALStatus = "finished_airing"
	MALStatusCurrentlyAiring MALStatus = "currently_airing"
	MALStatusNotYetAired     MALStatus = "not_yet_aired"
)
