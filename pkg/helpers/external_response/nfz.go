package external_response

// Response from NFZ api
type NFZResponse struct {
	Meta struct {
		Context       string `json:"context"`
		Count         int    `json:"count"`
		Title         string `json:"title"`
		Page          int    `json:"page"`
		URL           string `json:"url"`
		Limit         int    `json:"limit"`
		Provider      string `json:"provider"`
		DatePublished string `json:"date-published"`
		DateModified  string `json:"date-modified"`
		Description   string `json:"description"`
		Keywords      string `json:"keywords"`
		Language      string `json:"language"`
		ContentType   string `json:"content-type"`
		IsPartOf      string `json:"is-part-of"`
		Message       struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Self  string `json:"self"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"links"`
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Case                int     `json:"case"`
			Benefit             string  `json:"benefit"`
			ManyPlaces          string  `json:"many-places"`
			Provider            string  `json:"provider"`
			ProviderCode        string  `json:"provider-code"`
			RegonProvider       string  `json:"regon-provider"`
			NipProvider         string  `json:"nip-provider"`
			TerytProvider       string  `json:"teryt-provider"`
			Place               string  `json:"place"`
			Address             string  `json:"address"`
			Locality            string  `json:"locality"`
			Phone               string  `json:"phone"`
			TerytPlace          string  `json:"teryt-place"`
			RegistryNumber      string  `json:"registry-number"`
			IDResortPartVII     string  `json:"id-resort-part-VII"`
			IDResortPartVIII    string  `json:"id-resort-part-VIII"`
			BenefitsForChildren string  `json:"benefits-for-children"`
			Covid19             string  `json:"covid-19"`
			Toilet              string  `json:"toilet"`
			Ramp                string  `json:"ramp"`
			CarPark             string  `json:"car-park"`
			Elevator            string  `json:"elevator"`
			Latitude            float32 `json:"latitude"`
			Longitude           float32 `json:"longitude"`
			Statistics          struct {
				ProviderData struct {
					Awaiting      int    `json:"awaiting"`
					Removed       int    `json:"removed"`
					AveragePeriod int    `json:"average-period"`
					Update        string `json:"update"`
				} `json:"provider-data"`
				ComputedData struct {
					AveragePeriod int    `json:"average-period"`
					Update        string `json:"update"`
				} `json:"computed-data"`
			} `json:"statistics"`
			Dates struct {
				Applicable        bool   `json:"applicable"`
				Date              string `json:"date"`
				DateSituationAsAt string `json:"date-situation-as-at"`
			} `json:"dates"`
			BenefitsProvided struct {
				TypeOfBenefit int `json:"type-of-benefit"`
				Year          int `json:"year"`
				Amount        int `json:"amount"`
			} `json:"benefits-provided"`
		} `json:"attributes"`
	} `json:"data"`
}
