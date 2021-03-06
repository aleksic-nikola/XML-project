const DOCKERIZED = 'yes'
const AUTH_SERVICE_URL = getCorrectOrigin('auth')
const CAMPAIGN_SERVICE_URL = getCorrectOrigin('campaign')
const CONTENT_SERVICE_URL = getCorrectOrigin('content')
const INTERACTION_SERVICE_URL = getCorrectOrigin('interaction')
const MONOLIT_SERVICE_URL = getCorrectOrigin('monolit')
const PROFILE_SERVICE_URL = getCorrectOrigin('profile')
const REQUEST_SERVICE_URL = getCorrectOrigin('request')
const SEARCH_SERVICE_URL = getCorrectOrigin('service')
const FOLDER_PREFIX = getCorrectOrigin('prefix')

function getCorrectOrigin(name) {

	if (DOCKERIZED != 'yes') {
		// return local variables

		if (name == 'auth') {
			return "http://localhost:9090"
		}
		if (name == 'campaign') {
			return "http://localhost:7071"
		}
		if (name == 'content') {
			return "http://localhost:8080"
		}
		if (name == 'interaction') {
			return "http://localhost:5050"
		}
		if (name == 'monolit') {
			return "http://localhost:2902"
		}
		if (name == 'profile') {
			return "http://localhost:3030"
		}
		if (name == 'request') {
			return "http://localhost:9211"
		}
		if (name == 'search') {
			return "http://localhost:9494"
		}
		if (name == 'prefix') {
			return ""
		}

	}
	// else return gateway 
	if (name == 'auth') {
		return "http://localhost:8888/api/auth"
	}
	if (name == 'campaign') {
		return "http://localhost:8888/api/campaign"
	}
	if (name == 'content') {
		return "http://localhost:8888/api/content"
	}
	if (name == 'interaction') {
		return "http://localhost:8888/api/interaction"
	}
	if (name == 'monolit') {
		return "http://localhost:8888/api/monolit"
	}
	if (name == 'profile') {
		return "http://localhost:8888/api/profile"
	}
	if (name == 'request') {
		return "http://localhost:8888/api/request"
	}
	if (name == 'search') {
		return "http://localhost:8888/api/search"
	}
	if (name == 'prefix') {
		return "/app"
	}

	return "this should not be happening"
}

function printOriginVariables() {

	console.log(AUTH_SERVICE_URL)
	console.log(CAMPAIGN_SERVICE_URL)
	console.log(CONTENT_SERVICE_URL)
	console.log(INTERACTION_SERVICE_URL)
	console.log(MONOLIT_SERVICE_URL)
	console.log(PROFILE_SERVICE_URL)
	console.log(REQUEST_SERVICE_URL)
	console.log(SEARCH_SERVICE_URL)
}