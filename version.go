package igdb

// VersionService handles all the API calls
// for the IGDB Versions endpoint.
type VersionService service

// Version contains information on an IGDB entry for details about game
// editions and versions. Version does not support the Search function.
//
// For more information, visit: https://igdb.github.io/api/endpoints/versions/
type Version struct {
	ID        int       `json:"id"`
	Game      int       `json:"game"`
	CreatedAt int       `json:"created_at"` // Unix time in milliseconds
	UpdatedAt int       `json:"updated_at"` // Unix time in milliseconds
	Games     []int     `json:"games"`
	URL       URL       `json:"url"`
	Features  []Feature `json:"features"`
}

// Feature contains information on a feature
// included with a particular version of a game.
type Feature struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Category    FeatureCategory `json:"category"`
	Position    int             `json:"position"`
	Values      []FeatureValue  `json:"values"`
}

// FeatureValue describes a type of Feature.
type FeatureValue struct {
	Game  int    `json:"game"`
	Value string `json:"value"`
}

// Get returns a single Version identified by the provided IGDB ID. Provide
// the SetFields functional option if you need to specify which fields to retrieve.
// If the ID does not match any Versions, an error is returned.
func (vs *VersionService) Get(id int, opts ...FuncOption) (*Version, error) {
	url, err := vs.client.singleURL(VersionEndpoint, id, opts...)
	if err != nil {
		return nil, err
	}

	var v []Version

	err = vs.client.get(url, &v)
	if err != nil {
		return nil, err
	}

	return &v[0], nil
}

// List returns a list of Versions identified by the provided list of IGDB IDs.
// Provide functional options to sort, filter, and paginate the results. Omitting
// IDs will instead retrieve an index of Versions based solely on the provided
// options. Any ID that does not match a Version is ignored. If none of the IDs
// match a Version, an error is returned.
func (vs *VersionService) List(ids []int, opts ...FuncOption) ([]*Version, error) {
	url, err := vs.client.multiURL(VersionEndpoint, ids, opts...)
	if err != nil {
		return nil, err
	}

	var v []*Version

	err = vs.client.get(url, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Count returns the number of Versions available in the IGDB.
// Provide the SetFilter functional option if you need to filter
// which Versions to count.
func (vs *VersionService) Count(opts ...FuncOption) (int, error) {
	ct, err := vs.client.getEndpointCount(VersionEndpoint, opts...)
	if err != nil {
		return 0, err
	}

	return ct, nil
}

// ListFields returns the up-to-date list of fields in an
// IGDB Version object.
func (vs *VersionService) ListFields() ([]string, error) {
	fl, err := vs.client.getEndpointFieldList(VersionEndpoint)
	if err != nil {
		return nil, err
	}

	return fl, nil
}
