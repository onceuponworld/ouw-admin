package main

const (
	APP_NAME							= "ouw-admin"
	APP_VERSION						= "1.0"
)


const (
	DEFAULT_CONF_FILE			= "./config.json"
)


const (
	DEFAULT_KINGDOM_NAME_LENGTH			= 2
	DEFAULT_CHARS_MIN       				= 1
	DEFAULT_FORM_SIZE								= 20000
)


const (
	API_PARAM_PROFILE					= "profile"
	API_PARAM_COWS						= "cows"
	API_PARAM_KINGDOM_NAME		= "name"
	API_PARAM_LAND						= "land"
	API_PARAM_NAME						= "name"
	API_PARAM_POPULATION			= "population"
	API_PARAM_ROCKS						= "rocks"
	API_PARAM_TREES						= "trees"
	API_PARAM_WEALTH					= "wealth"
)


const (
	ERR_KINGDOM_NAME					= "Kingdom name length must be a-z and at least 2 characters."
	ERR_KINGDOM_PARAM_ENPTY   = "Parameter cannot by empty."
	ERR_UNSUPPORTED_METHOD    = "HTTP method not supported"
)


type Municipal struct {
	Name									string      `json:"name"`
	Population            string      `json:"population"`
	Land									int					`json:"land"`
	Wealth								int					`json:"wealth"`
	Trees             		int					`json:"trees"`
	Rocks                 int         `json:"rocks"`
	Grass                 int         `json:"grass"`
	Cows                 	int         `json:"cows"`
}


type Profile struct {
	Name									string      `json:"name"`
	Capital               string      `json:"capital"`
	MaleRatio							int					`json:"maleRatio"`
	MedianAge							int					`json:"medianAge"`
	BirthRate             int					`json:"birthRate"`
	DeathRate      				int					`json:"deathRate"`
	Population            int					`json:"population"`
	Land                  int         `json:"land"`
	Wealth                int         `json:"wealth"`
	Trees                 int         `json:"trees"`
	Rocks                 int         `json:"rocks"`
	Grass                 int         `json:"grass"`
	Cows                 	int         `json:"cows"`
	TaxRate               int         `json:"taxRate"`
	ConscriptAge    			int         `json:"conscriptAge"`
	Municipals            [] Municipal 	`json"municipals"`
}