const search_field = $('#search_input')
var public_profiles 
var public_names = {}
var allowed_names = {}
var allowed_profiles
var final_profiles = {}

search_field.focusin(function() {

	console.log('focus')
})


function who_am_I() {

	$.ajax({
		type: 'GET',
		crossDomain: true,
		url: AUTH_SERVICE_URL + '/whoami',
		contentType: 'application/json',
		dataType: 'JSON',
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success: function (data) {
		    console.log('using authenticated search..')
		    authenticatedSearch(data)
		},
		error: function () {
		    console.log('using non authenticated search..')
		    nonAuthenticatedSearch(data)
		}
	    })

}

function authenticatedSearch() {
	getAllPublicProfiles()
	getAllPublicPosts()
	getAllAllowedProfiles()
	getAllAllowedPosts()
	createFinalProfileList()
}

function createFinalProfileList() {
	
}

function nonAuthenticatedSearch() {
	getAllPublicProfiles()
	getAllPublicPosts()
}

function getAllPublicProfiles() {

}

function getAllPublicPosts() {

}

function getAllAllowedProfiles() {

}

function getAllAllowedPosts() {

}